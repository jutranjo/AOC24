package main

import (
	"bufio"
	"os"
)

type Equation struct {
	TestValue        int
	RemainingNumbers int
}

func readInput(filename string) ([]Equation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Equation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitUp := scanner.Text()
		//split it correctly
		lines = append(lines, Equation{TestValue: len(splitUp)})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func solvePart1(filename string) (int, error) {
	listOfNumbers, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	return len(listOfNumbers), nil
}
