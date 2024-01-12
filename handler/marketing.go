package handler

import (
	"net/http"

	"github.com/tombrereton/go-hot-reload/view/marketing"
)

type Marketing struct {
}

func (h *Marketing) GetLandingPage(w http.ResponseWriter, r *http.Request) {
	marketing.LandingPage().Render(r.Context(), w)
}

func (h *Marketing) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	marketing.AboutPage().Render(r.Context(), w)
}
