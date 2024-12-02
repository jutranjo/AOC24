package main

import (
	"fmt"
	"testing"
)

func TestCountSafeReports(t *testing.T) {
	result, err := CountSafeReports("test2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 2

	if result != expected {
		t.Errorf("CountSafeReports(test1.txt) = %d; want %d", result, expected)
	}
}

func TestCountBadSafeReports(t *testing.T) {
	result, err := CountTolerableReports("test2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 4

	if result != expected {
		t.Errorf("CountSafeReports(test1.txt) = %d; want %d", result, expected)
	}
}

func TestTolerableEdgeCases(t *testing.T) {
	result, err := CountTolerableReports("additional_test2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 12

	if result != expected {
		t.Errorf("CountSafeReports(test1.txt) = %d; want %d", result, expected)
	}
}
