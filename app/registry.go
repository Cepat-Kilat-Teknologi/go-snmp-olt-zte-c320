package app

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/config"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/internal/handler"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/internal/repository"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/internal/usecase"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/pkg/logger"
	"github.com/Cepat-Kilat-Teknologi/snmp-olt-zte/pkg/snmp"
	"go.uber.org/zap"
)

// defaultPollInterval is the default device-registry poll interval. Override
// with REGISTRY_POLL_INTERVAL (Go duration); "0" disables polling entirely.
const defaultPollInterval = 30 * time.Second

// OLTEntry holds the full runtime stack for a single OLT: SNMP connection pool,
// repository, usecase, handler, and the topology metadata the router needs for
// board/pon validation and per-tenant auth. Exported so routes.go can read its
// fields through the registry.
type OLTEntry struct {
	OLT       config.OLTRuntimeConfig
	Repo      repository.SnmpRepositoryInterface
	UC        usecase.OnuUseCaseInterface
	Handler   *handler.OnuHandler
	BoardPons map[int]int // physical slot -> PON count (used by ValidateBoardPonParams)
	UserID    int         // owner tenant id (used by RequireOLTOwner)
}

// OLTRegistry is a thread-safe, dynamically-updatable registry of OLT runtime
// stacks. It replaces the static []oltRoute built at startup: new OLTs appear,
// changed OLTs reconnect, and removed OLTs clean up — all without a restart.
type OLTRegistry struct {
	mu         sync.RWMutex
	entries    map[string]*OLTEntry
	defaultOLT string
	cfg        *config.Config
	redisRepo  repository.OnuRedisRepositoryInterface
}

// NewOLTRegistry creates an empty registry. Call Reconcile to populate it with
// the initial OLT set, then StartPoller to keep it in sync with device-registry.
func NewOLTRegistry(cfg *config.Config, redisRepo repository.OnuRedisRepositoryInterface, defaultOLT string) *OLTRegistry {
	return &OLTRegistry{
		entries:    make(map[string]*OLTEntry),
		defaultOLT: defaultOLT,
		cfg:        cfg,
		redisRepo:  redisRepo,
	}
}

// Get returns the OLTEntry for the given OLT ID (read-locked). This is the hot
// path — every inbound request calls it via the resolveOLT middleware.
func (reg *OLTRegistry) Get(oltID string) (*OLTEntry, bool) {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	e, ok := reg.entries[oltID]
	return e, ok
}

// GetDefault returns the default OLT entry. Falls back to the first entry if
// the configured default is missing (e.g. it failed to init).
func (reg *OLTRegistry) GetDefault() (*OLTEntry, bool) {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	if e, ok := reg.entries[reg.defaultOLT]; ok {
		return e, true
	}
	// Fallback: return any entry (deterministic: sorted by ID).
	for _, e := range reg.entries {
		return e, true
	}
	return nil, false
}

// DefaultOLTID returns the configured default OLT identifier.
func (reg *OLTRegistry) DefaultOLTID() string {
	return reg.defaultOLT
}

// List returns the IDs of all registered OLTs (sorted, for deterministic output
// in health probes and logs).
func (reg *OLTRegistry) List() []string {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	ids := make([]string, 0, len(reg.entries))
	for id := range reg.entries {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

// Len returns the number of registered OLTs.
func (reg *OLTRegistry) Len() int {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	return len(reg.entries)
}

// Reconcile compares the current registry against newOLTs and performs the
// minimum set of add/remove/update operations to converge. It is the core diff
// engine called both at startup (from the initial config) and on every poll tick.
//
// Change detection: an OLT is "updated" when its connection-relevant fields
// (host, port, community, walk, boards) change — requiring a full SNMP
// reconnection. Organization/user ID changes are applied in-place without
// reconnection.
func (reg *OLTRegistry) Reconcile(newOLTs []config.OLTRuntimeConfig) {
	// Build lookup of incoming OLTs.
	incoming := make(map[string]config.OLTRuntimeConfig, len(newOLTs))
	for _, o := range newOLTs {
		incoming[o.ID] = o
	}

	reg.mu.Lock()
	defer reg.mu.Unlock()

	// Remove OLTs no longer in the incoming set.
	for id, entry := range reg.entries {
		if _, exists := incoming[id]; !exists {
			reg.removeOLTLocked(id, entry)
		}
	}

	// Add or update OLTs from the incoming set.
	for _, o := range newOLTs {
		existing, exists := reg.entries[o.ID]
		if !exists {
			// New OLT — create full stack.
			if err := reg.addOLTLocked(o); err != nil {
				logger.Error("olt_add_failed",
					zap.String("olt_id", o.ID),
					zap.String("host", o.Host),
					zap.Error(err))
			}
			continue
		}

		// Existing OLT — check if connection params changed.
		if oltChangeKey(existing.OLT) != oltChangeKey(o) {
			// Connection-relevant fields changed → full rebuild.
			logger.Info("olt_updated",
				zap.String("olt_id", o.ID),
				zap.String("old_host", existing.OLT.Host),
				zap.String("new_host", o.Host))
			reg.removeOLTLocked(o.ID, existing)
			if err := reg.addOLTLocked(o); err != nil {
				logger.Error("olt_update_failed",
					zap.String("olt_id", o.ID),
					zap.Error(err))
			}
			continue
		}

		// Connection params unchanged — update metadata in-place (user/org ID
		// may have changed in device-registry without requiring reconnection).
		existing.OLT = o
		existing.UserID = o.UserID
	}
}

// addOLTLocked creates the full per-OLT stack (SNMP conn → repo → usecase →
// handler) and stores it in the registry. Caller must hold reg.mu write lock.
func (reg *OLTRegistry) addOLTLocked(olt config.OLTRuntimeConfig) error {
	snmpConn, err := snmp.SetupSnmpConnectionWith(olt.Host, olt.Port, olt.Community)
	if err != nil {
		return fmt.Errorf("snmp connect %s:%d: %w", olt.Host, olt.Port, err)
	}

	snmpRepo := repository.NewPonRepositoryWithConcurrency(snmpConn, olt.MaxConcurrent, olt.UseWalk)

	cachePrefix := olt.ID
	if olt.ID == reg.defaultOLT {
		cachePrefix = "" // default OLT keeps unprefixed keys (back-compat)
	}
	uc := usecase.NewOnuUsecaseForOLT(snmpRepo, reg.redisRepo, reg.cfg.ForOLT(olt), cachePrefix)
	h := handler.NewOnuHandler(uc)

	reg.entries[olt.ID] = &OLTEntry{
		OLT:       olt,
		Repo:      snmpRepo,
		UC:        uc,
		Handler:   h,
		BoardPons: olt.BoardPons,
		UserID:    olt.UserID,
	}

	logger.Info("olt_added",
		zap.String("olt_id", olt.ID),
		zap.String("host", olt.Host),
		zap.Uint16("port", olt.Port),
		zap.Ints("boards", olt.Boards),
		zap.Bool("default", olt.ID == reg.defaultOLT))

	return nil
}

// removeOLTLocked closes the SNMP connection pool and deletes the entry from
// the registry. Caller must hold reg.mu write lock.
func (reg *OLTRegistry) removeOLTLocked(id string, entry *OLTEntry) {
	entry.Repo.Close()
	delete(reg.entries, id)

	logger.Info("olt_removed",
		zap.String("olt_id", id),
		zap.String("host", entry.OLT.Host))
}

// StartPoller runs a background goroutine that fetches the OLT list from
// device-registry at the given interval and reconciles the registry. It blocks
// until ctx is cancelled. Poll failures are logged but never wipe the registry
// — the current OLTs keep serving.
func (reg *OLTRegistry) StartPoller(ctx context.Context, registryURL, apiKey string, interval time.Duration) {
	if interval <= 0 {
		logger.Info("registry_poll_disabled")
		return
	}

	logger.Info("registry_poll_started",
		zap.String("url", registryURL),
		zap.Duration("interval", interval))

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			logger.Info("registry_poll_stopped")
			return
		case <-ticker.C:
			olts, err := config.FetchRegistryOLTConfigs(registryURL, apiKey)
			if err != nil {
				logger.Warn("registry_poll_failed, keeping current OLTs",
					zap.Error(err))
				continue
			}
			if olts == nil {
				logger.Warn("registry_poll_empty, keeping current OLTs")
				continue
			}
			reg.Reconcile(olts)
		}
	}
}

// Close tears down all OLT stacks, closing every SNMP connection pool.
// Safe to call multiple times (subsequent calls are no-ops on an empty map).
func (reg *OLTRegistry) Close() {
	reg.mu.Lock()
	defer reg.mu.Unlock()
	for id, entry := range reg.entries {
		entry.Repo.Close()
		logger.Info("olt_closed", zap.String("olt_id", id))
		delete(reg.entries, id)
	}
}

// HealthCheck pings every registered OLT's SNMP connection. Returns the first
// error encountered (with the OLT ID in the message) or nil if all are healthy.
func (reg *OLTRegistry) HealthCheck(ctx context.Context) error {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	for id, entry := range reg.entries {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err := entry.Repo.Ping(); err != nil {
			return fmt.Errorf("olt %s: %w", id, err)
		}
	}
	return nil
}

// oltChangeKey builds a string that uniquely identifies the connection-relevant
// configuration of an OLT. When this key changes between the current and
// incoming config, the SNMP connection must be rebuilt.
func oltChangeKey(o config.OLTRuntimeConfig) string {
	// Sort board specs for deterministic comparison.
	boards := make([]string, 0, len(o.BoardPons))
	for slot, pons := range o.BoardPons {
		boards = append(boards, fmt.Sprintf("%d:%d", slot, pons))
	}
	sort.Strings(boards)

	return fmt.Sprintf("%s|%d|%s|%t|%s",
		o.Host, o.Port, o.Community, o.UseWalk, strings.Join(boards, ","))
}
