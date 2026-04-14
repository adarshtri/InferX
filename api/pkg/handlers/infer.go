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

func InferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.InferenceRequest
	
	// Phase 2: Decoding the JSON body into our struct
	// We use json.NewDecoder for better performance with HTTP streams
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

	log.Printf("Received inference request for model: %s", req.Model)

	// Prepare the response
	resp := models.InferenceResponse{
		Status:  "received",
		Message: "Inference request successfully parsed",
		Model:   req.Model,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
