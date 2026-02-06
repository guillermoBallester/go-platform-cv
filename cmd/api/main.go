package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/guillermoBallester/go-platform-cv/internal/adapter/handler/http"
	"github.com/guillermoBallester/go-platform-cv/internal/adapter/storage/postgres"
	"github.com/guillermoBallester/go-platform-cv/internal/config"
	"github.com/guillermoBallester/go-platform-cv/internal/service"
	"github.com/guillermoBallester/go-platform-cv/sql"
	"github.com/guillermoBallester/go-platform-cv/sql/data"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	// Infra: Database Connection
	dbPool, err := initDB(ctx, cfg)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}
	defer dbPool.Close()

	// Infra: migrations
	if err := sql.RunMigrations(dbPool); err != nil {
		log.Fatal("Cannot migrate to DB: ", err)
	}

	// Dependency Injection
	queries := postgres.New(dbPool)
	skillsRepo := postgres.NewSkillRepository(queries)
	expRepo := postgres.NewExperienceRepository(queries)
	cvSvc := service.NewCVService(skillsRepo, expRepo)

	// Seed data
	if err := cvSvc.SeedSkills(ctx, data.SkillsJSON); err != nil {
		log.Printf("Warning: could not seed skills: %v", err)
	}
	if err := cvSvc.SeedExperiences(ctx, data.ExperiencesJSON); err != nil {
		log.Printf("Warning: could not seed experiences: %v", err)
	}

	// Init server
	router := http.NewRouter(cvSvc)
	log.Printf("Server initiated at http://localhost%s", cfg.Server.Address())
	if err := router.Run(cfg.Server.Address()); err != nil {
		log.Fatal(err)
	}
}

func initDB(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.Database.ConnectionString())
	if err != nil {
		return nil, err
	}

	// Apply connection pool settings from config
	poolConfig.MaxConns = int32(cfg.Database.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.Database.MaxIdleConns)
	poolConfig.MaxConnLifetime = cfg.Database.ConnMaxLifetime

	return pgxpool.NewWithConfig(ctx, poolConfig)
}
