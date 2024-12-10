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
		fmt.Printf("%c\n", line)
	}
}

func solvePart1(filename string) (int, error) {
	antennaRuneMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	printRuneMap(antennaRuneMap)

	mapOfAntennta := parseMap(antennaRuneMap)
	printHashmap(mapOfAntennta)

	return 0, nil
}
