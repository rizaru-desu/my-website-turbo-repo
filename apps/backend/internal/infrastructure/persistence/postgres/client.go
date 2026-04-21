package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
		if err := prepareLegacyProfileContentJSONColumns(ctx, db); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("prepare legacy profile content json columns: %w", err)
		}

		if err := client.Schema.Create(ctx); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("run postgres migration: %w", err)
		}
	}

	return client, nil
}

func prepareLegacyProfileContentJSONColumns(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	var tableExists bool
	if err := tx.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = current_schema()
				AND table_name = 'profileContent'
		)
	`).Scan(&tableExists); err != nil {
		return fmt.Errorf("check profileContent table: %w", err)
	}

	if !tableExists {
		return tx.Commit()
	}

	if _, err := tx.ExecContext(ctx, `
		CREATE OR REPLACE FUNCTION pg_temp.__portfolio_safe_jsonb(value text, fallback jsonb, expected_type text)
		RETURNS jsonb
		LANGUAGE plpgsql
		AS $$
		DECLARE
			parsed jsonb;
		BEGIN
			IF value IS NULL OR btrim(value) = '' THEN
				RETURN fallback;
			END IF;

			parsed := value::jsonb;
			IF expected_type <> '' AND jsonb_typeof(parsed) <> expected_type THEN
				RETURN fallback;
			END IF;

			RETURN parsed;
		EXCEPTION WHEN others THEN
			RETURN fallback;
		END;
		$$;
	`); err != nil {
		return fmt.Errorf("create safe jsonb helper: %w", err)
	}

	repairs := []struct {
		column       string
		fallback     string
		expectedType string
	}{
		{column: "socialLinks", fallback: "{}", expectedType: "object"},
		{column: "stats", fallback: "[]", expectedType: "array"},
	}

	for _, repair := range repairs {
		if err := normalizeProfileContentJSONColumn(ctx, tx, repair.column, repair.fallback, repair.expectedType); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func normalizeProfileContentJSONColumn(ctx context.Context, tx *sql.Tx, column, fallback, expectedType string) error {
	var udtName string
	if err := tx.QueryRowContext(ctx, `
		SELECT udt_name
		FROM information_schema.columns
		WHERE table_schema = current_schema()
			AND table_name = 'profileContent'
			AND column_name = $1
	`, column).Scan(&udtName); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		return fmt.Errorf("check profileContent.%s column: %w", column, err)
	}

	if udtName == "jsonb" {
		return nil
	}

	quotedColumn := quotePostgresIdentifier(column)
	if _, err := tx.ExecContext(ctx, fmt.Sprintf(`
		ALTER TABLE "profileContent"
		ALTER COLUMN %s DROP DEFAULT
	`, quotedColumn)); err != nil {
		return fmt.Errorf("drop default for profileContent.%s: %w", column, err)
	}

	if _, err := tx.ExecContext(ctx, fmt.Sprintf(`
		ALTER TABLE "profileContent"
		ALTER COLUMN %s TYPE jsonb
		USING pg_temp.__portfolio_safe_jsonb(%s::text, $1::jsonb, $2)
	`, quotedColumn, quotedColumn), fallback, expectedType); err != nil {
		return fmt.Errorf("normalize profileContent.%s to jsonb: %w", column, err)
	}

	return nil
}

func quotePostgresIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}
