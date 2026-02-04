package sql

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"io/fs"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func setup() error {
	fsys, err := fs.Sub(embedMigrations, "migrations")
	if err != nil {
		return fmt.Errorf("failed to get subdirectory: %w", err)
	}

	goose.SetBaseFS(fsys)

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return errors.Join(fmt.Errorf("failed to set goose dialect to postgres"), err)
	}

	return nil
}

// RunMigrations We might actually have to consider a recover, since goose does panic
func RunMigrations(p *pgxpool.Pool) error {
	return run(stdlib.OpenDBFromPool(p))
}

// Run runs all pending migrations
func run(db *sql.DB) error {
	if err := setup(); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("failed to run up migrations: %w", err)
	}

	return nil
}
