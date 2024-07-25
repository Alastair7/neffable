package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SoulConnection struct {
	ID              uuid.UUID `json:"id"`
	FirstSoul       uuid.UUID `json:"first_soul"`
	SecondSoul      uuid.UUID `json:"second_soul"`
	ConnectionCode  string    `json:"connection_code"`
	ConnectionDate  time.Time `json:"connection_date"`
}

func CreateSoulConnection(context *gin.Context) {
	response := &SoulConnection{
		ID: uuid.New(),
		FirstSoul: uuid.New(),
		SecondSoul: uuid.New(),
		ConnectionCode: "1234",
		ConnectionDate: time.Now(),
	}

	context.JSON(http.StatusCreated, response)
}

func ConnectWithSoul() {}