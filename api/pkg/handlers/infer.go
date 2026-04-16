package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atrivedi/InferX/api/pkg/models"
)

// List of supported models (for basic validation)
var supportedModels = map[string]bool{
	"gpt-4":   true,
	"llama-3": true,
	"mistral": true,
}

// Server holds dependencies for our HTTP handlers, such as the task queue.
type Server struct {
	Queue chan models.InferenceRequest
}

// NewServer creates a new Server instance with the provided queue.
func NewServer(queue chan models.InferenceRequest) *Server {
	return &Server{
		Queue: queue,
	}
}

func (s *Server) InferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.InferenceRequest
	
	// Phase 2: Decoding the JSON body into our struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	// Basic validation
	if !supportedModels[req.Model] {
		log.Printf("Warning: received request for unsupported model: %s", req.Model)
	}

	// Phase 3: Pushing the request into our buffered queue
	// Note: This will block if the queue is full (capacity reached)
	log.Printf("Queuing inference request for model: %s", req.Model)
	s.Queue <- req
	log.Printf("Request successfully queued. Current queue depth: %d", len(s.Queue))

	// Prepare the response
	resp := models.InferenceResponse{
		Status:  "queued",
		Message: "Inference request successfully queued for processing",
		Model:   req.Model,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // 202 Accepted is more semantically correct for async work
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
