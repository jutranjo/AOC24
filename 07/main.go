package main

import "fmt"

const FileToRun = "smallExample.txt"

func main() {
	result, err := solvePart1(FileToRun)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Part 1 result for file %s is %d", FileToRun, result)
}
