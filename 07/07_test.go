package main

import (
	"fmt"
	"testing"
)

func TestSmallExample(t *testing.T) {
	got, err := solvePart1("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	want := 3749

	if got != want {
		t.Errorf("Part 1 not working on test example, got %d, want %d", got, want)
	}
}
