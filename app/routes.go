package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/handler"
)

func loadRoutes(onuHandler *handler.OnuHandler) *chi.Mux {
	router := chi.NewRouter()

	// Middleware for logging requests
	router.Use(middleware.Logger)

	// Define a simple root endpoint
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte("Hello, this is the root endpoint!"))
		if err != nil {
			return
		}
		_ = write
	})

	// Route for "pon" related endpoints
	router.Route("/pon", func(r chi.Router) {
		// Use the "onuHandler" to handle specific endpoints
		r.Get("/", onuHandler.List)
	})

	router.Route("/gtgo", func(r chi.Router) {
		r.Get("/{gtgo_id}/pon/{id}", onuHandler.GetByGtGoIDAndPonID)
	})

	return router
}
