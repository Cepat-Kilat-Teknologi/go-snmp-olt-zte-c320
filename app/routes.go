package app

import (
	"github.com/go-chi/chi/v5"
	mdl "github.com/go-chi/chi/v5/middleware"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/handler"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/middleware"
	"net/http"
)

func loadRoutes(onuHandler *handler.OnuHandler) http.Handler {
	router := chi.NewRouter()

	// Middleware for logging requests
	router.Use(mdl.Logger)

	// Middleware for CORS
	router.Use(middleware.CorsMiddleware())

	// Define a simple root endpoint
	router.Get("/", rootHandler)

	// Create a group for /api/v1/
	apiV1Group := chi.NewRouter()

	apiV1Group.Route("/board", func(r chi.Router) {
		r.Get("/{board_id}/pon/{pon_id}", onuHandler.GetByBoardIDAndPonID)
		r.Get("/{board_id}/pon/{pon_id}/onu/{onu_id}", onuHandler.GetByBoardIDPonIDAndOnuID)
		r.Get("/{board_id}/pon/{pon_id}/onu_id/empty", onuHandler.GetEmptyOnuID)
		r.Get("/{board_id}/pon/{pon_id}/onu_id/update", onuHandler.UpdateEmptyOnuID)
		r.Get("/{board_id}/page/pon/{pon_id}", onuHandler.GetByBoardIDAndPonIDWithPaginate)
		r.Get("/{board_id}/pon/{pon_id}/onu_id/empty/queue", onuHandler.GetEmptyOnuIDQueue)
	})
	router.Mount("/api/v1", apiV1Group)
	return router
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, this is the root endpoint!"))
}
