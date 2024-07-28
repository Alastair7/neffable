package api

import (
	"crypto/rand"
	"io"
	"neffable/backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestSoulConnection struct {
	FirstSoul       uuid.UUID `json:"firstSoul"`
}

type ResponseSoulConnection struct {
	ID              uuid.UUID `json:"id"`
	FirstSoul       uuid.UUID `json:"firstSoul"`
	SecondSoul      uuid.UUID `json:"secondSoul"`
	ConnectionCode  string    `json:"connectionCode"`
}

func MapToSoulConnectionFrom(RequestSoulConnection *RequestSoulConnection) *db.SoulConnection {
	result := &db.SoulConnection{
		FirstSoul: RequestSoulConnection.FirstSoul,
		ConnectionCode: GenerateConnectionCode(),
	}

	return result
}
func MapToSoulConnectionResponseFrom(soulConnection *db.SoulConnection) *ResponseSoulConnection {
	result := &ResponseSoulConnection{
		ID:  soulConnection.ID,
		FirstSoul: soulConnection.FirstSoul,
		SecondSoul: soulConnection.SecondSoul,
		ConnectionCode: soulConnection.ConnectionCode,
	}

	return result
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateConnectionCode() string {
	byteArray := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, byteArray, 6)

	if n != 6 {
		panic(err)
	}

	for i := 0; i < len(byteArray); i++ {
		byteArray[i] = table[int(byteArray[i])%len(table)]
	}

	return string(byteArray)
}
func(s *APIServer) CreateSoulConnection(ctx *gin.Context) {

	var requestData RequestSoulConnection

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		errorResponse := &BaseErrorResponse{
			Message: "An error occurred binding the request. Some data may be incorrect.",
			Details: err.Error(),
		}

		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	soulConnection := MapToSoulConnectionFrom(&requestData)
	createdSoulConnection := s.db.CreateSoulConnection(soulConnection, ctx)

	response := MapToSoulConnectionResponseFrom(createdSoulConnection)

	ctx.JSON(http.StatusCreated, response)
}

func(s *APIServer) ConnectWithSoul(ctx *gin.Context) {
	var request struct {
		SecondSoul uuid.UUID `json:"secondSoul" binding:"required"`
		ConnectionCode string `json:"connectionCode" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errorResponse := &BaseErrorResponse{
			Message: "An error occurred binding the request. Some data may be incorrect.",
			Details: err.Error(),
		}

		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	s.db.UpdateSoulConnection(request.ConnectionCode, request.SecondSoul, ctx)

	ctx.Status(http.StatusOK)
}