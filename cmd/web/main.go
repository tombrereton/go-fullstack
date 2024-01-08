package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	s := NewWebServer()
	s.MountMiddleware()
	s.MountPageHandlers()
	s.MountStaticFiles()

	http.ListenAndServe(":8080", s.Router)
}

type Server struct {
	Router *chi.Mux
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
	s.Router.Get("/", s.HomePageHandler)
}

func (s *Server) MountStaticFiles() {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	s.Router.Handle("/static/*", handler)
}

func (s *Server) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))
}
