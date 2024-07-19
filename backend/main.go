package main

import (
	"backend/cmd/api-server/backend/api"
	"backend/cmd/api-server/backend/db"
	"context"
	"fmt"
	"os"
)

func main() {
	dbConn, err := db.ConnectDB(context.Background(), "Connection_String")

	if err != nil {
		fmt.Println("Error connecting to DB")
		os.Exit(1)
	}

	server := api.NewAPIServer(":8080", dbConn)

	if err := server.Run(); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}