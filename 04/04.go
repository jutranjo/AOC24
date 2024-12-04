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
	re1, _ := regexp.Compile(`XMAS`)
	re2, _ := regexp.Compile(`SAMX`)

	matches1 := re1.FindAllStringIndex(input, -1)
	matches2 := re2.FindAllStringIndex(input, -1)

	return len(matches1) + len(matches2)
}

func rotateTextClockwise90(matrix []string) []string {
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

func rotateTextClockwise45(matrix []string) []string {
	n := len(matrix)
	rotatedText := []string{}

	//upper triangle of the returned diamond matrix
	for i := 0; i < n; i++ {
		newRow := []byte{}
		for j := 0; j < i+1; j++ {
			newRow = append(newRow, matrix[i-j][j])
		}

		rotatedText = append(rotatedText, string(newRow))
	}

	//lower triangle of the returned diamond matrix
	for i := 1; i < n; i++ {
		newRow := []byte{}
		for j := 0; j < n-i; j++ {
			newRow = append(newRow, matrix[n-1-j][i+j])
		}

		rotatedText = append(rotatedText, string(newRow))
	}
	return rotatedText
}

func countAllXMAS(filename string) (int, error) {
	readFile, err := readMemory(filename)
	totalXMAS := 0

	//Left and right XMAS
	for _, line := range readFile {
		totalXMAS += countXMASInLine(line)
	}

	//Up and down XMAS
	for _, line := range rotateTextClockwise90(readFile) {
		totalXMAS += countXMASInLine(line)
	}

	//Diagonal in this direction \
	for _, line := range rotateTextClockwise45(readFile) {
		totalXMAS += countXMASInLine(line)
	}

	//Diagonal in this direction
	for _, line := range rotateTextClockwise45(rotateTextClockwise90(readFile)) {
		totalXMAS += countXMASInLine(line)
	}
	return totalXMAS, err
}
