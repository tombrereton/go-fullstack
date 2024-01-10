package server

import (
	"html/template"
	"net/http"

	"github.com/aarol/reload"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tombrereton/go-hot-reload/internal/routes"
)

type Server struct {
	Router *chi.Mux
	Config *Config
}

func NewWebServer(cfg *Config) *Server {
	s := &Server{
		Router: chi.NewRouter(),
		Config: cfg,
	}
	templates := parseTemplates(s.Config.TemplateDir)

	if s.Config.IsDevelopment {
		s.MountHotReload(templates)
	}
	s.MountMiddleware()
	s.MountStaticFiles()
	s.MountPageHandlers(templates)
	return s
}

func (s *Server) MountHotReload(t *template.Template) {
	reload := reload.New(s.Config.TemplateDir, s.Config.StaticDir)
	reload.OnReload = func() {
		t = parseTemplates(s.Config.TemplateDir)
	}
	s.Router.Use(reload.Handle)
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountPageHandlers(t *template.Template) {
	s.Router.NotFound(routes.NotFound(t))
	s.Router.Mount("/", routes.LandingPage(t))
}

func (s *Server) MountStaticFiles() {
	fileServer := http.FileServer(http.Dir(s.Config.StaticDir))
	handler := http.StripPrefix("/static/", fileServer)
	s.Router.Handle("/static/*", handler)
}

func parseTemplates(templatePath string) *template.Template {
	return template.Must(template.ParseGlob(templatePath + "*.html"))
}
