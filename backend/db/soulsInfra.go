package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Soul struct {
	ID uuid.UUID `json:"id"`
	DisplayName string `json:"displayName"`
	HasConnection bool `json:"hasConnection"`
	CreatedAt time.Time `json:"createdAt"`
}

func (pg *Postgres) Save(soul *Soul, ctx context.Context) *Soul {
	sql := `
	INSERT INTO souls (display_name)
	VALUES ($1)
	RETURNING id, display_name, has_connection
	`
	err := pg.db.QueryRow(ctx, sql, soul.DisplayName).Scan(
		&soul.ID, 
		&soul.DisplayName,
		&soul.HasConnection)

	if err != nil {
		panic(err)
	}

	return soul
}

func (pg *Postgres) GetSoulByID(ID uuid.UUID, ctx context.Context) (*Soul, error) {
	result := &Soul{}

	sql := `
	SELECT id, display_name, has_connection, created_at FROM souls
	WHERE id = $1
	`

	err := pg.db.QueryRow(ctx, sql, ID).Scan(
		&result.ID,
		&result.DisplayName,
		&result.HasConnection,
		&result.CreatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("soul with ID %s not found", ID)
		}
		return nil, fmt.Errorf("failed to query Soul by ID. Details: %v", err)
	}

	return result, nil
}