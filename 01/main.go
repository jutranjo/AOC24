package main

import "fmt"

func main() {
	totalDiff, err := CalculateDistance("1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sum of Absolute Differences: ", totalDiff)
}
