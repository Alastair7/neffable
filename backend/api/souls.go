package api

import (
	"log"
	"neffable/backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestSoul struct {
	DisplayName string `json:"displayName"`
}

type ResponseSoul struct {
	ID uuid.UUID `json:"id"`
	DisplayName string `json:"displayName"`
	HasConnection bool `json:"hasConnection"`
}

type BaseErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func MapToSoulFrom(requestSoul RequestSoul) *db.Soul {
	dbSoul := &db.Soul{DisplayName: requestSoul.DisplayName}

	return dbSoul
}

func MapToSoulResponsetFrom(soul *db.Soul) *ResponseSoul {
	responseSoul := &ResponseSoul{
		ID: soul.ID, 
		DisplayName: soul.DisplayName,
		HasConnection: soul.HasConnection }

		return responseSoul
}

func(s *APIServer) CreateSoul(ctx *gin.Context) {
	var requestSoul RequestSoul

	if err := ctx.ShouldBindJSON(&requestSoul); err != nil {
		errorResponse := &BaseErrorResponse{
			Message: "An error occurred binding the request. Some data may be incorrect.",
			Details: err.Error(),
		}

		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	log.Print(&requestSoul)
	soul := MapToSoulFrom(requestSoul)

	createdSoul := s.db.Save(soul, ctx)

	response := MapToSoulResponsetFrom(createdSoul)

	ctx.JSON(http.StatusCreated, response);
}

func (s *APIServer) GetSoulByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		errorResponse := &BaseErrorResponse{
			Message: "Invalid message format.",
			Details: err.Error(),
		}

		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	soul, err := s.db.GetSoulByID(id, ctx)

	if err != nil {
		errorResponse := &BaseErrorResponse{
			Message: "Soul not found.",
			Details: err.Error(),
		}

		ctx.JSON(http.StatusNotFound, errorResponse)
		return
	}

	response := MapToSoulResponsetFrom(soul)

	ctx.JSON(http.StatusOK, response)
}