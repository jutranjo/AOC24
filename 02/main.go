package main

import "fmt"

func main() {
	safeCount, err := CountSafeReports("02.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("Safe report count: ", safeCount)

	tolerableCount, err := CountTolerableReports("02.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("Tolerable report count: ", tolerableCount)

}
