package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"api/config"
	"api/internal/infrastructure/persistence/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func NewClient(ctx context.Context, cfg config.DatabaseConfig) (*ent.Client, error) {
	db, err := sql.Open(dialect.Postgres, cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))
	if cfg.AutoMigrate {
		if err := client.Schema.Create(ctx); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("run postgres migration: %w", err)
		}
	}

	return client, nil
}
