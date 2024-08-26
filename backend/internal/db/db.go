package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseClient struct {
	Address string
}

var db *pgxpool.Pool

func(d *DatabaseClient) InitDB() error {
	var err error

	db, err = pgxpool.Connect(context.Background(), d.Address)

	defer db.Close()
	
	if err != nil {
		return err
	}

	return db.Ping(context.Background()) 
}