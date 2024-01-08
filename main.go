package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/aarol/reload"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	s := NewWebServer()
	templates := parseTemplates()

	isDevelopment := os.Getenv("ENV") == "development"
	if isDevelopment {
		s.MountHotReload(templates)
	}
	s.MountMiddleware()
	s.MountStaticFiles()
	s.MountPageHandlers(templates)

	address := "localhost:3000"
	log.Default().Printf("Listening at http://%s", address)
	http.ListenAndServe(address, s.Router)
}

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

func (s *Server) MountHotReload(t *template.Template) {
	reload := reload.New("templates/")
	reload.OnReload = func() {
		t = parseTemplates()
	}
	s.Router.Use(reload.Handle)
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountPageHandlers(t *template.Template) {
	landingPageRouter := LandingPageRouter(t)
	s.Router.Mount("/", landingPageRouter)
}

func (s *Server) MountStaticFiles() {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	s.Router.Handle("/static/*", handler)
}

func LandingPageRouter(t *template.Template) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageVariable := PageVariables{Name: "Hot Reloading"}
		err := t.ExecuteTemplate(w, "landing.html", pageVariable)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return r
}

func parseTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}
