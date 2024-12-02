package main

import (
	"fmt"
	"testing"
)

// Test01 tests the CalculateDistance function.
func TestDistance(t *testing.T) {
	result, err := CalculateDistance("test1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 11

	if result != expected {
		t.Errorf("CalculateDistance(test1.txt) = %d; want %d", result, expected)
	}
}

func TestSimilarity(t *testing.T) {
	result, err := CalculateSimilarity("test1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 31

	if result != expected {
		t.Errorf("CalculateSimilarity(test1.txt) = %d; want %d", result, expected)
	}
}
