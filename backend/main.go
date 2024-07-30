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
		envPath = "../.env"
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	 fmt.Print(dbConnectionString)

	dbConn, err := db.ConnectDB(context.Background(), dbConnectionString)

	if err != nil {
		fmt.Println("Error connecting to DB")
		os.Exit(1)
	}

	serverPort := os.Getenv("SERVER_PORT")
	serverAddress := fmt.Sprintf(":%s", serverPort)
	server := api.NewAPIServer(serverAddress, dbConn)

	if err := server.Run(); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}