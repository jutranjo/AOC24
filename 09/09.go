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

func findEmptySectorIndexes(fullDiskMap []Sector) []int {
	var listOfEmptySectorIndexes []int

	for index, sector := range fullDiskMap {
		if !sector.isFile {
			listOfEmptySectorIndexes = append(listOfEmptySectorIndexes, index)
		}
	}

	return listOfEmptySectorIndexes
}

func findFileSectorIndexes(fullDiskMap []Sector) []int {
	var listOfFileSectorIndexes []int

	for index, sector := range fullDiskMap {
		if sector.isFile {
			listOfFileSectorIndexes = append(listOfFileSectorIndexes, index)
		}
	}

	return listOfFileSectorIndexes
}

func isDiscFullyLeftmost(fullDiskMap []Sector) bool {
	//isFirstSectorFile := fullDiskMap[0].isFile
	haveWeReachedEmptySectors := false

	for _, sector := range fullDiskMap {
		//fmt.Printf("have we reached empty sectors? %v | ", haveWeReachedEmptySectors)
		//fmt.Println("Looking at sector ", index, sector)
		if !sector.isFile && !haveWeReachedEmptySectors {
			haveWeReachedEmptySectors = true
			//fmt.Println("Reached empty sectors at ", index, sector)
		}
		if sector.isFile && haveWeReachedEmptySectors {
			//fmt.Println("Found files again at ", index, sector)
			return false
		}
	}
	return true
}

func moveOneFileSectorToLeftmostEmptySpace(fullDiskMap []Sector, fileSectorIndexes *[]int, emptySectorIndexes *[]int) {
	destinationFreeSectorIndex := (*emptySectorIndexes)[0]
	*emptySectorIndexes = (*emptySectorIndexes)[1:]

	filesLen := len(*fileSectorIndexes)
	fileSectorToMoveIndex := (*fileSectorIndexes)[filesLen-1]
	*fileSectorIndexes = (*fileSectorIndexes)[:filesLen-1]

	//tempFreeSector := fullDiskMap[destinationFreeSector]
	fullDiskMap[destinationFreeSectorIndex], fullDiskMap[fileSectorToMoveIndex] = fullDiskMap[fileSectorToMoveIndex], fullDiskMap[destinationFreeSectorIndex]
	//= tempFreeSector
}

func fragmentTheDisk(fullDiskMap []Sector) {
	emptySectorIndexes := findEmptySectorIndexes(fullDiskMap)
	fileSectorIndexes := findFileSectorIndexes(fullDiskMap)

	for !isDiscFullyLeftmost(fullDiskMap) {
		moveOneFileSectorToLeftmostEmptySpace(fullDiskMap, &fileSectorIndexes, &emptySectorIndexes)
		//printFullDiskMap(fullDiskMap)
	}
}

func defragTheDisk(fullDiskMap []Sector) {

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

func calculateCheckSum(fullDiskMap []Sector) int {
	checkSum := 0
	for index, sector := range fullDiskMap {
		checkSum += sector.FileID * index
	}
	return checkSum
}

func solvePart1(filename string) (int, error) {
	denseDiskMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	fullDiskMap := expandDenseDiskMap(denseDiskMap)
	//printFullDiskMap(fullDiskMap)
	fragmentTheDisk(fullDiskMap)

	result := calculateCheckSum(fullDiskMap)

	return result, nil
}

func solvePart2(filename string) (int, error) {
	denseDiskMap, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	fullDiskMap := expandDenseDiskMap(denseDiskMap)

	defragTheDisk(fullDiskMap)

	result := calculateCheckSum(fullDiskMap)

	return result, nil
}
