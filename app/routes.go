package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mdl "github.com/go-chi/chi/v5/middleware"
	"github.com/megadata-dev/go-snmp-olt-c320/internal/handler"
	"github.com/megadata-dev/go-snmp-olt-c320/internal/middleware"
)

func loadRoutes(onuHandler *handler.OnuHandler) *chi.Mux {
	router := chi.NewRouter()

	// Middleware for logging requests
	router.Use(mdl.Logger)

	// Middleware for CORS
	router.Use(middleware.CorsMiddleware())

	// Define a simple root endpoint
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte("Hello, this is the root endpoint!"))
		if err != nil {
			return
		}
		_ = write
	})

	router.Route("/gtgo", func(r chi.Router) {
		r.Get("/{gtgo_id}/pon/{pon_id}", onuHandler.GetByGtGoIDAndPonID)
		//r.Get("/{gtgo_id}/pon/{pon_id}", onuHandler.GetByGtGoIDAndPonIDWithPaginate)
		r.Get("/{gtgo_id}/pon/{pon_id}/onu/{onu_id}", onuHandler.GetByGtGoIDPonIDAndOnuID)
		r.Get("/{gtgo_id}/pon/{pon_id}/empty", onuHandler.GetEmptyOnuID)
	})

	return router
}
