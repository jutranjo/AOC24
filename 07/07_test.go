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

	want := 9 //3749

	if got != want {
		t.Errorf("Part 1 not working on test example, got %d, want %d", got, want)
	}
}

func TestExampleLine2(t *testing.T) {
	listOfNumbers, err := readInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("input is", listOfNumbers[1])
	got := isTrueEquation(listOfNumbers[1])
	want := true

	if got != want {
		t.Errorf("2nd example line not correct, got %v, want %v", got, want)
	}
}
