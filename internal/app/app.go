package app

import (
	"context"
	"fmt"
	"github.com/guillermoBallester/go-platform-cv/internal/adapter/storage/postgres"
	"github.com/guillermoBallester/go-platform-cv/internal/config"
	"github.com/guillermoBallester/go-platform-cv/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Cfg       *config.Config
	CvService *service.CVService
	SeedSvc   *service.SeedService
	DB        *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	dbPool, err := initDB(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("db init: %w", err)
	}

	repos := postgres.NewRepositories(dbPool)
	cvSvc := service.NewCVService(*repos)
	seedSvc := service.NewSeedService(*repos)

	return &App{
		Cfg:       cfg,
		CvService: cvSvc,
		SeedSvc:   seedSvc,
		DB:        dbPool,
	}, nil
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
