package handlers

import (
	"go-api/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes) // strips the URL of trailing slashes
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/dob", GetUserDoB)
	})
}
