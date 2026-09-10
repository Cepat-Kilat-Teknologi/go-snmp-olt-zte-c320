package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/config"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/internal/handler"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/internal/reqctx"
	"github.com/gosnmp/gosnmp"
)

// ---------------------------------------------------------------------------
// Mock SnmpRepositoryInterface — lightweight stub for registry tests.
// ---------------------------------------------------------------------------

type mockSnmpRepo struct {
	pingErr error
	closed  bool
}

func (m *mockSnmpRepo) Get([]string) (*gosnmp.SnmpPacket, error) { return nil, nil }
func (m *mockSnmpRepo) Walk(string, func(gosnmp.SnmpPDU) error) error {
	return nil
}
func (m *mockSnmpRepo) BulkWalk(string, func(gosnmp.SnmpPDU) error) error {
	return nil
}
func (m *mockSnmpRepo) Ping() error { return m.pingErr }
func (m *mockSnmpRepo) Close()      { m.closed = true }

// ---------------------------------------------------------------------------
// Helper: build a pre-populated OLTRegistry for tests. Each entry gets a mock
// SNMP repo and the shared mockOnuUsecase handler, avoiding real SNMP.
// ---------------------------------------------------------------------------

func testOLTEntry(id string, host string, userID int, boardPons map[int]int) *OLTEntry {
	uc := &mockOnuUsecase{}
	return &OLTEntry{
		OLT: config.OLTRuntimeConfig{
			ID:        id,
			Host:      host,
			Port:      161,
			Community: "public",
			UserID:    userID,
			BoardPons: boardPons,
		},
		Repo:      &mockSnmpRepo{},
		UC:        uc,
		Handler:   handler.NewOnuHandler(uc),
		BoardPons: boardPons,
		UserID:    userID,
	}
}

func testRegistry(defaultOLT string, entries ...*OLTEntry) *OLTRegistry {
	reg := &OLTRegistry{
		entries:    make(map[string]*OLTEntry, len(entries)),
		defaultOLT: defaultOLT,
		cfg:        &config.Config{},
	}
	for _, e := range entries {
		reg.entries[e.OLT.ID] = e
	}
	return reg
}

// ===================================================================
// OLTRegistry data-access methods
// ===================================================================

func TestOLTRegistry_Get(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 1, map[int]int{1: 16})
	reg := testRegistry("c320", e)

	got, ok := reg.Get("c320")
	if !ok || got != e {
		t.Fatal("Get(c320) should return the registered entry")
	}
	_, ok = reg.Get("nope")
	if ok {
		t.Fatal("Get(nope) should return false for missing OLT")
	}
}

func TestOLTRegistry_GetDefault(t *testing.T) {
	e1 := testOLTEntry("c320", "10.0.0.1", 1, map[int]int{1: 16})
	e2 := testOLTEntry("c300", "10.0.0.2", 2, map[int]int{3: 16})

	t.Run("configured default present", func(t *testing.T) {
		reg := testRegistry("c320", e1, e2)
		got, ok := reg.GetDefault()
		if !ok || got != e1 {
			t.Fatalf("expected c320, got %v ok=%v", got, ok)
		}
	})

	t.Run("fallback when default missing", func(t *testing.T) {
		reg := testRegistry("missing", e1, e2)
		got, ok := reg.GetDefault()
		if !ok || got == nil {
			t.Fatal("should fall back to any entry when default is missing")
		}
	})

	t.Run("empty registry", func(t *testing.T) {
		reg := testRegistry("c320")
		_, ok := reg.GetDefault()
		if ok {
			t.Fatal("empty registry should return false")
		}
	})
}

func TestOLTRegistry_DefaultOLTID(t *testing.T) {
	reg := testRegistry("c320")
	if reg.DefaultOLTID() != "c320" {
		t.Fatalf("expected c320, got %s", reg.DefaultOLTID())
	}
}

func TestOLTRegistry_List(t *testing.T) {
	e1 := testOLTEntry("b", "10.0.0.2", 1, nil)
	e2 := testOLTEntry("a", "10.0.0.1", 1, nil)
	reg := testRegistry("a", e1, e2)

	ids := reg.List()
	if len(ids) != 2 || ids[0] != "a" || ids[1] != "b" {
		t.Fatalf("expected sorted [a, b], got %v", ids)
	}
}

func TestOLTRegistry_Len(t *testing.T) {
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 1, nil),
		testOLTEntry("c300", "10.0.0.2", 2, nil),
	)
	if reg.Len() != 2 {
		t.Fatalf("expected 2, got %d", reg.Len())
	}
}

// ===================================================================
// Reconcile — remove / no-op / metadata update
// ===================================================================

func TestOLTRegistry_Reconcile_RemovesStaleOLTs(t *testing.T) {
	repo1 := &mockSnmpRepo{}
	repo2 := &mockSnmpRepo{}
	e1 := testOLTEntry("c320", "10.0.0.1", 1, nil)
	e1.Repo = repo1
	e2 := testOLTEntry("c300", "10.0.0.2", 2, nil)
	e2.Repo = repo2
	reg := testRegistry("c320", e1, e2)

	// Reconcile with only c320 — c300 should be removed.
	reg.Reconcile([]config.OLTRuntimeConfig{
		{ID: "c320", Host: "10.0.0.1", Port: 161, Community: "public"},
	})

	if reg.Len() != 1 {
		t.Fatalf("expected 1 OLT after reconcile, got %d", reg.Len())
	}
	if _, ok := reg.Get("c300"); ok {
		t.Fatal("c300 should be removed")
	}
	if !repo2.closed {
		t.Fatal("removed OLT's repo should be closed")
	}
	if repo1.closed {
		t.Fatal("kept OLT's repo should NOT be closed")
	}
}

func TestOLTRegistry_Reconcile_EmptyListRemovesAll(t *testing.T) {
	repo := &mockSnmpRepo{}
	e := testOLTEntry("c320", "10.0.0.1", 1, nil)
	e.Repo = repo
	reg := testRegistry("c320", e)

	reg.Reconcile(nil)

	if reg.Len() != 0 {
		t.Fatalf("expected 0 OLTs after empty reconcile, got %d", reg.Len())
	}
	if !repo.closed {
		t.Fatal("repo should be closed after removal")
	}
}

func TestOLTRegistry_Reconcile_UnchangedOLTNotRebuilt(t *testing.T) {
	repo := &mockSnmpRepo{}
	e := testOLTEntry("c320", "10.0.0.1", 1, map[int]int{1: 16, 2: 16})
	e.Repo = repo
	reg := testRegistry("c320", e)

	// Same connection params → should keep the existing entry.
	reg.Reconcile([]config.OLTRuntimeConfig{
		{ID: "c320", Host: "10.0.0.1", Port: 161, Community: "public", BoardPons: map[int]int{1: 16, 2: 16}},
	})

	if repo.closed {
		t.Fatal("unchanged OLT should NOT have its repo closed")
	}
	got, ok := reg.Get("c320")
	if !ok || got.Repo != repo {
		t.Fatal("unchanged OLT should keep the same repo instance")
	}
}

func TestOLTRegistry_Reconcile_MetadataUpdateInPlace(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 1, map[int]int{1: 16})
	repo := &mockSnmpRepo{}
	e.Repo = repo
	reg := testRegistry("c320", e)

	// Same host/port/community/walk/boards, but different UserID.
	reg.Reconcile([]config.OLTRuntimeConfig{
		{ID: "c320", Host: "10.0.0.1", Port: 161, Community: "public", UserID: 42, BoardPons: map[int]int{1: 16}},
	})

	got, _ := reg.Get("c320")
	if got.UserID != 42 {
		t.Fatalf("expected UserID=42 after metadata update, got %d", got.UserID)
	}
	if repo.closed {
		t.Fatal("metadata-only change should NOT rebuild the SNMP stack")
	}
}

func TestOLTRegistry_Reconcile_ConnectionChangeTriggersRebuild(t *testing.T) {
	oldRepo := &mockSnmpRepo{}
	e := testOLTEntry("c320", "10.0.0.1", 1, map[int]int{1: 16})
	e.Repo = oldRepo
	reg := testRegistry("c320", e)

	// Change host → old entry removed (repo closed), new entry created with
	// a fresh SNMP stack. SNMP uses UDP so Connect succeeds even when the host
	// is unreachable — the entry IS re-added with a real (but non-functional)
	// SNMP connection.
	reg.Reconcile([]config.OLTRuntimeConfig{
		{ID: "c320", Host: "10.0.0.99", Port: 161, Community: "public", BoardPons: map[int]int{1: 16}},
	})

	if !oldRepo.closed {
		t.Fatal("old repo should be closed on connection-param change")
	}
	// The new entry should exist with the updated host.
	got, ok := reg.Get("c320")
	if !ok {
		t.Fatal("c320 should be re-added after connection-param change")
	}
	if got.OLT.Host != "10.0.0.99" {
		t.Fatalf("expected new host 10.0.0.99, got %s", got.OLT.Host)
	}
	// The new repo should NOT be the old mock — it's a real SNMP repo now.
	if got.Repo == oldRepo {
		t.Fatal("new entry should have a different repo instance")
	}
}

// ===================================================================
// oltChangeKey
// ===================================================================

func Test_oltChangeKey(t *testing.T) {
	base := config.OLTRuntimeConfig{
		Host: "10.0.0.1", Port: 161, Community: "public",
		UseWalk: false, BoardPons: map[int]int{1: 16, 2: 16},
	}

	t.Run("same config same key", func(t *testing.T) {
		a := oltChangeKey(base)
		b := oltChangeKey(base)
		if a != b {
			t.Fatalf("same config should produce same key: %q vs %q", a, b)
		}
	})

	t.Run("different host different key", func(t *testing.T) {
		modified := base
		modified.Host = "10.0.0.2"
		if oltChangeKey(base) == oltChangeKey(modified) {
			t.Fatal("different host should produce different key")
		}
	})

	t.Run("different port different key", func(t *testing.T) {
		modified := base
		modified.Port = 162
		if oltChangeKey(base) == oltChangeKey(modified) {
			t.Fatal("different port should produce different key")
		}
	})

	t.Run("different community different key", func(t *testing.T) {
		modified := base
		modified.Community = "private"
		if oltChangeKey(base) == oltChangeKey(modified) {
			t.Fatal("different community should produce different key")
		}
	})

	t.Run("different walk different key", func(t *testing.T) {
		modified := base
		modified.UseWalk = true
		if oltChangeKey(base) == oltChangeKey(modified) {
			t.Fatal("different walk should produce different key")
		}
	})

	t.Run("different boards different key", func(t *testing.T) {
		modified := base
		modified.BoardPons = map[int]int{1: 16, 3: 8}
		if oltChangeKey(base) == oltChangeKey(modified) {
			t.Fatal("different board topology should produce different key")
		}
	})

	t.Run("userID change does NOT change key", func(t *testing.T) {
		modified := base
		modified.UserID = 999
		if oltChangeKey(base) != oltChangeKey(modified) {
			t.Fatal("UserID is metadata, should not change key")
		}
	})
}

// ===================================================================
// HealthCheck
// ===================================================================

func TestOLTRegistry_HealthCheck_AllHealthy(t *testing.T) {
	e1 := testOLTEntry("c320", "10.0.0.1", 1, nil)
	e1.Repo = &mockSnmpRepo{pingErr: nil}
	e2 := testOLTEntry("c300", "10.0.0.2", 2, nil)
	e2.Repo = &mockSnmpRepo{pingErr: nil}
	reg := testRegistry("c320", e1, e2)

	if err := reg.HealthCheck(context.Background()); err != nil {
		t.Fatalf("all healthy: expected nil, got %v", err)
	}
}

func TestOLTRegistry_HealthCheck_OneUnhealthy(t *testing.T) {
	e1 := testOLTEntry("c320", "10.0.0.1", 1, nil)
	e1.Repo = &mockSnmpRepo{pingErr: nil}
	e2 := testOLTEntry("c300", "10.0.0.2", 2, nil)
	e2.Repo = &mockSnmpRepo{pingErr: errors.New("timeout")}
	reg := testRegistry("c320", e1, e2)

	err := reg.HealthCheck(context.Background())
	if err == nil {
		t.Fatal("expected error when one OLT is unhealthy")
	}
}

func TestOLTRegistry_HealthCheck_ContextCancelled(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 1, nil)
	e.Repo = &mockSnmpRepo{pingErr: nil}
	reg := testRegistry("c320", e)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // immediately cancel
	err := reg.HealthCheck(ctx)
	if err == nil {
		t.Fatal("expected context error")
	}
}

func TestOLTRegistry_HealthCheck_EmptyRegistry(t *testing.T) {
	reg := testRegistry("c320")
	if err := reg.HealthCheck(context.Background()); err != nil {
		t.Fatalf("empty registry should be healthy, got %v", err)
	}
}

// ===================================================================
// Close
// ===================================================================

func TestOLTRegistry_Close(t *testing.T) {
	repo1 := &mockSnmpRepo{}
	repo2 := &mockSnmpRepo{}
	e1 := testOLTEntry("c320", "10.0.0.1", 1, nil)
	e1.Repo = repo1
	e2 := testOLTEntry("c300", "10.0.0.2", 2, nil)
	e2.Repo = repo2
	reg := testRegistry("c320", e1, e2)

	reg.Close()

	if !repo1.closed || !repo2.closed {
		t.Fatal("Close should close all repos")
	}
	if reg.Len() != 0 {
		t.Fatal("Close should empty the entries map")
	}
}

func TestOLTRegistry_Close_Idempotent(t *testing.T) {
	reg := testRegistry("c320", testOLTEntry("c320", "10.0.0.1", 1, nil))
	reg.Close()
	reg.Close() // second call should be a no-op
	if reg.Len() != 0 {
		t.Fatal("double Close should leave entries empty")
	}
}

// ===================================================================
// Dynamic route middleware: resolveOLT, resolveDefaultOLT, dynamicHandler
// ===================================================================

// TestLoadRoutesWithRegistry_PerOLTValidation mirrors
// TestLoadRoutesMulti_PerOLTValidation but exercises the dynamic registry path.
func TestLoadRoutesWithRegistry_PerOLTValidation(t *testing.T) {
	t.Setenv("API_KEY", "")

	e1 := testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16, 2: 16})
	e2 := testOLTEntry("c300a", "10.0.0.2", 0, map[int]int{3: 16, 5: 8})
	reg := testRegistry("c320", e1, e2)

	router := loadRoutesWithRegistry(reg, nil, nil, "")

	cases := []struct {
		path    string
		want400 bool
	}{
		{"/api/v1/olt/c320/board/1/pon/1", false},
		{"/api/v1/olt/c320/board/3/pon/1", true},    // slot 3 not on C320
		{"/api/v1/olt/c300a/board/3/pon/16", false},  // GTGH (16)
		{"/api/v1/olt/c300a/board/5/pon/8", false},   // GTGO (8)
		{"/api/v1/olt/c300a/board/5/pon/9", true},    // only 8 PONs
		{"/api/v1/olt/c300a/board/1/pon/1", true},    // slot 1 not on c300a
		{"/api/v1/board/1/pon/1", false},             // bare → default c320
		{"/api/v1/board/3/pon/1", true},              // bare default c320: slot 3 invalid
	}
	for _, tc := range cases {
		req := httptest.NewRequest("GET", tc.path, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		got400 := rr.Code == http.StatusBadRequest
		if got400 != tc.want400 {
			t.Errorf("%s: status=%d want400=%v", tc.path, rr.Code, tc.want400)
		}
	}
}

// TestLoadRoutesWithRegistry_UnknownOLT verifies an unknown OLT ID returns 404.
func TestLoadRoutesWithRegistry_UnknownOLT(t *testing.T) {
	t.Setenv("API_KEY", "")
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16}),
	)
	router := loadRoutesWithRegistry(reg, nil, nil, "")

	req := httptest.NewRequest("GET", "/api/v1/olt/nope/board/1/pon/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("unknown OLT: status=%d, want 404", rr.Code)
	}
}

// TestLoadRoutesWithRegistry_TenantIsolation mirrors the static test in
// multi_olt_routes_test.go but exercises the dynamic resolveOLT middleware.
func TestLoadRoutesWithRegistry_TenantIsolation(t *testing.T) {
	uc := okUsecase{&mockOnuUsecase{}}
	h1 := handler.NewOnuHandler(uc)
	h2 := handler.NewOnuHandler(uc)

	e1 := &OLTEntry{
		OLT:       config.OLTRuntimeConfig{ID: "c320", Host: "10.0.0.1", Port: 161, Community: "pub", UserID: 1, BoardPons: map[int]int{1: 16}},
		Repo:      &mockSnmpRepo{},
		UC:        uc,
		Handler:   h1,
		BoardPons: map[int]int{1: 16},
		UserID:    1,
	}
	e2 := &OLTEntry{
		OLT:       config.OLTRuntimeConfig{ID: "c300a", Host: "10.0.0.2", Port: 161, Community: "pub", UserID: 2, BoardPons: map[int]int{3: 16}},
		Repo:      &mockSnmpRepo{},
		UC:        uc,
		Handler:   h2,
		BoardPons: map[int]int{3: 16},
		UserID:    2,
	}
	reg := testRegistry("c320", e1, e2)

	users := map[string]reqctx.Principal{
		"keyA":     {UserID: 1},
		"keyB":     {UserID: 2},
		"adminKey": {Admin: true},
	}
	router := loadRoutesWithRegistry(reg, nil, users, "")

	cases := []struct {
		name string
		key  string
		path string
		want int
	}{
		{"user1 own OLT", "keyA", "/api/v1/olt/c320/board/1/pon/1", http.StatusOK},
		{"user1 other OLT -> 404", "keyA", "/api/v1/olt/c300a/board/3/pon/1", http.StatusNotFound},
		{"user2 own OLT", "keyB", "/api/v1/olt/c300a/board/3/pon/1", http.StatusOK},
		{"user2 other OLT -> 404", "keyB", "/api/v1/olt/c320/board/1/pon/1", http.StatusNotFound},
		{"admin sees c320", "adminKey", "/api/v1/olt/c320/board/1/pon/1", http.StatusOK},
		{"admin sees c300a", "adminKey", "/api/v1/olt/c300a/board/3/pon/1", http.StatusOK},
		{"missing key -> 401", "", "/api/v1/olt/c320/board/1/pon/1", http.StatusUnauthorized},
		{"bad key -> 401", "nope", "/api/v1/olt/c320/board/1/pon/1", http.StatusUnauthorized},
		{"user2 bare default(c320) -> 404", "keyB", "/api/v1/board/1/pon/1", http.StatusNotFound},
		{"user1 bare default(c320) -> 200", "keyA", "/api/v1/board/1/pon/1", http.StatusOK},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.path, nil)
			if tc.key != "" {
				req.Header.Set("X-API-Key", tc.key)
			}
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if rr.Code != tc.want {
				t.Errorf("%s %s: status=%d, want %d", tc.key, tc.path, rr.Code, tc.want)
			}
		})
	}
}

// TestLoadRoutesWithRegistry_CommonEndpoints verifies that unauthenticated
// endpoints (root, healthz, readyz, version, metrics) work through the
// registry-backed router just as they do through the static router.
func TestLoadRoutesWithRegistry_CommonEndpoints(t *testing.T) {
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16}),
	)
	router := loadRoutesWithRegistry(reg, nil, nil, "")

	for _, path := range []string{"/", "/healthz", "/readyz", "/version", "/metrics"} {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest("GET", path, nil))
		if rr.Code != http.StatusOK {
			t.Errorf("%s: status=%d, want 200", path, rr.Code)
		}
	}
}

// TestOLTEntryFromContext verifies the context round-trip.
func TestOLTEntryFromContext(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 1, nil)
	ctx := context.WithValue(context.Background(), oltEntryCtxKey, e)
	got := OLTEntryFromContext(ctx)
	if got != e {
		t.Fatal("OLTEntryFromContext should return the stored entry")
	}
}

// TestOLTEntryFromContext_PanicsOnMissing verifies the "must have middleware"
// contract: calling OLTEntryFromContext without resolveOLT panics.
func TestOLTEntryFromContext_PanicsOnMissing(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on missing context entry")
		}
	}()
	OLTEntryFromContext(context.Background())
}

// TestDynamicHandler verifies the method-expression wrapper resolves the
// handler from context and dispatches correctly.
func TestDynamicHandler(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 0, nil)
	ctx := context.WithValue(context.Background(), oltEntryCtxKey, e)

	// Build a request with the OLTEntry already in context (simulating the
	// resolveOLT middleware having run).
	req := httptest.NewRequest("GET", "/test", nil).WithContext(ctx)
	rr := httptest.NewRecorder()

	// Wrap GetUplinkTopology (it returns empty UplinkTopology → 200).
	h := dynamicHandler((*handler.OnuHandler).GetUplinkTopology)
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("dynamicHandler: status=%d, want 200", rr.Code)
	}
}

// ===================================================================
// StartPoller — only the "disabled" and "context cancellation" paths
// are testable without a real device-registry server.
// ===================================================================

func TestOLTRegistry_StartPoller_DisabledByZeroInterval(t *testing.T) {
	reg := testRegistry("c320")
	// interval <= 0 returns immediately without starting a ticker.
	reg.StartPoller(context.Background(), "http://unused", "", 0)
	// If it blocked, the test would time out. Reaching here = pass.
}

func TestOLTRegistry_StartPoller_StopsOnContextCancel(t *testing.T) {
	reg := testRegistry("c320")
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already cancelled

	done := make(chan struct{})
	go func() {
		reg.StartPoller(ctx, "http://unused", "", 1) // 1ns — tick immediately
		close(done)
	}()

	select {
	case <-done:
		// ok — poller exited
	case <-time.After(2 * time.Second): // 2s safety
		t.Fatal("StartPoller did not exit after context cancel")
	}
}

// ===================================================================
// loadRoutesWithRegistry — additional route coverage
// ===================================================================

// TestLoadRoutesWithRegistry_UplinkRoute verifies the /uplinks endpoint works
// through both per-OLT and bare (default) paths.
func TestLoadRoutesWithRegistry_UplinkRoute(t *testing.T) {
	t.Setenv("API_KEY", "")
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16}),
	)
	router := loadRoutesWithRegistry(reg, nil, nil, "")

	for _, path := range []string{"/api/v1/olt/c320/uplinks", "/api/v1/uplinks"} {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest("GET", path, nil))
		if rr.Code != http.StatusOK {
			t.Errorf("%s: status=%d, want 200", path, rr.Code)
		}
	}
}

// TestLoadRoutesWithRegistry_PaginateRoute verifies the /paginate path is
// mounted through the dynamic router.
func TestLoadRoutesWithRegistry_PaginateRoute(t *testing.T) {
	t.Setenv("API_KEY", "")
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16}),
	)
	router := loadRoutesWithRegistry(reg, nil, nil, "")

	paths := []string{
		"/api/v1/olt/c320/paginate/board/1/pon/1",
		"/api/v1/paginate/board/1/pon/1",
	}
	for _, path := range paths {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest("GET", path, nil))
		// The mock usecase returns empty data → handler returns a 404 JSON
		// response ("no data found"). That IS a valid handler response — the
		// route is mounted. Chi's unmatched-route 404 returns plain text
		// "404 page not found\n" (19 bytes). Distinguish by body length.
		if rr.Code == http.StatusNotFound && rr.Body.Len() < 30 {
			t.Errorf("%s: got chi 404, route not mounted", path)
		}
	}
}

// TestLoadRoutesWithRegistry_OnuDetailRoute verifies the nested
// /board/{b}/pon/{p}/onu/{o} route is mounted.
func TestLoadRoutesWithRegistry_OnuDetailRoute(t *testing.T) {
	t.Setenv("API_KEY", "")
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16}),
	)
	router := loadRoutesWithRegistry(reg, nil, nil, "")

	paths := []string{
		"/api/v1/olt/c320/board/1/pon/1/onu/1",
		"/api/v1/board/1/pon/1/onu/1",
	}
	for _, path := range paths {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest("GET", path, nil))
		// Same as PaginateRoute: handler 404 (JSON, >30 bytes) means mounted;
		// chi 404 ("404 page not found\n", 19 bytes) means not mounted.
		if rr.Code == http.StatusNotFound && rr.Body.Len() < 30 {
			t.Errorf("%s: got chi 404, route not mounted", path)
		}
	}
}

// TestLoadRoutesWithRegistry_MiddlewareHeaders verifies that global middleware
// (RequestID, version headers, security headers, CORS) is applied.
func TestLoadRoutesWithRegistry_MiddlewareHeaders(t *testing.T) {
	reg := testRegistry("c320",
		testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16}),
	)
	router := loadRoutesWithRegistry(reg, nil, nil, "")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, httptest.NewRequest("GET", "/", nil))

	if rr.Header().Get("X-Request-ID") == "" {
		t.Error("missing X-Request-ID header")
	}
	if rr.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Error("missing X-Content-Type-Options header")
	}
	if rr.Header().Get("X-API-Version") == "" {
		t.Error("missing X-API-Version header")
	}
}

// TestDynamicValidateBoardPon verifies the dynamic board/pon validator reads
// the topology from the context-injected OLTEntry, not a static closure.
func TestDynamicValidateBoardPon(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 0, map[int]int{1: 16})

	t.Run("valid board/pon passes through", func(t *testing.T) {
		// chi URL params are set by chi router context; for unit testing the
		// middleware in isolation we use the full router. This test is already
		// covered by TestLoadRoutesWithRegistry_PerOLTValidation above, so
		// we keep this as a minimal sanity check via the full router path.
		t.Setenv("API_KEY", "")
		reg := testRegistry("c320", e)
		router := loadRoutesWithRegistry(reg, nil, nil, "")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest("GET", "/api/v1/board/1/pon/1", nil))
		if rr.Code == http.StatusBadRequest {
			t.Error("valid board/pon should not return 400")
		}
	})

	t.Run("invalid board returns 400", func(t *testing.T) {
		t.Setenv("API_KEY", "")
		reg := testRegistry("c320", e)
		router := loadRoutesWithRegistry(reg, nil, nil, "")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest("GET", "/api/v1/board/99/pon/1", nil))
		if rr.Code != http.StatusBadRequest {
			t.Errorf("invalid board: status=%d, want 400", rr.Code)
		}
	})
}

// ===================================================================
// NewOLTRegistry constructor
// ===================================================================

func TestNewOLTRegistry(t *testing.T) {
	cfg := &config.Config{}
	reg := NewOLTRegistry(cfg, nil, "c320")

	if reg.defaultOLT != "c320" {
		t.Fatalf("expected default c320, got %s", reg.defaultOLT)
	}
	if reg.Len() != 0 {
		t.Fatal("new registry should be empty")
	}
	if reg.cfg != cfg {
		t.Fatal("config reference should be stored")
	}
}

// ===================================================================
// Helpers — verify test helper correctness
// ===================================================================

func Test_testOLTEntry(t *testing.T) {
	e := testOLTEntry("c320", "10.0.0.1", 42, map[int]int{1: 16})
	if e.OLT.ID != "c320" || e.OLT.Host != "10.0.0.1" || e.UserID != 42 {
		t.Fatalf("testOLTEntry fields wrong: %+v", e)
	}
	if e.Handler == nil || e.UC == nil || e.Repo == nil {
		t.Fatal("testOLTEntry should populate Handler, UC, Repo")
	}
	if fmt.Sprintf("%v", e.BoardPons) != "map[1:16]" {
		t.Fatalf("BoardPons wrong: %v", e.BoardPons)
	}
}
