package handlers

import (
	"log"
	"time"
)

// StartWorkerPool launches numWorkers background goroutines, 
// each pulling from the shared Queue.
func (s *Server) StartWorkerPool(numWorkers int) {
	log.Printf("Day 5: Starting worker pool with %d workers...", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Launch each worker in its own goroutine
		go s.runWorker(i)
	}
}

// runWorker is the internal loop that a single worker goroutine executes.
func (s *Server) runWorker(id int) {
	log.Printf("[Worker %d] Started, waiting for tasks...", id)

	for req := range s.Queue {
		log.Printf("[Worker %d] Starting processing for model: %s", id, req.Model)

		// Simulate inference latency
		time.Sleep(500 * time.Millisecond)

		log.Printf("[Worker %d] Completed processing for model: %s", id, req.Model)
		log.Printf("[Worker %d] Task finished. Current queue depth: %d", id, len(s.Queue))
	}

	log.Printf("[Worker %d] Shutting down.", id)
}
