package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	apiServer := APIServer{}

	r.GET("/ping", apiServer.Ping)

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response TestResponse

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response body: %V", err)
	}

	expectedMessage := "Pong!"

	if response.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, response.Message)
	}
}