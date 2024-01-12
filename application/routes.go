package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tombrereton/go-hot-reload/handler"
)

func loadRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", marketingRoutes)

	return r
}

func marketingRoutes(router chi.Router) {
	marketingHandler := &handler.Marketing{}

	router.Get("/", marketingHandler.GetLandingPage)
	router.Get("/about", marketingHandler.GetAboutPage)
}
