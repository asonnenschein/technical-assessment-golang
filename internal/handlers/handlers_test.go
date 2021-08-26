package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/asonnenschein/technical-assessment-golang/internal/handlers"
	"github.com/asonnenschein/technical-assessment-golang/internal/models"
	"github.com/gorilla/mux"
)

const (
	originStr string = "https://maps.wikimedia.org/"
	testPath  string = "/osm-intl/1/0/0.png"
)

func TestWildCardHandler(t *testing.T) {
	originUrl, err := url.Parse(originStr)
	if err != nil {
		t.Fatal(err)
	}

	meta := models.Meta{}
	meta.SetOriginUrl(*originUrl)

	req, err := http.NewRequest("GET", testPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/{wildcard:[a-zA-Z0-9/.-]+}", handlers.WildCard(meta))
	router.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf(
			"WildCard handler returned wrong HTTP status code.  Want %v Got %v.",
			status,
			http.StatusOK,
		)
	}

	if contentType := resp.Header().Get("Content-Type"); contentType != "image/png" {
		t.Errorf(
			"WildCard handler returned wrong Content-Type header.  Want %v Got %v.",
			"image/png",
			contentType,
		)
	}
}
