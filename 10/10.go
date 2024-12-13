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
			if char != rune('.') {
				height, _ := strconv.Atoi(string(char))
				topographicMap[i][j] = height
			} else {
				//height, _ := strconv.Atoi(string(char))
				topographicMap[i][j] = -1
			}

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

func neighboursWithinBounds(trailhead Position, topographicMap [][]int) []Position {
	var neighbourList []Position
	width := len(topographicMap[0])
	height := len(topographicMap)

	if trailhead.x > 0 {
		neighbourX := trailhead.x - 1
		neighbourY := trailhead.y
		neighbourList = append(neighbourList, Position{x: neighbourX, y: neighbourY, height: topographicMap[neighbourX][neighbourY]})
	}

	if trailhead.y > 0 {
		neighbourX := trailhead.x
		neighbourY := trailhead.y - 1
		neighbourList = append(neighbourList, Position{x: neighbourX, y: neighbourY, height: topographicMap[neighbourX][neighbourY]})
	}

	if trailhead.x < width-1 {
		neighbourX := trailhead.x + 1
		neighbourY := trailhead.y
		neighbourList = append(neighbourList, Position{x: neighbourX, y: neighbourY, height: topographicMap[neighbourX][neighbourY]})
	}

	if trailhead.y < height-1 {
		neighbourX := trailhead.x
		neighbourY := trailhead.y + 1
		neighbourList = append(neighbourList, Position{x: neighbourX, y: neighbourY, height: topographicMap[neighbourX][neighbourY]})
	}

	return neighbourList
}

func findNeighboursOfOneHeigher(trailhead Position, topographicMap [][]int) []Position {
	validPositions := neighboursWithinBounds(trailhead, topographicMap)
	var goodPathForward []Position

	// fmt.Printf("Searching for higher spots to position %#v \n", trailhead)

	for _, validNeighbour := range validPositions {
		if validNeighbour.height == trailhead.height+1 {
			goodPathForward = append(goodPathForward, validNeighbour)
		}
	}
	return goodPathForward
}

func findAllPaths(trailhead Position, topographicMap [][]int, thePathSoFar []Position) [][]Position {
	// fmt.Println("Path so far", thePathSoFar)
	newPath := append([]Position{}, thePathSoFar...)
	//allFuturePaths := [][]Position{thePathSoFar}
	newPath = append(newPath, trailhead)
	//allFuturePaths[0] = append(allFuturePaths[0], trailhead)

	nextHeightPositions := findNeighboursOfOneHeigher(trailhead, topographicMap) //do this 9 times?
	//if len(nextPos) > 1, expand allPossiblePaths here?
	if len(nextHeightPositions) == 0 {
		return [][]Position{newPath}
	}

	var allFuturePaths [][]Position

	for _, nextPosition := range nextHeightPositions {
		// if index > 0 {
		// 	allFuturePaths = append(allFuturePaths, thePathSoFar)
		// }
		pathsFromNext := findAllPaths(nextPosition, topographicMap, newPath)
		allFuturePaths = append(allFuturePaths, pathsFromNext...)
	}

	//if len 9, stop?

	return allFuturePaths
}

func countTrailheadScore(trailhead Position, topographicMap [][]int) int {
	allPossiblePaths := []Position{}
	allFutures := findAllPaths(trailhead, topographicMap, allPossiblePaths)

	// trailHeadScore := 0
	// var uniquePeaks map[Position]bool
	uniquePeaks := make(map[Position]bool)
	// fmt.Println("all futures: ", allFutures)
	for _, pathFromTrailhead := range allFutures {
		// fmt.Println(pathFromTrailhead, len(pathFromTrailhead))

		if len(pathFromTrailhead) == 10 {
			uniquePeaks[pathFromTrailhead[9]] = true
		}
	}
	//rekurzivno pripni sebe k seznamu in se spet klici se je sosed prave visine
	//trivialen primer je da je seznam dolzine 9, takrat se sam vrne seznam

	return len(uniquePeaks)
}

func countTrailheadRating(trailhead Position, topographicMap [][]int) int {
	allPossiblePaths := []Position{}
	allFutures := findAllPaths(trailhead, topographicMap, allPossiblePaths)

	trailHeadRating := 0
	for _, pathFromTrailhead := range allFutures {
		// fmt.Println(pathFromTrailhead, len(pathFromTrailhead))

		if len(pathFromTrailhead) == 10 {
			trailHeadRating++
		}
	}
	//rekurzivno pripni sebe k seznamu in se spet klici se je sosed prave visine
	//trivialen primer je da je seznam dolzine 9, takrat se sam vrne seznam

	return trailHeadRating
}

func solvePart1(filename string) (int, error) {
	rawTopographicMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	topographicMap := parseRuneMap(rawTopographicMap)

	listOfTrailHeads := findAllTrailheads(topographicMap)

	totalScore := 0
	for _, trailhead := range listOfTrailHeads {
		totalScore += countTrailheadScore(trailhead, topographicMap)
	}

	return totalScore, nil
}

func solvePart2(filename string) (int, error) {
	rawTopographicMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	topographicMap := parseRuneMap(rawTopographicMap)

	listOfTrailHeads := findAllTrailheads(topographicMap)

	totalScore := 0
	for _, trailhead := range listOfTrailHeads {
		// fmt.Println("Doing trailhead ", trailhead)
		totalScore += countTrailheadRating(trailhead, topographicMap)
	}

	return totalScore, nil
}

func PrintMap(topographicMap [][]int) {

	for _, line := range topographicMap {
		for _, number := range line {
			if number != -1 {
				fmt.Printf("%d ", number)
			} else {
				fmt.Printf(". ")
			}

		}
		fmt.Println()

	}
}
