package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	number int
	next   *Stone
}

type LinkedStones struct {
	head *Stone
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers = strings.Fields(line)
	}

	return numbers, nil
}

func parseStartingStone(stoneStrings []string) LinkedStones {
	linkedStones := LinkedStones{}

	for _, numberString := range stoneStrings {
		newStoneNumber, _ := strconv.Atoi(numberString)
		// fmt.Println("number is ", newStoneNumber)

		newNode := &Stone{number: newStoneNumber}

		if linkedStones.head == nil {
			linkedStones.head = newNode
		} else {
			currentStone := linkedStones.head
			for currentStone.next != nil {
				currentStone = currentStone.next
			}
			currentStone.next = newNode
		}
	}

	return linkedStones
}

func printStones(stones LinkedStones) {
	stone := stones.head

	for stone != nil {
		fmt.Println(stone.number)
		stone = stone.next
	}
}

func applyRule1(stone *Stone) {
	stone.number = 1
}

func applyRule2(stone *Stone) {
	//split number into two
	numberToSplitString := strconv.Itoa(stone.number)
	// fmt.Println("Splitting", numberToSplitString)
	mid := len(numberToSplitString) / 2
	leftString := numberToSplitString[:mid]
	rightString := numberToSplitString[mid:]
	// fmt.Printf("left number %s, right number %s \n", leftString, rightString)
	leftNumber, _ := strconv.Atoi(leftString)
	rightNumber, _ := strconv.Atoi(rightString)

	newStone := Stone{number: rightNumber, next: stone.next}
	stone.number = leftNumber
	stone.next = &newStone
}

func applyRule3(stone *Stone) {
	stone.number *= 2024
}

func blink(stones LinkedStones) {
	stone := stones.head

	for stone != nil {
		if stone.number == 0 {
			applyRule1(stone)
		} else if len(strconv.Itoa(stone.number))%2 == 0 {
			applyRule2(stone)
			stone = stone.next
		} else {
			applyRule3(stone)
		}
		stone = stone.next
	}
}

func countStones(stones LinkedStones) int {
	stone := stones.head

	stoneCount := 0
	for stone != nil {
		stoneCount++
		stone = stone.next
	}
	return stoneCount
}

func highestStone(stones LinkedStones) int {
	stone := stones.head

	highStone := 0
	for stone != nil {
		if stone.number > highStone {
			highStone = stone.number
		}
		stone = stone.next
	}
	return highStone
}

func solvePart1(filename string) (int, error) {
	startingStonesString, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	stones := parseStartingStone(startingStonesString)

	for index := range 75 {
		fmt.Println("Blink count is ", index)
		fmt.Println("Stone count is ", countStones(stones))
		fmt.Println("Highest Stone number is ", highestStone(stones))
		blink(stones)
		//part 2: make a hashmap of all the stone counts maybe
	}

	stoneCount := countStones(stones)

	return stoneCount, nil
}
