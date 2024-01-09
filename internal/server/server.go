package server

import (
	"html/template"
	"net/http"

	"github.com/aarol/reload"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tombrereton/go-hot-reload/internal/handlers"
)

type Server struct {
	Router *chi.Mux
	Config *Config
}

func NewWebServer(c *Config) *Server {
	s := &Server{}
	s.Router = chi.NewRouter()

	templates := parseTemplates(c.TemplatesPath)

	if c.IsDevelopment {
		s.MountHotReload(templates, c)
	}
	s.MountMiddleware()
	s.MountStaticFiles()
	s.MountPageHandlers(templates)
	return s
}

func (s *Server) MountHotReload(t *template.Template, c *Config) {
	reload := reload.New(c.TemplatesPath)
	reload.OnReload = func() {
		t = parseTemplates(c.TemplatesPath)
	}
	s.Router.Use(reload.Handle)
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountPageHandlers(t *template.Template) {
	s.Router.Mount("/", handlers.LandingPageRoutes(t))
}

func (s *Server) MountStaticFiles() {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	s.Router.Handle("/static/*", handler)
}

func parseTemplates(templatePath string) *template.Template {
	return template.Must(template.ParseGlob(templatePath + "*.html"))
}
