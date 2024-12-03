package main

import "fmt"

func main() {
	output, err := sumAllLines("input03.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("main output", output)

}
