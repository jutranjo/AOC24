package main

import "fmt"

const FileToRun = "input10.txt"

func main() {
	resultPart1, err := solvePart1(FileToRun)
	if err != nil {
		fmt.Println("Error while running part 1: ", err)
		return
	}

	fmt.Println("Part 1 trailhead score is ", resultPart1)

	resultPart2, err := solvePart2(FileToRun)
	if err != nil {
		fmt.Println("Error while running part 1: ", err)
		return
	}

	fmt.Println("Part 2 trailhead rating is ", resultPart2)
}
