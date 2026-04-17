package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/atrivedi/InferX/api/internal/server"
	"github.com/atrivedi/InferX/api/pkg/config"
	"github.com/atrivedi/InferX/api/pkg/models"
)

func main() {
	// Initialize structured logging (JSON format for production-readiness)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load centralized configuration
	cfg := config.Load()
	slog.Info("initializing InferX server", 
		"port", cfg.Port, 
		"queue_size", cfg.QueueSize, 
		"worker_count", cfg.WorkerCount,
		"inference_delay_ms", cfg.InferenceDelayMS)

	// Day 7: WaitGroup for tracking background workers
	var wg sync.WaitGroup

	// Initialize the buffered request queue
	inferenceQueue := make(chan models.InferenceRequest, cfg.QueueSize)

	// Initialize our server struct with the dependencies
	inferenceDelay := time.Duration(cfg.InferenceDelayMS) * time.Millisecond
	srv := server.NewServer(inferenceQueue, inferenceDelay)

	// Launch our background worker pool (pass the waitgroup)
	srv.StartWorkerPool(cfg.WorkerCount, &wg)

	// Setup our HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/infer", srv.InferHandler)

	// Manual http.Server for graceful shutdown
	httpServer := &http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}

	// Channel to listen for interrupts (Ctrl+C, termination)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the HTTP server in a goroutine so it doesn't block the main thread
	go func() {
		slog.Info("server starting", "port", cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for the stop signal
	<-stop
	slog.Info("shutting down gracefully...")

	// 1. Stop accepting new HTTP requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}

	// 2. Signal the worker pool to stop (close the queue)
	close(inferenceQueue)

	// 3. Wait for workers to finish the backlog
	slog.Info("waiting for workers to finish current tasks...")
	wg.Wait()

	slog.Info("shutdown complete. Goodbye!")
}
