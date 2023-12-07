package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/handler"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func loadRoutes(onuHandler *handler.OnuHandler) http.Handler {

	// Initialize logger
	l := log.Output(zerolog.ConsoleWriter{
		Out: os.Stdout,
	})

	// Initialize router using chi
	router := chi.NewRouter()

	// Middleware for logging requests
	router.Use(middleware.Logger(l))

	// Middleware for CORS
	router.Use(middleware.CorsMiddleware())

	// Define a simple root endpoint
	router.Get("/", rootHandler)

	// Create a group for /api/v1/
	apiV1Group := chi.NewRouter()

	// Define routes for /api/v1/
	apiV1Group.Route("/board", func(r chi.Router) {
		r.Get("/{board_id}/pon/{pon_id}", onuHandler.GetByBoardIDAndPonID)
		r.Get("/{board_id}/pon/{pon_id}/onu/{onu_id}", onuHandler.GetByBoardIDPonIDAndOnuID)
		r.Get("/{board_id}/pon/{pon_id}/onu_id/empty", onuHandler.GetEmptyOnuID)
		r.Get("/{board_id}/pon/{pon_id}/onu_id_sn", onuHandler.GetOnuIDAndSerialNumber)
		r.Get("/{board_id}/pon/{pon_id}/onu_id/update", onuHandler.UpdateEmptyOnuID)
	})

	// Define routes for /api/v1/paginate
	apiV1Group.Route("/paginate", func(r chi.Router) {
		r.Get("/board/{board_id}/pon/{pon_id}", onuHandler.GetByBoardIDAndPonIDWithPaginate)
	})

	// Mount /api/v1/ to root router
	router.Mount("/api/v1", apiV1Group)

	return router
}

// rootHandler is a simple handler for root endpoint
func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)                                // Set HTTP status code to 200
	_, _ = w.Write([]byte("Hello, this is the root endpoint!")) // Write response body
}
