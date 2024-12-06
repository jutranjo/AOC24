package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const GuardSymbols = "^v<>"
const Obstacle = rune('#')
const Traversed = rune('X')

type GuardDirection int

const (
	Up GuardDirection = iota
	Down
	Left
	Right
)

func (gd GuardDirection) String() string {
	switch gd {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "NOT A DIRECTION"
	}
}

type GuardPosition struct {
	x, y int
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var readMap [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		readMap = append(readMap, []rune(line))
	}

	return readMap, nil
}

func printMap(roomMap [][]rune) {
	for _, row := range roomMap {
		fmt.Printf("%c \n", row)
	}
}

func findGuard(roomMap [][]rune) GuardPosition {
	var guardStartPosition GuardPosition
	for i, roomLine := range roomMap {
		for j, square := range roomLine {
			if strings.ContainsRune(GuardSymbols, square) {
				guardStartPosition = GuardPosition{x: i, y: j}
			}
		}
	}
	return guardStartPosition
}

func insideBounds(currentGuardPosition GuardPosition, roomMaxWidthIndex int, roomMaxHeightIndex int) bool {
	if currentGuardPosition.x < 0 || currentGuardPosition.y < 0 || currentGuardPosition.x > roomMaxWidthIndex-1 || currentGuardPosition.y > roomMaxHeightIndex-1 {
		return false
	}
	return true
}

func findGuardDirection(guardRune rune) GuardDirection {
	switch guardRune {
	case rune('^'):
		return Up
	case rune('v'):
		return Down
	case rune('<'):
		return Left
	case rune('>'):
		return Right
	}
	return 5
}

func findNextSpot

func moveGuard(roomMap [][]rune, currentPosition GuardPosition) {
	direction := findGuardDirection(roomMap[currentPosition.x][currentPosition.y])

	//nextSpace := findNextSpot(roomMap, currentPosition, direction)
	//			 are we out of bounds now? true -> return
	//		is there an obstacle?
	//			-> rotate guard 90 degrees, return
	//		replace guardposition with X, new guard position if . replaced with with <>v^, return
}

func solvePart1(filename string) (int, error) {
	roomMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	printMap(roomMap)
	currentPosition := findGuard(roomMap)
	fmt.Printf("Start position is [%d %d]", currentPosition.x, currentPosition.y)

	roomWidth := len(roomMap[0])
	roomHeight := len(roomMap)

	for insideBounds(currentPosition, roomWidth, roomHeight) {
		moveGuard(roomMap, currentPosition)
	}

	return 0, nil
}
