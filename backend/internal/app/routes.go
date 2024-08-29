package routes

import (
	"net/http"

	middlewares "github.com/Alastair7/backend/internal/middllewares"
	handler "github.com/Alastair7/backend/internal/server"
	"github.com/gorilla/mux"
)

func BuildRoutes() *mux.Router  {
	router := mux.NewRouter()

	router.Handle("/api/ping", middlewares.EnsureValidToken()(
		http.HandlerFunc( handler.PingHandler),
	))

	return router
}