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

type Block struct {
	startIndex int
	length     int
	sector     Sector
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
				newBlock := Sector{FileID: -1, isFile: false}
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

func findSectorBlockIndexes(fullDiskMap []Sector, findFiles bool) []Block {
	var listOfEmptySectorBlockIndexes []Block
	isNewBlock := false

	nextBlockToAdd := Block{sector: fullDiskMap[0]}
	for index, sector := range fullDiskMap {
		isNewID := !(sector.FileID == nextBlockToAdd.sector.FileID)
		//fmt.Printf("DISK INDEX: %d , at sector ID %d: new Block ID? %v bool:%v\n", index, sector.FileID, nextBlockToAdd.sector.FileID, isNewID)
		if (sector.isFile != findFiles || isNewID) && !isNewBlock {
			//fmt.Println("found end of bloc")
			isNewBlock = true
			nextBlockToAdd.length = index - nextBlockToAdd.startIndex
			if nextBlockToAdd.length != 0 {
				listOfEmptySectorBlockIndexes = append(listOfEmptySectorBlockIndexes, nextBlockToAdd)
			}
		}

		if sector.isFile == findFiles && isNewBlock {
			//fmt.Println("starting new block")
			isNewBlock = false
			nextBlockToAdd.startIndex = index
			nextBlockToAdd.sector = sector
		}

	}

	nextBlockToAdd.length = len(fullDiskMap) - nextBlockToAdd.startIndex
	lastSector := fullDiskMap[len(fullDiskMap)-1]

	if lastSector.isFile == findFiles && !isNewBlock {
		listOfEmptySectorBlockIndexes = append(listOfEmptySectorBlockIndexes, nextBlockToAdd)
	}

	return listOfEmptySectorBlockIndexes
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

func moveBlock(fullDiskMap []Sector, fileBlock Block, destinationIndex int) {
	for index := range fileBlock.length {
		fullDiskMap[destinationIndex+index] = fileBlock.sector
	}

}

func defragTheDisk(fullDiskMap []Sector) {
	emptySectorBlocks := findSectorBlockIndexes(fullDiskMap, false)
	fileSectorBlocks := findSectorBlockIndexes(fullDiskMap, true)

fileBlocks:
	for i := len(fileSectorBlocks) - 1; i > 0; i-- {
		requiredSpace := fileSectorBlocks[i].length
		for _, freeBlock := range emptySectorBlocks {
			if freeBlock.length >= requiredSpace && freeBlock.startIndex < fileSectorBlocks[i].startIndex {
				moveBlock(fullDiskMap, fileSectorBlocks[i], freeBlock.startIndex)
				newEmptyBlock := fileSectorBlocks[i]
				newEmptyBlock.sector = Sector{FileID: -1, isFile: false}
				moveBlock(fullDiskMap, newEmptyBlock, newEmptyBlock.startIndex)

				emptySectorBlocks = findSectorBlockIndexes(fullDiskMap, false)

				continue fileBlocks
			}
		}
		//printFullDiskMap(fullDiskMap)
	}
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
		if sector.isFile {
			checkSum += sector.FileID * index
		}

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

	//printFullDiskMap(fullDiskMap)
	defragTheDisk(fullDiskMap)

	result := calculateCheckSum(fullDiskMap)

	return result, nil
}
