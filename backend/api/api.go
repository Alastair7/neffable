package api

import (
	"backend/cmd/api-server/backend/db"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	address string
	db      *db.Postgres
}

type BaseEmotionResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Emotion    string `json:"emotion"`
}

type TestResponse struct {
	Message string `json:"message"`
}

func NewAPIServer(address string, db *db.Postgres) *APIServer {
	return &APIServer{
		address: address,
		db: db,
	}
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

func (s *APIServer) Run() error {
	router := gin.Default()

	err := s.db.Ping(context.Background())

	if err != nil {
		return err
	}

	router.GET("api/test/ping", Ping)
	router.POST("api/emotions", SendEmotion) 

	return router.Run()
}