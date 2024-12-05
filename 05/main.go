package main

import "fmt"

func main() {
	result, err := solvePart1("input05.txt")
	if err != nil {
		fmt.Println("got error ", err)
		return
	}
	fmt.Println("Part1 sum =", result)

	result2, err2 := solvePart2("input05.txt")
	if err2 != nil {
		fmt.Println("got error ", err)
		return
	}
	fmt.Println("Part1 sum =", result2)
}
