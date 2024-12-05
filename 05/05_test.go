package main

import (
	"fmt"
	"testing"
)

func TestIsUpdateCorrect(t *testing.T) {
	rules, updates, err := parseInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := []bool{true, true, true, false, false, false}

	for index, update := range updates {
		result := isUpdateCorrect(update, rules)
		if expected[index] != result {
			t.Errorf("error; got %v; want %v", result, expected)
		}
	}

}

func TestSummerPart1(t *testing.T) {
	result, err := solvePart1("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 143

	if expected != result {
		t.Errorf("Wrong sum of example part1, got %d, want %d", result, expected)
	}

}
