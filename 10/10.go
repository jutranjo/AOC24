package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	x, y   int
	height int
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

func parseRuneMap(runeMap [][]rune) [][]int {
	n := len(runeMap)
	m := len(runeMap[0])
	topographicMap := make([][]int, n)
	for i := range topographicMap {
		topographicMap[i] = make([]int, m)
	}

	for i, line := range runeMap {
		for j, char := range line {
			height, _ := strconv.Atoi(string(char))
			topographicMap[i][j] = height
		}
	}
	return topographicMap
}

func findAllTrailheads(topographicMap [][]int) []Position {
	var allTrailHeads []Position

	for i, row := range topographicMap {
		for j, height := range row {
			if height == 0 {
				newTrailhead := Position{x: i, y: j, height: 0}
				allTrailHeads = append(allTrailHeads, newTrailhead)
			}

		}
	}

	return allTrailHeads
}

func countTrailheadScore(trailhead Position, topographicMap [][]int) int {

	return 1
}

func solvePart1(filename string) (int, error) {
	rawTopographicMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	topographicMap := parseRuneMap(rawTopographicMap)

	for _, line := range topographicMap {
		fmt.Printf("%d \n", line)
	}

	listOfTrailHeads := findAllTrailheads(topographicMap)

	totalScore := 0
	for _, trailhead := range listOfTrailHeads {
		totalScore += countTrailheadScore(trailhead, topographicMap)
	}

	return totalScore, nil
}
