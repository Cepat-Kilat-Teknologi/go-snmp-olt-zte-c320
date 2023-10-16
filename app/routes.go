package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/pon", loadPonRoutes)

	return router
}

func loadPonRoutes(router chi.Router) {
	ponHandler := &handler.PonHanlder{}

	router.Post("/", ponHandler.Create)
	router.Get("/", ponHandler.List)
	router.Get("/{id}", ponHandler.GetByID)
	router.Get("/{id}/pon", ponHandler.GetByPonID)
	router.Put("/{id}", ponHandler.UpdateByID)
	router.Delete("/{id}", ponHandler.DeleteByID)
}
