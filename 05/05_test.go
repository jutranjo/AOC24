package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsUpdateCorrect(t *testing.T) {
	rules, updates, err := parseInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := []bool{true, true, true, false, false, false}

	for index, update := range updates {
		result := isUpdateCorrect(update, rules)
		if expected[index] != result {
			t.Errorf("error; got %v; want %v", result, expected)
		}
	}

}

func TestSummerPart1(t *testing.T) {
	result, err := solvePart1("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 143

	if expected != result {
		t.Errorf("Wrong sum of example part1, got %d, want %d", result, expected)
	}

}

func TestSummerPart2(t *testing.T) {
	result, err := solvePart2("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 123

	if expected != result {
		t.Errorf("Wrong sum of example part1, got %d, want %d", result, expected)
	}
}

func TestFindError(t *testing.T) {
	rules, updates, err := parseInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	result1, result2 := findWrongPlacement(updates[3], rules)
	expected1 := 1
	expected2 := 0

	if result1 != expected1 || result2 != expected2 {
		t.Errorf("One of two indexes wrong, either #1 got %d; wanted %d or #2 got %d; wanted %d", result1, result2, expected1, expected2)
	}
}

func TestFixUpdate(t *testing.T) {
	rules, updates, err := parseInput("smallExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var result []int
	var expected []int

	{
		fixUpdate(updates[3], rules)

		result = updates[3]
		expected = []int{97, 75, 47, 61, 53}

		if !reflect.DeepEqual(result, expected) {
			fmt.Println("Got: ", result)
			fmt.Println("Want: ", expected)
			t.Errorf("Got wrong update back!")
		}
	}

	{
		fixUpdate(updates[4], rules)

		result = updates[4]
		expected = []int{61, 29, 13}

		if !reflect.DeepEqual(result, expected) {
			fmt.Println("Got: ", result)
			fmt.Println("Want: ", expected)
			t.Errorf("Got wrong update back!")
		}
	}

	{
		for !(isUpdateCorrect(updates[5], rules)) {
			fixUpdate(updates[5], rules)
		}

		result = updates[5]
		expected = []int{97, 75, 47, 29, 13}

		if !reflect.DeepEqual(result, expected) {
			fmt.Println("Got: ", result)
			fmt.Println("Want: ", expected)
			t.Errorf("Got wrong update back!")
		}
	}

}
