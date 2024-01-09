package handlers

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PageVariables struct {
	Name string
}

func LandingPageRoutes(t *template.Template) *chi.Mux {
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
