package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ctrixcode/go-chi-postgres/internal/server"
	"github.com/ctrixcode/go-chi-postgres/pkg/logger"
)

func main() {
	// Setup structured logging
	logger.Init()

	s := server.NewServer()

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				slog.Error("graceful shutdown timed out.. forcing exit.")
				os.Exit(1)
			}
		}()

		// Trigger graceful shutdown
		err := s.GetHTTPServer().Shutdown(shutdownCtx)
		if err != nil {
			slog.Error("server shutdown error", "error", err)
			os.Exit(1)
		}
		serverStopCtx()
	}()

	// Run the server
	slog.Info("server starting", "port", s.GetHTTPServer().Addr)
	err := s.Start()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	// Close database connection
	if err := s.Shutdown(); err != nil {
		slog.Error("database shutdown error", "error", err)
		os.Exit(1)
	}

	slog.Info("server exited properly")
}
