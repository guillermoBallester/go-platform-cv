package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config holds all configuration for the application.
// It is immutable after loading.
type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

// AppConfig holds application-level configuration.
type AppConfig struct {
	Env      string `env:"APP_ENV" envDefault:"production"`
	SeedData bool   `env:"SEED_DATA" envDefault:"true"`
}

// IsDevelopment returns true if running in development mode.
func (a AppConfig) IsDevelopment() bool {
	return a.Env == "development"
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Port         string        `env:"PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"30s"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"60s"`
}

// Address returns the server address in the format ":port".
func (s ServerConfig) Address() string {
	return ":" + s.Port
}

// DatabaseConfig holds database connection configuration.
type DatabaseConfig struct {
	URL             string        `env:"DATABASE_URL"`
	Host            string        `env:"DB_HOST"`
	Port            int           `env:"DB_PORT" envDefault:"5432"`
	User            string        `env:"DB_USER"`
	Password        string        `env:"DB_PASSWORD"`
	DBName          string        `env:"DB_NAME"`
	SSLMode         string        `env:"DB_SSLMODE" envDefault:"disable"`
	MaxOpenConns    int           `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns    int           `env:"DB_MAX_IDLE_CONNS" envDefault:"5"`
	ConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME" envDefault:"5m"`
}

// ConnectionString returns the database connection string.
// If DATABASE_URL is set, it returns that directly.
// Otherwise, it builds the connection string from individual components.
func (d DatabaseConfig) ConnectionString() string {
	if d.URL != "" {
		return d.URL
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.DBName, d.SSLMode,
	)
}

// hasComponents returns true if individual database components are configured.
func (d DatabaseConfig) hasComponents() bool {
	return d.Host != "" && d.User != "" && d.DBName != ""
}

// Load parses environment variables and returns a validated Config.
// It fails fast if required configuration is missing.
func Load() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return cfg, nil
}

// Validate checks that all required configuration is present.
// In development mode, local defaults are allowed.
// In production mode, database configuration is strictly required.
func (c *Config) Validate() error {
	if c.App.IsDevelopment() {
		return c.validateDevelopment()
	}
	return c.validateProduction()
}

func (c *Config) validateDevelopment() error {
	// In development mode, apply local defaults if no database config is provided
	if c.Database.URL == "" && !c.Database.hasComponents() {
		c.Database.URL = "postgres://postgres:postgres@localhost:5432/gocv?sslmode=disable"
	}
	return nil
}

func (c *Config) validateProduction() error {
	if c.Database.URL != "" {
		return nil
	}

	if !c.Database.hasComponents() {
		return errors.New("database configuration required: set DATABASE_URL or (DB_HOST, DB_USER, DB_NAME)")
	}

	var missing []string
	if c.Database.Host == "" {
		missing = append(missing, "DB_HOST")
	}
	if c.Database.User == "" {
		missing = append(missing, "DB_USER")
	}
	if c.Database.DBName == "" {
		missing = append(missing, "DB_NAME")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required database configuration: %v", missing)
	}

	return nil
}
