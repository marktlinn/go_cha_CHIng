package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marktlinn/go_cha_CHIng/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/health-check", loadHealthCheck)

	return router
}

func loadHealthCheck(router chi.Router) {
	router.Get("/", handler.HealthCheck)
}
