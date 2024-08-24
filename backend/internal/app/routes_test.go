package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Alastair7/backend/internal/server"
)

func TestBuildRoutes(t *testing.T) {
	const basePath = "/api"
	t.Run("Ping should return 'Pong!'", func(t *testing.T) {
		path := fmt.Sprintf("%q/ping", basePath)
		req, err := http.NewRequest(http.MethodPost, path, nil)
		if err != nil {
			t.Fatalf("error creating the request")
		}

		res := httptest.NewRecorder()
		
		handlerTest := http.HandlerFunc(handler.PingHandler)
		handlerTest.ServeHTTP(res, req)

		want := "Pong!"
		got := res.Body.String()

		if (got != want) {
			t.Fatalf("expected %q but got %q", want, got)
		}

	})
}