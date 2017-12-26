package main_test

import (
	"testing"
	"net/http"
)

func TestGetOriginalURL(t *testing.T) {
	// make a dummy reques
	response, err := http.Get("http://localhost:8000/v1/short/1")
 
    if http.StatusOK != response.StatusCode {
    	t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
    }

    if err != nil {
    	t.Errorf("Encountered an error:", err)
    }
}
