package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atrivedi/InferX/api/pkg/handlers"
)

func main() {
	// Register the handler from our custom package
	http.HandleFunc("/infer", handlers.InferHandler)

	port := ":8080"
	fmt.Printf("InferX Server starting on %s...\n", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
