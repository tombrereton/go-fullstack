package main

import (
	"net/http"
)

func main() {
	s := NewWebServer()
	s.MountMiddleware()
	s.MountPageHandlers()
	s.MountStaticFiles()

	http.ListenAndServe(":8080", s.Router)
}