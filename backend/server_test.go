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

		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, req)

		decodedResponse := &TestResponse{}

		errDecode := json.NewDecoder(response.Body).Decode(decodedResponse)

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

	t.Run("Returns a mock love emotion", func(t *testing.T) {
		testRouter.POST("/api/emotions", SendEmotion)

		req,err := http.NewRequest(http.MethodPost, "/api/emotions", nil)

		if err != nil {
			t.Fatalf("Couldn't create a request: %v\n", err)
		}

		responseW := httptest.NewRecorder()

		testRouter.ServeHTTP(responseW, req)

		response := &BaseEmotionResponse{}
		errDecode := json.NewDecoder(responseW.Body).Decode(response)

		if errDecode != nil {
			t.Fatalf("Error decoding the response body %s\n", errDecode)
		}

		if responseW.Code != http.StatusOK {
			t.Fatalf("Expected result code %d but instead got %d", http.StatusOK, responseW.Code)
		}

		if response.Emotion != "love" {
			t.Fatalf("Expected %s emotion but instead got %s", "love", response.Emotion)
		}


	})
} 