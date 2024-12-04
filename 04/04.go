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

func lookDiagonal(grid []string, row int, column int, rowDirection int, columnDirection int) rune {
	return rune(grid[row+rowDirection][column+columnDirection])
}

func isACenteredOnTwoMAS(puzzleCharacters []string, row int, column int) bool {
	upLeft := lookDiagonal(puzzleCharacters, row, column, -1, -1)
	downRight := lookDiagonal(puzzleCharacters, row, column, 1, 1)
	if (upLeft == 'M' && downRight == 'S') || (upLeft == 'S' && downRight == 'M') {
		upRight := lookDiagonal(puzzleCharacters, row, column, -1, 1)
		downLeft := lookDiagonal(puzzleCharacters, row, column, 1, -1)
		if (upRight == 'M' && downLeft == 'S') || (upRight == 'S' && downLeft == 'M') {
			return true
		}
	}
	return false
}

func Part2MASCount(filename string) (int, error) {
	puzzleCharacters, err := readMemory(filename)

	n, m := len(puzzleCharacters), len(puzzleCharacters[0])

	totalXMAS := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if puzzleCharacters[i][j] == 'A' {
				if isACenteredOnTwoMAS(puzzleCharacters, i, j) {
					totalXMAS++
				}
			}
		}
	}

	return totalXMAS, err
}
