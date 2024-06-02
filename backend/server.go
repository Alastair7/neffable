package main

import (
	"encoding/json"
	"net/http"
)

type BaseEmotionResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Emotion    string `json:"emotion"`
}

func NeffableServer(w http.ResponseWriter, r *http.Request) {
	emotionResponse := BaseEmotionResponse{StatusCode: 200, Message: "SUCCESS", Emotion: "Love"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emotionResponse)
}