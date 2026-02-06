package main

import (
	"context"
	_ "embed"
	"github.com/guillermoBallester/go-platform-cv/internal/app"
	"log"

	"github.com/guillermoBallester/go-platform-cv/internal/adapter/handler/http"
	"github.com/guillermoBallester/go-platform-cv/internal/config"
	"github.com/guillermoBallester/go-platform-cv/sql"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	a, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatal("Failed to initialize app: ", err)
	}
	defer a.DB.Close()

	if err := sql.RunMigrations(a.DB); err != nil {
		log.Fatal("Cannot migrate to DB: ", err)
	}

	if cfg.App.SeedData {
		if err := a.SeedSvc.Run(ctx); err != nil {
			log.Printf("SeedSvc warning: %v", err)
		}
	}

	router := http.NewRouter(a.CvService)
	log.Printf("Server initiated at http://localhost%s", cfg.Server.Address())
	if err := router.Run(cfg.Server.Address()); err != nil {
		log.Fatal(err)
	}
}
