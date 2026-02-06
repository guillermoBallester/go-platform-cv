package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	web "github.com/guillermoBallester/go-platform-cv/internal/adapter/handler/http"
	"github.com/guillermoBallester/go-platform-cv/internal/app"
	"github.com/guillermoBallester/go-platform-cv/internal/config"
	"github.com/guillermoBallester/go-platform-cv/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Startup error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load: %w", err)
	}

	a, err := app.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("app init: %w", err)
	}

	defer func() {
		log.Println("Closing DB connection pool...")
		a.DB.Close()
	}()

	if err := sql.RunMigrations(a.DB); err != nil {
		return fmt.Errorf("migrations: %w", err)
	}

	if cfg.App.SeedData {
		if err := a.SeedSvc.Run(ctx); err != nil {
			log.Printf("SeedSvc warning: %v", err)
		}
	}

	router := web.NewRouter(a.CvService)
	srv := web.NewServer(cfg, router)
	go func() {
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server runtime error: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received...")

	if err := srv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("server shutdown: %w", err)
	}

	log.Println("Server exited successfully")
	return nil
}
