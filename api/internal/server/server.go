package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/atrivedi/InferX/api/pkg/models"
)

// List of supported models (for basic validation)
var supportedModels = map[string]bool{
	"gpt-4":   true,
	"llama-3": true,
	"mistral": true,
}

// Server holds dependencies for our HTTP handlers and inference engine.
type Server struct {
	Queue          chan models.InferenceRequest
	InferenceDelay time.Duration
	BatchSize      int
	BatchTimeout   time.Duration
}

// NewServer creates a new Server instance with the provided dependencies.
func NewServer(queue chan models.InferenceRequest, delay time.Duration, batchSize int, timeout time.Duration) *Server {
	return &Server{
		Queue:          queue,
		InferenceDelay: delay,
		BatchSize:      batchSize,
		BatchTimeout:   timeout,
	}
}

// InferHandler handles the HTTP POST requests for inference.
func (s *Server) InferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.InferenceRequest
	
	// Decoding the JSON body into our struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		slog.Error("error decoding JSON", "error", err)
		return
	}

	// Basic validation
	if !supportedModels[req.Model] {
		slog.Warn("unsupported model requested", "model", req.Model)
	}

	// Pushing the request into our buffered queue
	slog.Info("queuing inference request", "model", req.Model)
	s.Queue <- req
	slog.Info("request successfully queued", "model", req.Model, "depth", len(s.Queue))

	// Prepare the response
	resp := models.InferenceResponse{
		Status:  "queued",
		Message: "Inference request successfully queued for processing",
		Model:   req.Model,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

// StartWorkerPool launches numWorkers background goroutines.
func (s *Server) StartWorkerPool(numWorkers int, wg *sync.WaitGroup) {
	slog.Info("starting worker pool", "num_workers", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go s.runWorker(i, wg)
	}
}

// runWorker is the internal loop for processing tasks in batches.
func (s *Server) runWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	slog.Info("worker started", "worker_id", id, "batch_size", s.BatchSize, "batch_timeout", s.BatchTimeout)

	for {
		// 1. Wait for the FIRST request to arrive (this blocks indefinitely)
		req, ok := <-s.Queue
		if !ok {
			slog.Info("worker shutting down", "worker_id", id)
			return
		}

		// Day 9: We have the first request! Start the batch and the clock.
		batch := models.InferenceBatch{
			Requests: make([]models.InferenceRequest, 0, s.BatchSize),
		}
		batch.Requests = append(batch.Requests, req)

		// Start a timer for the batch timeout
		timer := time.NewTimer(s.BatchTimeout)

		// 2. Collection loop: Grab more till batch is full OR timer fires
		collecting := true
		for collecting && len(batch.Requests) < s.BatchSize {
			select {
			case nextReq, ok := <-s.Queue:
				if !ok {
					collecting = false
					break
				}
				batch.Requests = append(batch.Requests, nextReq)
			case <-timer.C:
				// Timeout! We flush whatever we have.
				collecting = false
			}
		}

		// Be a good citizen: stop the timer if it hasn't fired yet
		if !timer.Stop() {
			// If it already fired and we didn't catch it in the select, drain it
			select {
			case <-timer.C:
			default:
			}
		}

		// 3. Process whatever we collected (might be a partial batch)
		slog.Info("worker processing batch", 
			"worker_id", id, 
			"batch_size", len(batch.Requests))

		// Simulate inference latency
		time.Sleep(s.InferenceDelay)

		slog.Info("worker completed batch", 
			"worker_id", id, 
			"batch_size", len(batch.Requests),
			"queue_depth", len(s.Queue))
	}
}
