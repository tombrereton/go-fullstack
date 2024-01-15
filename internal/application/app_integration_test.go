package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestApp_Start(t *testing.T) {
	cfg := NewConfigurationBuilder().WithPort("7000").Build()
	app := NewApp(cfg)

	server := httptest.NewServer(app.router)
	defer server.Close()

	client := server.Client()
	resp, err := client.Get(server.URL + "/")
	if err != nil {
		t.Errorf("Failed to send test request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	aboutLink := doc.Find("#about-link")
	if aboutLink.Length() == 0 {
		t.Errorf("Element with id 'about-link' not found")
	}

	attr, exists := aboutLink.Attr("hx-get")
	if !exists {
		t.Errorf("Href attribute not found for element with id 'about-link'")
	} else if attr != "/about" {
		t.Errorf("Unexpected href value: got %s, want /about", attr)
	}

	label := aboutLink.Text()
	if label != "About" {
		t.Errorf("Unexpected label value: got %s, want About", label)
	}
}
