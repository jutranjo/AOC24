package main

import "fmt"

func main() {

	total, err := countAllXMAS("input04.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("total XMAS in part 1: ", total)

	total2, err := Part2MASCount("input04.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("total X shaped MAS in part2: ", total2)
}
