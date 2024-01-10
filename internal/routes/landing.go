package routes

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PageVariables struct {
	Name string
}

func LandingPage(t *template.Template) *chi.Mux {
	r := chi.NewRouter()
	pageVariable := PageVariables{Name: "Hot Reloading"}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "landing.html", pageVariable)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return r
}
