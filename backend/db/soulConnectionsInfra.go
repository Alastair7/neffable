package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SoulConnection struct {
	ID             uuid.UUID
	FirstSoul      uuid.UUID
	SecondSoul     uuid.UUID
	ConnectionDate time.Time
	ConnectionCode string
}

func (pg *Postgres) CreateSoulConnection(soulConnection *SoulConnection, ctx context.Context) *SoulConnection {
	sql := `
	INSERT INTO soul_connections (first_soul, second_soul, connection_code)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	err := pg.db.QueryRow(ctx, sql, 
		soulConnection.FirstSoul, 
		nil,
		soulConnection.ConnectionCode).Scan(
			&soulConnection.ID,
		)

		if err != nil {
			panic(err)
		}

	return soulConnection
}