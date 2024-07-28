package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestResponse struct {
	Message string `json:"message"`
}

func (APIServer) Ping(context *gin.Context) {
	response := &TestResponse{Message: "Pong!"}

	context.JSON(http.StatusOK, response)
}