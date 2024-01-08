package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}

func TestWebServer(t *testing.T) {
	s := NewWebServer()
	s.MountMiddleware()
	s.MountPageHandlers()
	s.MountStaticFiles()
	req, _ := http.NewRequest("GET", "/", nil)

	response := executeRequest(req, s)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	element := doc.Find("#hello").Text()
	got := strings.TrimSpace(element)
	want := `Hello, Hot Reloading!`

	if got != want {
		t.Errorf("Expected response body to be %s. Got %s", want, got)
	}
}
