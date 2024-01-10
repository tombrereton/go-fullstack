package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestIntegrationWebServer(t *testing.T) {
	config := Config{IsDevelopment: false, TemplateDir: "../../web/templates/", Port: "3000"}
	app := NewWebServer(&config)
	ts := httptest.NewServer(app.Router)
	defer ts.Close()
	client := ts.Client()
	response := GetReponse(client, ts, t)

	isOk(response, t)

	got := getElementText(response, t, "#hello")
	want := `Hello, Hot Reloading!`

	if got != want {
		t.Errorf("Expected response body to be %s. Got %s", want, got)
	}
}

func GetReponse(client *http.Client, ts *httptest.Server, t *testing.T) *http.Response {
	response, err := client.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	return response
}

func getElementText(response *http.Response, t *testing.T, elementId string) string {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	element := doc.Find(elementId).Text()
	got := strings.TrimSpace(element)
	return got
}

func isOk(response *http.Response, t *testing.T) {
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200. Got %d", response.StatusCode)
	}
}
