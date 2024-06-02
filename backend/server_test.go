package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSendEmotion(t *testing.T) {

	t.Run("Sends an emotion and returns it", func(t *testing.T) {
		request, _:= http.NewRequest(http.MethodPost, "/emotions", nil)
		response := httptest.NewRecorder()
		
		NeffableServer(response, request)
		
		emotionResponse := decodeResponse(t, response)


		want := BaseEmotionResponse{StatusCode: 200, Message: "SUCCESS", Emotion: "Love"}
		got := emotionResponse

		assertResponseBody(t, want, got)
	})
}

func decodeResponse(t *testing.T, response *httptest.ResponseRecorder) BaseEmotionResponse {
	var emotionResponse BaseEmotionResponse
	err := json.NewDecoder(response.Body).Decode(&emotionResponse)

	if err != nil {
		t.Fatalf("Unable decoding %q. '%v'", response.Body, err)
	}

	return emotionResponse
}

func assertResponseBody(t testing.TB, want,got BaseEmotionResponse) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %#v got %#v", want, got)
	}
} 