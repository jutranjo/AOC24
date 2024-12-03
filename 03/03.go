package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func findMulPartsThenMultiply(input string) int {
	pattern := "mul\\((\\d+),(\\d+)\\)"
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	var sum int

	for _, match := range matches {
		a, err1 := strconv.Atoi(match[1])
		b, err2 := strconv.Atoi(match[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting strings to numbers:,", err1, err2)
			return 0
		}

		sum += a * b
	}

	return sum
}

func findMulPartsDoDont(input string, startState bool) (int, bool) {
	pattern := "mul\\((\\d+),(\\d+)\\)|(do\\(\\))|(don\\'t\\(\\))"
	re := regexp.MustCompile(pattern)

	isDoing := startState

	matches := re.FindAllStringSubmatch(input, -1)

	var sum int

	for _, match := range matches {
		if match[0] == "do()" {
			isDoing = true
		} else if match[0] == `don't()` {
			isDoing = false
		} else if isDoing {
			a, err1 := strconv.Atoi(match[1])
			b, err2 := strconv.Atoi(match[2])

			if err1 != nil || err2 != nil {
				return 0, true
			}
			sum += a * b
		} else {
			//fmt.Println("not mul match: ", match)
		}

	}

	return sum, isDoing
}

func readMemory(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, err
}

func sumAllLines(filename string) (int, error) {
	lines, err := readMemory(filename)
	if err != nil {
		return 0, err
	}

	var sum int
	for _, line := range lines {
		sum += findMulPartsThenMultiply(line)
	}

	return sum, nil
}

func sumDoDontAllLines(filename string) (int, error) {
	lines, err := readMemory(filename)
	if err != nil {
		return 0, err
	}

	var sum int
	doingState := true
	var partOfSum int
	for _, line := range lines {
		partOfSum, doingState = findMulPartsDoDont(line, doingState)
		sum += partOfSum
	}

	return sum, nil
}
