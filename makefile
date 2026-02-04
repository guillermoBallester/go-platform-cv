# Variables
BINARY_NAME=hablanorsk
DOCKER_COMPOSE=docker compose

.PHONY: help build up down restart logs lint sqlc test clean

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

## --- Docker Commands ---

up: ## Build and start the containers (detached)
	$(DOCKER_COMPOSE) up --build -d

down: ## Stop containers and remove volumes (clean reset)
	$(DOCKER_COMPOSE) down -v

restart: down up ## Full reset: stop, remove volumes, and start again

logs: ## Tail container logs
	$(DOCKER_COMPOSE) logs -f

## --- Development & Quality ---

lint: ## Run golangci-lint
	golangci-lint run ./...

sqlc: ## Generate Go code from SQL queries
	sqlc generate

test: ## Run all tests
	go test -v ./...

tidy: ## Clean up go.mod and format code
	go mod tidy
	go fmt ./...

build: ## Build the local binary (for testing without Docker)
	go build -o bin/$(BINARY_NAME) ./cmd/api/main.go

clean: ## Remove binaries and temp files
	rm -rf bin/