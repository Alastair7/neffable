package api

import (
	"context"
	"neffable/backend/db"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	address string
	db      *db.Postgres
}

func NewAPIServer(address string, db *db.Postgres) *APIServer {
	return &APIServer{
		address: address,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	
	err := s.db.Ping(context.Background())

	if err != nil {
		return err
	}

	router.GET("api/test/ping", s.Ping)
	router.POST("/api/souls", s.CreateSoul)
	router.GET("/api/souls/:id", s.GetSoulByID)
	router.POST("api/soulConnections", s.CreateSoulConnection) 

	return router.Run(s.address)
}