package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/aarol/reload"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tombrereton/go-hot-reload/internal/routes"
)

type Server struct {
	Router    *chi.Mux
	Config    *Config
	Templates *template.Template
}

func NewWebServer(cfg *Config) *Server {
	s := &Server{
		Router: chi.NewRouter(),
		Config: cfg,
	}
	s.AddHtmlTemplatesToServer()

	if s.Config.IsDevelopment {
		log.Default().Println("Hot reloading enabled")
		s.MountHotReload()
	}
	s.MountMiddleware()
	s.MountStaticFiles()
	s.MountPageHandlers()
	return s
}

func (s *Server) MountHotReload() {
	reload := reload.New(s.Config.TemplateDir, s.Config.StaticDir)
	reload.OnReload = func() {
		s.AddHtmlTemplatesToServer()
	}
	s.Router.Use(reload.Handle)
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountPageHandlers() {
	s.Router.NotFound(routes.NotFound(s.Templates))
	s.Router.Mount("/", routes.LandingPage(s.Templates))
	s.Router.Mount("/htmx", routes.Htmx(s.Templates))
}

func (s *Server) MountStaticFiles() {
	fileServer := http.FileServer(http.Dir(s.Config.StaticDir))
	handler := http.StripPrefix("/static/", fileServer)
	s.Router.Handle("/static/*", handler)
}

func (s *Server) AddHtmlTemplatesToServer() {
	s.Templates = template.Must(template.ParseGlob(s.Config.TemplateDir + "*.html"))
}
