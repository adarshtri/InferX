package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/atrivedi/InferX/api/pkg/models"
)

const (
	targetURLDefault = "http://localhost:8080/infer"
)

type TestResult struct {
	ID         int
	StatusCode int
	Latency    time.Duration
	Error      error
}

func main() {
	numRequests := flag.Int("n", 200, "Number of requests to send")
	concurrency := flag.Int("c", 20, "Number of concurrent requests")
	targetURL := flag.String("u", targetURLDefault, "Target URL")
	flag.Parse()

	fmt.Printf("Starting load test: %d requests total, %d at a time on %s\n", *numRequests, *concurrency, *targetURL)
	
	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan TestResult, *numRequests)

	// Semaphore to control concurrency
	sem := make(chan struct{}, *concurrency)

	for i := 0; i < *numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			sem <- struct{}{} // Acquire slot
			defer func() { <-sem }() // Release slot

			results <- sendRequest(id, *targetURL)
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
	var latencies []time.Duration

	for res := range results {
		if res.Error != nil {
			errorCount++
		} else {
			latencies = append(latencies, res.Latency)
			if res.StatusCode == http.StatusAccepted {
				successCount++
			} else if res.StatusCode == http.StatusTooManyRequests {
				rejectedCount++
			} else {
				errorCount++
			}
		}
	}

	duration := time.Since(start)
	
	// Calculate percentiles
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	var p50, p95, p99 time.Duration
	if len(latencies) > 0 {
		p50 = latencies[len(latencies)*50/100]
		p95 = latencies[len(latencies)*95/100]
		p99 = latencies[len(latencies)*99/100]
	}

	fmt.Println("\n--- Load Test Results ---")
	fmt.Printf("Total Requests:   %d\n", *numRequests)
	fmt.Printf("Successes (202):  %d\n", successCount)
	fmt.Printf("Rejections (429): %d\n", rejectedCount)
	fmt.Printf("Hard Failures:    %d\n", errorCount)
	fmt.Printf("Total Duration:   %v\n", duration)
	fmt.Printf("Requests/sec:     %.2f\n", float64(*numRequests)/duration.Seconds())
	fmt.Printf("P50 Latency:      %v\n", p50)
	fmt.Printf("P95 Latency:      %v\n", p95)
	fmt.Printf("P99 Latency:      %v\n", p99)
}

func sendRequest(id int, url string) TestResult {
	reqBody := models.InferenceRequest{
		Model:  "gpt-4",
		Prompt: fmt.Sprintf("Load test prompt %d", id),
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return TestResult{ID: id, Error: err}
	}

	start := time.Now()
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	latency := time.Since(start)

	if err != nil {
		return TestResult{ID: id, Error: err, Latency: latency}
	}
	defer resp.Body.Close()

	return TestResult{ID: id, StatusCode: resp.StatusCode, Latency: latency}
}

