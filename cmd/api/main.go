package main

import (
	"context"
	"fmt"
	"github.com/guillermoBallester/go-platform-cv/internal/adapter/storage/postgres"
	"github.com/guillermoBallester/go-platform-cv/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	// Infra: Database Connection
	dbPool := initDB(ctx)
	defer dbPool.Close()

	// Infra: migrations
	if err := sql.RunMigrations(dbPool); err != nil {
		log.Fatal("Cannot migrate to DB:", err)
	}

	// Dependency Injection
	queries := postgres.New(dbPool)

	skills, err := queries.ListSkills(ctx)
	if err != nil {
		return
	}
	if len(skills) > 0 {
		fmt.Printf("Found %d skills", len(skills))
	}

}

func initDB(ctx context.Context) *pgxpool.Pool {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5432/gocv?sslmode=disable"
	}
	dbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	return dbPool
}
