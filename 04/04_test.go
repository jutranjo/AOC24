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

func TestClockwiseRotation(t *testing.T) {
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

	result := rotateTextClockwise(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("rotateTextClockwise didn't work, got %s; want %s", result, expected)
	}
}
