package main

import "fmt"

const FileToRun = "smallExample.txt"

func main() {
	resultPart1, err := solvePart1(FileToRun)
	if err != nil {
		fmt.Println("Error while running part 1: ", err)
		return
	}

	fmt.Println("Part 1 checksum is ", resultPart1)

	resultPart2, err := solvePart2(FileToRun)
	if err != nil {
		fmt.Println("Error while running part 2: ", err)
		return
	}

	fmt.Println("Part 2 checksum is ", resultPart2)
}
