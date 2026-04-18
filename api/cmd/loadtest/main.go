package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/atrivedi/InferX/api/pkg/models"
)

const (
	targetURL   = "http://localhost:8080/infer"
	numRequests = 200
	concurrency = 20
)

type TestResult struct {
	ID         int
	StatusCode int
	Error      error
}

func main() {
	fmt.Printf("Starting load test: %d requests total, %d at a time\n", numRequests, concurrency)
	
	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan TestResult, numRequests)

	// Semaphore to control concurrency
	sem := make(chan struct{}, concurrency)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			sem <- struct{}{} // Acquire slot
			defer func() { <-sem }() // Release slot

			results <- sendRequest(id)
		}(i)
	}

	// Close results channel once all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	successCount := 0
	rejectedCount := 0
	errorCount := 0

	for res := range results {
		if res.Error != nil {
			errorCount++
		} else if res.StatusCode == http.StatusAccepted {
			successCount++
		} else if res.StatusCode == http.StatusTooManyRequests {
			rejectedCount++
		} else {
			errorCount++
		}
	}

	duration := time.Since(start)
	
	fmt.Println("\n--- Load Test Results ---")
	fmt.Printf("Total Requests:   %d\n", numRequests)
	fmt.Printf("Successes (202):  %d\n", successCount)
	fmt.Printf("Rejections (429): %d\n", rejectedCount)
	fmt.Printf("Hard Failures:    %d\n", errorCount)
	fmt.Printf("Total Duration:   %v\n", duration)
	fmt.Printf("Requests/sec:     %.2f\n", float64(numRequests)/duration.Seconds())
}

func sendRequest(id int) TestResult {
	reqBody := models.InferenceRequest{
		Model:  "gpt-4",
		Prompt: fmt.Sprintf("Load test prompt %d", id),
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return TestResult{ID: id, Error: err}
	}

	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return TestResult{ID: id, Error: err}
	}
	defer resp.Body.Close()

	return TestResult{ID: id, StatusCode: resp.StatusCode}
}
