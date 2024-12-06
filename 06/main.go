package main

import "fmt"

func main() {
	part1Result, err := solvePart1("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("total squares traversed in part 1: ", part1Result)

	part2Result, err := solvePart2("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("total squares traversed in part 1: ", part2Result)

}
