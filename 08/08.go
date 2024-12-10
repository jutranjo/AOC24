package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x, y int
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

func parseMap(inputGrid [][]rune) map[rune][]Coordinate {
	antennaHash := make(map[rune][]Coordinate)

	for i, line := range inputGrid {
		for j, symbol := range line {
			if symbol != rune('.') {
				antennaHash[symbol] = append(antennaHash[symbol], Coordinate{x: i, y: j})
			}

		}
	}

	return antennaHash
}

func printHashmap(mapOfAntennta map[rune][]Coordinate) {

	for key, values := range mapOfAntennta {
		fmt.Printf("key: %c \n", key)
		for i, value := range values {
			fmt.Printf("	Value %d is %d \n", i, value)
		}
	}

}

func printRuneMap(antennaRuneMap [][]rune) {
	for _, line := range antennaRuneMap {
		for _, char := range line {
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func findANCoordinate(ant1 Coordinate, ant2 Coordinate) Coordinate {
	deltaX := ant2.x - ant1.x
	deltaY := ant2.y - ant1.y

	return Coordinate{x: ant2.x + deltaX, y: ant2.y + deltaY}
}

func isAntiNodeInBounds(antinodeCoordinate Coordinate, antennaRuneMap [][]rune) bool {
	width := len(antennaRuneMap[0]) - 1
	height := len(antennaRuneMap) - 1
	return antinodeCoordinate.x <= width &&
		antinodeCoordinate.x >= 0 &&
		antinodeCoordinate.y >= 0 &&
		antinodeCoordinate.y <= height
}

func countAntinodes(antennaHashMap map[rune][]Coordinate, antennaRuneMap [][]rune) int {
	//antinodeHashMap := make(map[rune][]Coordinate)
	antinodesFound := make(map[Coordinate]bool)

	for _, values := range antennaHashMap {
		for _, antenna1 := range values {
			for _, antenna2 := range values {
				if antenna1 != antenna2 {
					//fmt.Printf("key %c antenna %d-%d ", key, antenna1, antenna2)

					antinodeCoordinate := findANCoordinate(antenna1, antenna2)

					//fmt.Printf("is %d in bounds? %v \n", antinodeCoordinate, isAntiNodeInBounds(antinodeCoordinate, antennaRuneMap))
					if isAntiNodeInBounds(antinodeCoordinate, antennaRuneMap) {
						antinodesFound[antinodeCoordinate] = true
						antennaRuneMap[antinodeCoordinate.x][antinodeCoordinate.y] = rune('#')
						//fmt.Printf("placed antinode of %c at %d \n", key, antinodeCoordinate)
					}
				}

			}
		}
	}
	printRuneMap(antennaRuneMap)
	return len(antinodesFound)
}

func solvePart1(filename string) (int, error) {
	antennaRuneMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	printRuneMap(antennaRuneMap)

	antennaHashMap := parseMap(antennaRuneMap)
	//printHashmap(antennaHashMap)

	antinodeCount := countAntinodes(antennaHashMap, antennaRuneMap)

	return antinodeCount, nil
}
