package models

import "time"

// InferenceRequest represents the incoming JSON payload for an inference task.
type InferenceRequest struct {
	Model     string    `json:"model"`
	Prompt    string    `json:"prompt"`
	CreatedAt time.Time `json:"-"`
}

// InferenceResponse represents the JSON payload sent back to the client.
type InferenceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Model   string `json:"model,omitempty"`
}

// InferenceBatch represents a group of requests to be processed together.
type InferenceBatch struct {
	Requests []InferenceRequest
}
