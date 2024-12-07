package main

import (
	"fmt"
	"testing"
)

func TestStartPosition(t *testing.T) {
	roomMap, err := readInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := findGuard(roomMap)

	expected := Position{x: 6, y: 4}

	if result != expected {
		t.Errorf("Guard not found at correct spot, want %d; got %d", expected, result)
	}
}

func TestInsideBounds(t *testing.T) {
	roomMap, err := readInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	BoundsTests := []struct {
		name     string
		position Position
		isInside bool
	}{
		{name: "Test (1 1)", position: Position{x: 1, y: 1}, isInside: true},
		{name: "Test (2 2)", position: Position{x: 2, y: 2}, isInside: true},
		{name: "Test (0 0)", position: Position{x: 0, y: 0}, isInside: true},
		{name: "Test (-1 0)", position: Position{x: -1, y: 0}, isInside: false},
		{name: "Test (0 -1)", position: Position{x: 0, y: -1}, isInside: false},
		{name: "Test (-1 -1)", position: Position{x: -1, y: -1}, isInside: false},
		{name: "Test (100 100)", position: Position{x: 100, y: 100}, isInside: false},
		{name: "Test (9 9)", position: Position{x: 9, y: 9}, isInside: true},
		{name: "Test (10 10)", position: Position{x: 10, y: 10}, isInside: false},
		{name: "Test (0 10)", position: Position{x: 0, y: 10}, isInside: false},
	}

	for _, tt := range BoundsTests {
		t.Run(tt.name, func(t *testing.T) {
			got := insideBounds(tt.position, len(roomMap[0]), len(roomMap))
			if got != tt.isInside {
				t.Errorf("Position not correctly found to be inside, want %t; got %t", tt.isInside, got)
			}
		})
	}

}

func TestDirectionDetection(t *testing.T) {
	DirectionTests := []struct {
		name      string
		symbol    rune
		direction GuardDirection
	}{
		{name: "Up", symbol: rune('^'), direction: Up},
		{name: "Right", symbol: rune('>'), direction: Right},
		{name: "Down", symbol: rune('v'), direction: Down},
		{name: "Left", symbol: rune('<'), direction: Left},
	}

	for _, tt := range DirectionTests {
		t.Run(tt.name, func(t *testing.T) {
			got := findGuardDirection(tt.symbol)
			if got != tt.direction {
				t.Errorf("Direction is wrong, want %d, got %d", tt.direction, got)
			}
		})
	}
}

func TestMovingOneSpace(t *testing.T) {
	roomMap, err := readInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currentPosition := findGuard(roomMap)
	moveGuard(roomMap, currentPosition)
}

func TestPart1(t *testing.T) {
	got, err := solvePart1("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	want := 41

	if got != want {
		t.Errorf("Part 1 solver wrong, want %d; got %d", want, got)
	}

}

func TestPart2(t *testing.T) {
	got, err := solvePart2("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	want := 6

	if got != want {
		t.Errorf("Part 1 solver wrong, want %d; got %d", want, got)
	}

}

func TestEdgeCase1(t *testing.T) {
	roomMap, err := readInput("edgeCase1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	got := isGuardLooping(roomMap, Position{x: 0, y: 0})
	expect := true

	if got != expect {
		t.Errorf("Looping not working on edge case, got %v, want %v", got, expect)
	}
}
