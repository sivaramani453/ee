package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOctocatGists(t *testing.T) {

	req, err := http.NewRequest("GET", "/octocat", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 but got %v", status)
	}
}
