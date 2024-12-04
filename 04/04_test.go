package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSmallestText(t *testing.T) {
	input := ".SAMX."
	result := countXMASInLine(input)

	expected := 1

	if result != expected {
		t.Errorf("countXMASInLine(%s) = %d; want %d", input, result, expected)
	}
}

func TestSmallText(t *testing.T) {
	result, err := countAllXMAS("smalltest.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	expected := 4

	if result != expected {
		t.Errorf("countAllXMAS(smalltest.txt) = %d; want %d", result, expected)
	}
}

func TestOtherText(t *testing.T) {
	result, err := countAllXMAS("example2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	expected := 2

	if result != expected {
		t.Errorf("countAllXMAS(smalltest.txt) = %d; want %d", result, expected)
	}
}

func TestBiggerExample(t *testing.T) {
	result, err := countAllXMAS("biggerExample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	expected := 18

	if result != expected {
		t.Errorf("countAllXMAS(smalltest.txt) = %d; want %d", result, expected)
	}
}

func TestClockwiseRotation90(t *testing.T) {
	input := []string{
		"abc",
		"def",
		"ghi",
	}

	expected := []string{
		"gda",
		"heb",
		"ifc",
	}

	result := rotateTextClockwise90(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("rotateTextClockwise90 didn't work, got %s; want %s", result, expected)
	}
}

func TestClockwiseRotation45(t *testing.T) {
	input := []string{
		"abc",
		"def",
		"ghi",
	}

	expected := []string{
		"a",
		"db",
		"gec",
		"hf",
		"i",
	}

	result := rotateTextClockwise45(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("rotateTextClockwise45 didn't work, got %s; want %s", result, expected)
	}
}

func TestPart2Small(t *testing.T) {
	result, err := Part2MASCount("part2_small.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 1

	if result != expected {
		t.Errorf("Part2MASCount(part2_small.txt) = %d; want %d", result, expected)
	}
}

func TestPart2Large(t *testing.T) {
	result, err := Part2MASCount("part2large.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expected := 9

	if result != expected {
		t.Errorf("Part2MASCount(part2large.txt) = %d; want %d", result, expected)
	}
}
