package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue        int
	RemainingNumbers []int
}

type operation func(int, int) int

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	return a / b
}

func parseLineIntoEquation(lineInput string) (Equation, error) {
	splitByColon := strings.Split(lineInput, ":")
	lineTestValue, err1 := strconv.Atoi(splitByColon[0])
	if err1 != nil {
		return Equation{}, err1
	}
	splitByComma := strings.Split(strings.TrimSpace(splitByColon[1]), " ")

	lineRemainingNumbers := make([]int, len(splitByComma))
	for i, s := range splitByComma {
		lineRemainingNumbers[i], _ = strconv.Atoi(s)
	}

	readLine := Equation{TestValue: lineTestValue, RemainingNumbers: lineRemainingNumbers}
	return readLine, nil
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
		lineEquation, err := parseLineIntoEquation(scanner.Text())
		if err != nil {
			return nil, err
		}

		lines = append(lines, lineEquation)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// make it recursive
func isTrueEquation(eq Equation) bool {

	forwardOperations := []operation{add, multiply}
	//are these even in the input data?
	if len(eq.RemainingNumbers) == 1 {
		return eq.TestValue == eq.RemainingNumbers[0]
	}
	//base example
	if len(eq.RemainingNumbers) == 2 {

		for _, op := range forwardOperations {
			a := eq.RemainingNumbers[0]
			b := eq.RemainingNumbers[1]
			result := op(a, b)
			fmt.Println("result:", result, "is it true?", result == eq.TestValue)

			if result == eq.TestValue {
				return true
			}
		}
		return false
	} else {
		forwardOperations := []operation{add, multiply}
		isTrue := false
		for _, op := range forwardOperations {
			shorterNumbers := append([]int{op(eq.RemainingNumbers[0], eq.RemainingNumbers[1])}, eq.RemainingNumbers[2:]...)
			shorterEq := Equation{TestValue: eq.TestValue, RemainingNumbers: shorterNumbers}
			fmt.Println("calling again with shorter eq: ", shorterEq)
			//return isTrueEquation(shorterEq)
			tempBool := isTrueEquation(shorterEq)
			if !isTrue {
				isTrue = tempBool
			}
		}
		return isTrue
	}
}

func solvePart1(filename string) (int, error) {
	listOfNumbers, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	sumOfTrueEquations := 0

	for _, equation := range listOfNumbers {
		fmt.Println("Line being looked at is :", equation)

		if isTrueEquation(equation) {
			sumOfTrueEquations += equation.TestValue
		}
	}

	return sumOfTrueEquations, nil
}
