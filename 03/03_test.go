package main

import (
	"fmt"
	"testing"
)

func TestRegexSlice1(t *testing.T) {
	result := findMulPartsThenMultiply("xmul(2,4)")

	expected := 8

	if result != expected {
		t.Errorf("findMulParts(xmul(2,4)) = %d; want %d", result, expected)
	}
}

func TestRegexSlice2(t *testing.T) {
	result := findMulPartsThenMultiply("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+")

	expected := 8 + 25

	if result != expected {
		t.Errorf("findMulParts(xmul(2,4)) = %d; want %d", result, expected)
	}
}

func TestRegexSlice3(t *testing.T) {
	result := findMulPartsThenMultiply("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)m")

	expected := 8 + 25 + 11*8

	if result != expected {
		t.Errorf("findMulParts(xmul(2,4)) = %d; want %d", result, expected)
	}
}

func TestRegexWhole(t *testing.T) {
	result, err := sumAllLines("test3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 8 + 25 + 11*8 + 8*5

	if result != expected {
		t.Errorf("findMulParts(xmul(2,4)) = %d; want %d", result, expected)
	}
}

func TestPart2Whole(t *testing.T) {
	lines, err := readMemory("test3_2.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	result, _ := findMulPartsDoDont(lines[0], true)

	expected := 8 + 8*5

	if result != expected {
		t.Errorf("findMulParts(xmul(2,4)) = %d; want %d", result, expected)
	}
}
