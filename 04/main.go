package main

import "fmt"

func main() {

	total, err := countAllXMAS("input04.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("total XMAS in part 1: ", total)
}
