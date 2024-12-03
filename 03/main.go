package main

import "fmt"

func main() {
	output, err := sumAllLines("input03.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("main output", output)

	output2, err := sumDoDontAllLines(("input03.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("part 2 output: ", output2)
}
