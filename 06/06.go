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

type PositionDirection struct {
	Pos Position
	Dir GuardDirection
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

func findAllX(roomMap [][]rune) []Position {
	var AllX []Position
	for i, roomLine := range roomMap {
		for j, square := range roomLine {
			if rune('X') == square {
				AllX = append(AllX, Position{x: i, y: j})
			}
		}
	}
	return AllX
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

func movePositionInDirection(position Position, direction GuardDirection) Position {
	switch direction {
	case Up:
		position.x -= 1
	case Down:
		position.x += 1
	case Left:
		position.y -= 1
	case Right:
		position.y += 1
	}
	return position
}

func moveGuard(roomMap [][]rune, currentPosition Position) Position {
	originalGuardPosition := currentPosition

	direction := findGuardDirection(roomMap[currentPosition.x][currentPosition.y])
	currentPosition = movePositionInDirection(currentPosition, direction)

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
	case rune('0'):
		oldGuardSymbol := lookAtPosition(roomMap, originalGuardPosition)
		direction = findGuardDirection(oldGuardSymbol)
		updateSpot(roomMap, originalGuardPosition, rune('X')) //old guard spot changed to X

		currentPosition = movePositionInDirection(currentPosition, direction)
		updateSpot(roomMap, currentPosition, oldGuardSymbol) //move guard to . spot
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

func deepCopyMap(original [][]rune) [][]rune {
	copyMap := make([][]rune, len(original))
	for i := range original {
		copyMap[i] = make([]rune, len(original[i]))
		copy(copyMap[i], original[i])
	}
	return copyMap
}

func isGuardLooping(roomMap [][]rune, xSpot Position) bool {
	currentPosition := findGuard(roomMap)

	roomWidth := len(roomMap[0])
	roomHeight := len(roomMap)

	stepCount := 0
	stepHistory := make(map[PositionDirection]bool)
	for insideBounds(currentPosition, roomWidth, roomHeight) {
		stepCount++
		if xSpot.x == 22 && xSpot.y == 93 && stepCount > 10000 {
			printMap(roomMap)
		}
		direction := findGuardDirection(roomMap[currentPosition.x][currentPosition.y])
		//printMap(roomMap)
		currentPosition = moveGuard(roomMap, currentPosition)
		currentPosDir := PositionDirection{Pos: currentPosition, Dir: direction}

		currentPosition2 := movePositionInDirection(currentPosition, direction)
		var nextRune rune
		if insideBounds(currentPosition2, roomWidth, roomHeight) {
			nextRune = lookAtPosition(roomMap, currentPosition2)
		}

		if nextRune == rune('#') {
			currentPosition = moveGuard(roomMap, currentPosition)
			currentPosDir = PositionDirection{Pos: currentPosition, Dir: direction}

		}

		//fmt.Println("step is history? ", stepHistory[currentPosDir])
		if !stepHistory[currentPosDir] {
			//fmt.Println("Added spot to history: ", currentPosDir)
			stepHistory[currentPosDir] = true
		} else {
			return true
		}

		//stepCount++
		//if stepCount > roomWidth {
		//	return true
		//}
	}

	return insideBounds(currentPosition, roomWidth, roomHeight)
}

func solvePart2(filename string) (int, error) {
	roomMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	originalMap := deepCopyMap(roomMap)
	//fill in 0 and 1 for start position, remember position later
	//POTENTIAL BUG: can # be placed where guard started after he leaves?

	currentPosition := findGuard(roomMap)
	startSpot := currentPosition
	currentPosition = moveGuard(roomMap, currentPosition)
	updateSpot(roomMap, startSpot, rune('0'))

	inFrontOfGuardStart := currentPosition
	_ = moveGuard(roomMap, currentPosition)
	updateSpot(roomMap, inFrontOfGuardStart, rune('0'))

	fillXinMap(roomMap)

	//store all X positions in array
	allXpositions := findAllX(roomMap)
	fmt.Println("X to be checked: ", len(allXpositions))

	goodSpots := 0
	//for each X try placing #
	for _, xSpot := range allXpositions {
		if !(xSpot.x == 22 && xSpot.y == 93) {
			continue
		}
		fmt.Println("Checking spot ", xSpot)
		alteredMap := deepCopyMap(originalMap)
		updateSpot(alteredMap, xSpot, rune('#'))

		//	see if it loops?? 130^2 steps should be enough?
		if isGuardLooping(alteredMap, xSpot) {
			fmt.Println("Guard is looping!")
			goodSpots++
		}

		//count up if it loops
	}

	return goodSpots, nil
}
