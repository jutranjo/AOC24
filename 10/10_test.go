package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	got, err := solvePart1("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	want := 36

	if got != want {
		t.Errorf("Part 1 solver wrong, want %d; got %d", want, got)
	}

}
