package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce sync.Once
)

func ConnectDB(ctx context.Context, dbUrl string) (*Postgres, error) {
	var err error

	pgOnce.Do(func() {
		db, dbErr := pgxpool.New(ctx, dbUrl)

		if dbErr != nil {
			err = fmt.Errorf("unable to create connection pool: %w", dbErr)
			return
		}

		pgInstance = &Postgres{db: db}
	})

	return pgInstance, err
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}

