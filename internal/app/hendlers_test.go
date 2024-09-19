package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// type resource struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }



func TestNotExistingLink(t *testing.T) {
    // Create a new Chi router
    r := chi.NewRouter()

    // Define the GET route
    r.Get("/{shorturl}", getURL)

    // Create a test request
    req, err := http.NewRequest("GET", "/NY34sq", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a test response recorder
    w := httptest.NewRecorder()

    // Serve the request
    r.ServeHTTP(w, req)

    // Check the response status code
    if w.Code != http.StatusBadRequest {
        t.Errorf("Expected status code 400, got %d", w.Code)
    }

    // Check the response body
}