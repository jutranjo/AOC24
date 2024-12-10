package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Sector struct {
	FileID int
	isFile bool
}

func readInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var diskMap []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			layoutNumber, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			diskMap = append(diskMap, layoutNumber)
		}
	}

	return diskMap, nil
}

func expandDenseDiskMap(denseDiskMap []int) []Sector {
	var fullDiskMap []Sector

	isNewBlockAFile := true
	currentBlockID := 0

	for _, numberOfBlocks := range denseDiskMap {
		if isNewBlockAFile {
			for range numberOfBlocks {
				newBlock := Sector{FileID: currentBlockID, isFile: true}
				fullDiskMap = append(fullDiskMap, newBlock)
			}
			currentBlockID++
			isNewBlockAFile = false
		} else {
			for range numberOfBlocks {
				newBlock := Sector{isFile: false}
				fullDiskMap = append(fullDiskMap, newBlock)
			}
			isNewBlockAFile = true
		}
	}

	return fullDiskMap
}

func fragmentTheDisk(fullDiskMap []Sector) {

}

func printFullDiskMap(fullDiskMap []Sector) {
	for _, sector := range fullDiskMap {
		if sector.isFile {
			fmt.Printf("%d", sector.FileID)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("\n")
}

func solvePart1(filename string) (int, error) {
	denseDiskMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	fullDiskMap := expandDenseDiskMap(denseDiskMap)
	printFullDiskMap(fullDiskMap)
	fragmentTheDisk(fullDiskMap)

	return 0, nil
}
