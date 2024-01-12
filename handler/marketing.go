package handler

import (
	"net/http"

	"github.com/tombrereton/go-hot-reload/model"
	"github.com/tombrereton/go-hot-reload/view/layout"
	"github.com/tombrereton/go-hot-reload/view/marketing"
)

type Marketing struct {
}

func (h *Marketing) GetLandingPage(w http.ResponseWriter, r *http.Request) {
	user := model.User{ID: 1, Name: "Bob Loblaw"}
	p := marketing.LandingPage(user)
	b := layout.Base(p)
	b.Render(r.Context(), w)
}

func (h *Marketing) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	p := marketing.AboutPage()
	b := layout.Base(p)
	b.Render(r.Context(), w)
}
