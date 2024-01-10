package routes

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Htmx(t *template.Template) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "htmx-demo.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return r
}
