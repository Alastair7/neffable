package main

import (
	"log"
	"net/http"

	routes "github.com/Alastair7/backend/internal/app"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", routes.BuildRoutes()))
}