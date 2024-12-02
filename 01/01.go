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

func ReadNumbers(filename string) ([]int, []int, error) {
	var leftNumbers []int
	var rightNumbers []int
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return leftNumbers, rightNumbers, err
	}
	defer file.Close()

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
	return leftNumbers, rightNumbers, err
}

func CalculateDistance(filename string) (int, error) {
	leftNumbers, rightNumbers, err := ReadNumbers(filename)
	if err != nil {
		fmt.Println("Error reading numbers:", err)
		return 0, err
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	totalDistance := 0.0
	for i := 0; i < len(leftNumbers); i++ {
		totalDistance += math.Abs(float64(leftNumbers[i] - rightNumbers[i]))
	}

	return int(totalDistance), nil
}

func CalculateSimilarity(filename string) (int, error) {
	leftNumbers, rightNumbers, err := ReadNumbers(filename)
	if err != nil {
		fmt.Println("Error reading numbers:", err)
		return 0, err
	}

	similarityScore := 0
	for _, target := range leftNumbers {
		count := 0
		for _, num := range rightNumbers {
			if num == target {
				count++
			}
		}
		similarityScore += target * count
	}

	fmt.Println("similarity score: ", similarityScore)

	return similarityScore, nil
}
