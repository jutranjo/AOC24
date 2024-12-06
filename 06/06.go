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

type Position struct {
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
	fmt.Printf("\n")
}

func findGuard(roomMap [][]rune) Position {
	var guardStartPosition Position
	for i, roomLine := range roomMap {
		for j, square := range roomLine {
			if strings.ContainsRune(GuardSymbols, square) {
				guardStartPosition = Position{x: i, y: j}
			}
		}
	}
	return guardStartPosition
}

func insideBounds(currentGuardPosition Position, roomMaxWidthIndex int, roomMaxHeightIndex int) bool {
	return !(currentGuardPosition.x < 0 || currentGuardPosition.y < 0 || currentGuardPosition.x > roomMaxWidthIndex-1 || currentGuardPosition.y > roomMaxHeightIndex-1)
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

func lookAtPosition(roomMap [][]rune, position Position) rune {
	return roomMap[position.x][position.y]
}

func updateSpot(roomMap [][]rune, position Position, newRune rune) {
	roomMap[position.x][position.y] = newRune
}

func moveGuard(roomMap [][]rune, currentPosition Position) Position {
	originalGuardPosition := currentPosition

	direction := findGuardDirection(roomMap[currentPosition.x][currentPosition.y])
	switch direction {
	case Up:
		currentPosition.x -= 1
	case Down:
		currentPosition.x += 1
	case Left:
		currentPosition.y -= 1
	case Right:
		currentPosition.y += 1
	}

	roomWidth := len(roomMap[0])
	roomHeight := len(roomMap)

	if !insideBounds(currentPosition, roomWidth, roomHeight) {
		updateSpot(roomMap, originalGuardPosition, rune('X')) //old guard spot changed to X
		return currentPosition
	}

	nextRune := lookAtPosition(roomMap, currentPosition)

	switch nextRune {
	case rune('#'):
		rotateGuardRight(roomMap)
	case rune('.'), rune('X'):
		oldGuardSymbol := lookAtPosition(roomMap, originalGuardPosition)
		updateSpot(roomMap, currentPosition, oldGuardSymbol)  //move guard to . spot
		updateSpot(roomMap, originalGuardPosition, rune('X')) //old guard spot changed to X
	}

	return findGuard(roomMap)
}

func rotateGuardRight(roomMap [][]rune) {
	guardPosition := findGuard(roomMap)
	guardRune := lookAtPosition(roomMap, guardPosition)
	switch guardRune {
	case rune('^'):
		updateSpot(roomMap, guardPosition, rune('>'))
	case rune('v'):
		updateSpot(roomMap, guardPosition, rune('<'))
	case rune('<'):
		updateSpot(roomMap, guardPosition, rune('^'))
	case rune('>'):
		updateSpot(roomMap, guardPosition, rune('v'))
	}
}

func countX(roomMap [][]rune) int {
	totalX := 0
	for _, line := range roomMap {
		for _, elementRune := range line {
			if elementRune == rune('X') {
				totalX++
			}
		}
	}
	return totalX
}

func fillXinMap(roomMap [][]rune) {
	currentPosition := findGuard(roomMap)

	roomWidth := len(roomMap[0])
	roomHeight := len(roomMap)

	//printMap(roomMap)
	//reader := bufio.NewReader(os.Stdin)

	for insideBounds(currentPosition, roomWidth, roomHeight) {
		currentPosition = moveGuard(roomMap, currentPosition)
		//printMap(roomMap)
		//_, _ = reader.ReadByte()
	}
}

func solvePart1(filename string) (int, error) {
	roomMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	fillXinMap(roomMap)

	return countX(roomMap), nil
}

func solvePart2(filename string) (int, error) {
	roomMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	originalMap := roomMap
	//fill in 0 and 1 for start position, remember position later

	//printMap(roomMap)

	fillXinMap(roomMap)

	//store all X positions in array
	//for each X try placing #
	//	see if it loops?? 130^2 steps should be enough?
	//count up if it loops

	return 0, nil
}
