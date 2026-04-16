package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atrivedi/InferX/api/pkg/handlers"
	"github.com/atrivedi/InferX/api/pkg/models"
)

func main() {
	// Day 3: Initialize a buffered channel (the queue)
	// We use a capacity of 100 for our demonstration.
	queueSize := 100
	inferenceQueue := make(chan models.InferenceRequest, queueSize)

	// Initialize our server struct with the queue
	srv := handlers.NewServer(inferenceQueue)

	// Day 4: Launch our background worker
	// The 'go' keyword turns this call into a background goroutine.
	go srv.StartWorker()

	// Register the handler method from our server instance
	http.HandleFunc("/infer", srv.InferHandler)

	port := ":8080"
	fmt.Printf("InferX Server starting on %s...\n", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
