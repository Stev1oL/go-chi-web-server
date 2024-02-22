package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Stev1oL/basic-web-server/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHancler := &handler.Order{}

	router.Post("/", orderHancler.Create)
	router.Get("/", orderHancler.List)
	router.Get("/{id}", orderHancler.GetByID)
	router.Put("/{id}", orderHancler.UpdateByID)
	router.Delete("/{id}", orderHancler.DeleteByID)
}
