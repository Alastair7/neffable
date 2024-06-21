package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendEmotion(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := gin.Default()
	
	t.Run("Sends a Ping message and returns Pong!", func(t *testing.T) {
		testRouter.GET("api/test/ping", Ping)
		
		req, err := http.NewRequest(http.MethodGet,"/api/test/ping", nil)

		if err != nil {
			t.Fatalf("Couldn't create a request: %v\n", err)
		}

		w := httptest.NewRecorder()

		testRouter.ServeHTTP(w, req)

		decodedResponse := &TestResponse{}

		response := httptest.NewRecorder()
		errDecode := json.NewDecoder(w.Body).Decode(decodedResponse)

		if errDecode != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		if response.Code != http.StatusOK {
			t.Fatalf("Expected to get status %d but instead got %d\n",http.StatusOK, response.Code)
		}

		if decodedResponse.Message != "Pong!" {
			t.Fatalf("Expected %s but instead got %s\n", "Pong!", decodedResponse.Message)
		}

	})
} 