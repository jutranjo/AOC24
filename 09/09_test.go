package main

import (
	"slices"
	"testing"
)

func TestXxx(t *testing.T) {
	buggedExample := []int{1, 1, 3, 0, 5}
	diskMap := expandDenseDiskMap(buggedExample)

	printFullDiskMap(diskMap)

	got := findSectorBlockIndexes(diskMap, true)
	want := []Block{
		{startIndex: 0, length: 1, sector: Sector{FileID: 0, isFile: true}},
		{startIndex: 2, length: 3, sector: Sector{FileID: 1, isFile: true}},
		{startIndex: 5, length: 5, sector: Sector{FileID: 2, isFile: true}},
	}
	if !slices.Equal(got, want) {
		t.Errorf("2nd example line not correct, got %v, want %v", got, want)
	}
}
