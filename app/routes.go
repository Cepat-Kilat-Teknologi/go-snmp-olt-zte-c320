package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/handler"
)

func loadRoutes(ponHandler *handler.PonHandler) *chi.Mux {
	router := chi.NewRouter()

	// Middleware for logging requests
	router.Use(middleware.Logger)

	// Define a simple root endpoint
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, this is the root endpoint!"))
	})

	// Route for "pon" related endpoints
	router.Route("/pon", func(r chi.Router) {
		// Use the "ponHandler" to handle specific endpoints
		r.Get("/", ponHandler.List)
		r.Get("/{id}", ponHandler.GetByPonID)
	})

	router.Route("/gtgo", func(r chi.Router) {
		r.Get("/{gtgo_id}/pon/{id}", ponHandler.GetByGtGoIDAndPonID)
	})

	return router
}

func loadPonRoutes(router chi.Router) {
	ponHandler := &handler.PonHandler{}
	gtGoHandler := &handler.PonHandler{}

	// Pass the "ponHandler" to the "loadRoutes" function
	router.Mount("/", loadRoutes(ponHandler))
	router.Mount("/", loadRoutes(gtGoHandler))
}
