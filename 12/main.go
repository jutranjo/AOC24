package main

import "fmt"

const FileToRun = "smallExample1.txt"

func main() {
	fmt.Println("Running file ", FileToRun)
	resultPart1, err := solvePart1(FileToRun)
	if err != nil {
		fmt.Println("Error while running part 1: ", err)
		return
	}

	fmt.Println("Part 1 fence cost is ", resultPart1)
}
