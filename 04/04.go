package main

import (
	"bufio"
	"os"
	"regexp"
)

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

func countXMASInLine(input string) int {
	re, _ := regexp.Compile("XMAS|SAMX")

	matches := re.FindAllStringIndex(input, -1)

	return len(matches)
}

func rotateTextClockwise(matrix []string) []string {
	n := len(matrix)
	rotatedMatrix := make([]string, n)

	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = matrix[n-1-j][i]
		}
		rotatedMatrix[i] = string(row)
	}
	return rotatedMatrix
}

func countAllXMAS(filename string) (int, error) {
	readFile, err := readMemory(filename)
	totalXMAS := 0

	//Left and right XMAS
	for _, line := range readFile {
		totalXMAS += countXMASInLine(line)
	}

	//Up and down XMAS
	for _, line := range rotateTextClockwise(readFile) {
		totalXMAS += countXMASInLine(line)
	}

	//Diagonal in this direction \

	//Diagonal in this direction /
	return totalXMAS, err
}
