package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Static struct {
}

func (h *Static) GetStaticFiles() *chi.Mux {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("web/static"))
	r.Handle("/*", http.StripPrefix("/static/", fs))

	return r
}
