package routes

import (
	handler "github.com/Alastair7/backend/internal/server"
	"github.com/gorilla/mux"
)

func BuildRoutes() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/api/ping", handler.PingHandler)

	return router
}