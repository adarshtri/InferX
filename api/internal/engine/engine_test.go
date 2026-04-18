package engine

import "testing"

func TestProcessBatch(t *testing.T) {
	err := ProcessBatch(10)
	if err != nil {
		t.Fatalf("Expected no error when calling C++ engine, got: %v", err)
	}
}
