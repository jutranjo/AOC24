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

	expected := GuardPosition{x: 6, y: 4}

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
		position GuardPosition
		isInside bool
	}{
		{name: "Test (1 1)", position: GuardPosition{x: 1, y: 1}, isInside: true},
		{name: "Test (2 2)", position: GuardPosition{x: 2, y: 2}, isInside: true},
		{name: "Test (0 0)", position: GuardPosition{x: 0, y: 0}, isInside: true},
		{name: "Test (-1 0)", position: GuardPosition{x: -1, y: 0}, isInside: false},
		{name: "Test (0 -1)", position: GuardPosition{x: 0, y: -1}, isInside: false},
		{name: "Test (-1 -1)", position: GuardPosition{x: -1, y: -1}, isInside: false},
		{name: "Test (100 100)", position: GuardPosition{x: 100, y: 100}, isInside: false},
		{name: "Test (9 9)", position: GuardPosition{x: 9, y: 9}, isInside: true},
		{name: "Test (10 10)", position: GuardPosition{x: 10, y: 10}, isInside: false},
		{name: "Test (0 10)", position: GuardPosition{x: 0, y: 10}, isInside: false},
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
