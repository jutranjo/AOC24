package main

import "fmt"

const FileToRun = "input07.txt"

func main() {
	result, err := solvePart1(FileToRun)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Part 1 result for file %s is %d \n", FileToRun, result)

	result2, err2 := solvePart2(FileToRun)
	if err2 != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Part 2 result for file %s is %d \n", FileToRun, result2)
}
