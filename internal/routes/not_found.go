package routes

import (
	"html/template"
	"net/http"
)

func NotFound(t *template.Template) http.HandlerFunc {
	pageVariable := PageVariables{Name: "Not Found!"}

	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "landing.html", pageVariable)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
