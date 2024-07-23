package main

import (
	"context"
	"fmt"
	"log"
	"neffable/backend/api"
	"neffable/backend/db"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envPath := os.Getenv("ENV_PATH")

	if envPath == "" {
		envPath = "/neffable/.env"
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	dbConn, err := db.ConnectDB(context.Background(), connectionString)

	if err != nil {
		fmt.Println("Error connecting to DB")
		os.Exit(1)
	}

	server := api.NewAPIServer(":8080", dbConn)

	if err := server.Run(); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}