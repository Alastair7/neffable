package main

import (
	"log"
	"net/http"
	"os"

	routes "github.com/Alastair7/backend/internal/app"
	"github.com/Alastair7/backend/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")
	db := db.DatabaseClient{Address: os.Getenv("DB_CONNECTION_STRING")}
	err := db.InitDB()

	if err != nil {
		log.Fatalf("error initiating the DB %s: ", err)
	}

	log.Fatal(http.ListenAndServe(":5000", routes.BuildRoutes()))
}