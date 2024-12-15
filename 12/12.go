package main

import (
	"bufio"
	"fmt"
	"os"
)

type Plot struct {
	x, y    int
	letter  rune
	groupID int
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

func parseGardenPlots(gardenPlotsStrings [][]rune) [][]Plot {
	gardenPlots := [][]Plot{}
	for i, gardenLine := range gardenPlotsStrings {
		gardenRow := []Plot{}
		for j, space := range gardenLine {
			newGardenPlot :=
				Plot{x: i,
					y:       j,
					letter:  space,
					groupID: -1}
			gardenRow = append(gardenRow, newGardenPlot)
		}
		gardenPlots = append(gardenPlots, gardenRow)
	}

	return gardenPlots
}

func plotGardenPlots(gardenPlots [][]Plot) {
	for _, plotRow := range gardenPlots {
		for _, plot := range plotRow {
			fmt.Printf("%d %d %c %d |", plot.x, plot.y, plot.letter, plot.groupID)
		}
		fmt.Println()

	}

}

func solvePart1(filename string) (int, error) {
	gardenPlotsStrings, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	gardenPlots := parseGardenPlots(gardenPlotsStrings)

	plotGardenPlots(gardenPlots)

	return 0, nil
}
