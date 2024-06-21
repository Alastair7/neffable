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

	router.Run()
}

func Ping(context *gin.Context) {
	response := &TestResponse{Message: "Pong!"}

	context.JSON(http.StatusOK, response)
}