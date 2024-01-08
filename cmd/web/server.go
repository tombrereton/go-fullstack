package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

type PageVariables struct {
	Name string
}

func NewWebServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountPageHandlers() {
	s.Router.Get("/", s.LandingPageHandler)
}

func (s *Server) MountStaticFiles() {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	s.Router.Handle("/static/*", handler)
}

func (s *Server) LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	pageVariable := PageVariables{Name: "Hot Reloading"}
	tmpl, err := template.ParseFiles("./templates/landing.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, pageVariable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
