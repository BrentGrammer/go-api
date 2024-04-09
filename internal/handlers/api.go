package handlers // this is how we reference in other files: handlers.Handler etc.

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/BrentGrammer/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// add middleware to a route with r.Use
	r.Use(chimiddle.StripSlashes) // on all routes - removes ending slash in url to prevent a 404 error

	r.Get("/health", GetHealth) // note these handler functions are in the same package handlers even though they are in different files (so they are available and in scope)

	// define a route
	r.Route("/account", func(router chi.Router) {
		// auth middleware in the route
		router.Use(middleware.Authorization) // from our middleware package

		// account/coins endpoint
		router.Get("/coins", GetCoinBalance)
	})

}