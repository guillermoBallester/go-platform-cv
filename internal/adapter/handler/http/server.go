package http

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/config"
	"log"
	"net/http"
	"time"
)

// Server represents an HTTP server instance with read, write, and idle timeouts
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new server instance with the provided configuration and request handler
func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Server.Address(),
			Handler:      handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  120 * time.Second,
		},
	}
}

// Run starts the server and listens for incoming requests on the specified address
func (s *Server) Run() error {
	log.Printf("Server running on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server with a specified context and timeout of 5 seconds
func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
