package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadReports(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into fields by whitespace
		parts := strings.Fields(line)

		// Convert the parts to integers
		var lineInts []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing number %q: %w", part, err)
			}
			lineInts = append(lineInts, num)
		}

		// Append the parsed integers for the current line
		result = append(result, lineInts)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func isSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	isIncreasing := true
	isDecreasing := true
	isCorrectDifference := true

	for i, number := range report {
		if i == 0 {
			continue
		}
		if number > report[i-1] {
			isDecreasing = false
		}
		if number < report[i-1] {
			isIncreasing = false
		}
		if diff := int(math.Abs(float64(number - report[i-1]))); 1 > diff || diff > 3 {
			isCorrectDifference = false
		}

	}
	return (isIncreasing || isDecreasing) && isCorrectDifference

}

func CountSafeReports(filename string) (int, error) {
	reports, err := ReadReports(filename)

	var safeCount int
	for _, line := range reports {
		if !isSafe(line) {
			continue
		}
		safeCount++
	}

	return safeCount, err
}

func CountTolerableReports(filename string) (int, error) {
	reports, err := ReadReports(filename)

	var TolerableCount int
	var isTolerableLine = false
	for _, line := range reports {
		//fmt.Println(line, isSafe(line), TolerableCount)

		if isSafe(line) {
			TolerableCount++
			continue
		}

		for i := range line {
			clonedLine := make([]int, len(line))
			copy(clonedLine, line)
			newline := append(clonedLine[:i], clonedLine[i+1:]...)
			//fmt.Println(newline, isSafe(newline))

			if isSafe(newline) {
				isTolerableLine = true
			}
		}

		if isTolerableLine {
			TolerableCount++
			isTolerableLine = false
		}

	}

	return TolerableCount, err
}
