package main

import "fmt"

const FileToRun = "input08.txt"

func main() {
	resultPart1, err := solvePart1(FileToRun)
	if err != nil {
		fmt.Println("Error while running part 1: ", err)
		return
	}

	fmt.Println("Part 1 antenna antinode count: ", resultPart1)
}
