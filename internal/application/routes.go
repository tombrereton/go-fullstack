package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tombrereton/go-hot-reload/internal/handler"
)

func loadRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Route("/", marketingRoutes)
	r.Route("/static", staticRoutes)
	return r
}

func marketingRoutes(r chi.Router) {
	marketingHandler := &handler.Marketing{}

	r.Get("/", marketingHandler.GetLandingPage)
	r.Get("/about", marketingHandler.GetAboutPage)
}

func staticRoutes(r chi.Router) {
	staticHandler := &handler.Static{}

	r.Handle("/*", staticHandler.GetStaticFiles())
}
