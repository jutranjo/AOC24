package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateDistance(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}
	defer file.Close()

	var leftNumbers []int
	var rightNumbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers:", err1, err2)
			continue
		}

		leftNumbers = append(leftNumbers, num1)
		rightNumbers = append(rightNumbers, num2)
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var totalDistance float64
	for i := 0; i < len(leftNumbers); i++ {
		totalDistance += math.Abs(float64(leftNumbers[i] - rightNumbers[i]))
	}

	return int(totalDistance), nil
}

func CalculateSimilarity(filename string) (int, error) {
	return 0, nil
}
