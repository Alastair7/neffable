package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	
	res := httptest.NewRecorder()
	
	req, err := http.NewRequest("POST", "/api/ping", nil)
	if err != nil {
		t.Fatalf("error creating a request: %q", err)
	}

	handler := http.HandlerFunc(PingHandler)
	handler.ServeHTTP(res, req)
	
	got := res.Body.String()
	want := "Pong!"

	if got != want {
		t.Fatalf("expected %q but got %q", want, got)
	}
}