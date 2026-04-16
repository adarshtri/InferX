package handlers

import (
	"log"
	"time"
)

// StartWorker begins a loop that consumes requests from the Queue.
// It should be launched as a background goroutine.
func (s *Server) StartWorker() {
	log.Println("Day 4: Background worker started, waiting for tasks...")

	// The 'range' loop over a channel is a common Go pattern.
	// It will wait efficiently for a new value, process it, 
	// and resume waiting until the channel is closed.
	for req := range s.Queue {
		log.Printf("[Worker] Starting processing for model: %s", req.Model)

		// Day 4: Simulate inference compute time (latency)
		// This simulates the actual heavy lifting of an LLM.
		time.Sleep(500 * time.Millisecond)

		log.Printf("[Worker] Completed processing for model: %s. Prompt snippet: '%.20s...'", req.Model, req.Prompt)
		log.Printf("[Worker] Task finished. Current queue depth: %d", len(s.Queue))
	}

	log.Println("Worker shutting down (channel closed).")
}
