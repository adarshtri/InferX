package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/atrivedi/InferX/api/pkg/models"
)

const (
	targetURL   = "http://localhost:8080/infer"
	numRequests = 100
	concurrency = 10
)

func main() {
	fmt.Printf("Starting load test: %d requests total, %d at a time\n", numRequests, concurrency)
	
	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan bool, numRequests)

	// Semaphore to control concurrency
	sem := make(chan struct{}, concurrency)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			sem <- struct{}{} // Acquire slot
			defer func() { <-sem }() // Release slot

			success := sendRequest(id)
			results <- success
		}(i)
	}

	// Close results channel once all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	successCount := 0
	for res := range results {
		if res {
			successCount++
		}
	}

	duration := time.Since(start)
	
	fmt.Println("\n--- Load Test Results ---")
	fmt.Printf("Total Requests: %d\n", numRequests)
	fmt.Printf("Successes:      %d\n", successCount)
	fmt.Printf("Failures:       %d\n", numRequests-successCount)
	fmt.Printf("Total Duration: %v\n", duration)
	fmt.Printf("Requests/sec:   %.2f\n", float64(numRequests)/duration.Seconds())
}

func sendRequest(id int) bool {
	reqBody := models.InferenceRequest{
		Model:  "gpt-4",
		Prompt: fmt.Sprintf("Load test prompt %d", id),
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("[Request %d] Encoding error: %v", id, err)
		return false
	}

	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		// Log error but don't stop the test
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusAccepted
}
