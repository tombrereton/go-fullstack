package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}

func checkReponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestWebServer(t *testing.T) {
	s := NewWebServer()
	s.MountMiddleware()
	s.MountPageHandlers()
	s.MountStaticFiles()
	req, _ := http.NewRequest("GET", "/", nil)

	response := executeRequest(req, s)
	got := response.Body.String()
	want := `<h1>Hello, World!</h1>`

	if got != want {
		t.Errorf("Expected response body to be %s. Got %s\n", want, got)
	}
}
