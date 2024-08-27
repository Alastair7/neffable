package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseProvider interface {
	InitDB() error
}

type PostgresClient struct {
	Address string
	Db *pgxpool.Pool
}

func (pgc *PostgresClient) InitDB() error {
	var err error
	
	pgc.Db, err = pgxpool.Connect(context.Background(), pgc.Address)

	defer pgc.Db.Close()
	
	if err != nil {
		return err
	}

	return pgc.Db.Ping(context.Background()) 
}

func InitDB(provider DatabaseProvider) error {
	if err := provider.InitDB(); err != nil {
		return err
	}

	return nil
}