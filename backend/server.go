package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseEmotionResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Emotion    string `json:"emotion"`
}

type TestResponse struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.GET("api/test/ping", Ping)
	router.POST("api/emotions", SendEmotion) 

	router.Run()
}

func Ping(context *gin.Context) {
	response := &TestResponse{Message: "Pong!"}

	context.JSON(http.StatusOK, response)
}

func SendEmotion(context *gin.Context) {
	response := &BaseEmotionResponse{
		StatusCode: http.StatusOK,
		Message: "Mock Emotion Sent",
		Emotion: "love",
	}

	context.JSON(http.StatusOK, response)
}