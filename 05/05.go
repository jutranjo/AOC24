package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func ReadRules() int {
	return 0
}

func ReadInput(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	emptyLineIndex := slices.Index(lines, "")

	rulesText := lines[:emptyLineIndex]
	updateText := lines[emptyLineIndex+1:]

	return rulesText, updateText, nil
}

func parseRulesText(rulesText []string) map[int][]int {
	rulesMap := make(map[int][]int)

	re := regexp.MustCompile(`(\d+)\|(\d+)`)
	for _, ruleString := range rulesText {
		matches := re.FindAllString(ruleString, -1)
		splitNumbers := strings.Split(matches[0], "|")
		leftNumber, err1 := strconv.Atoi(splitNumbers[0])
		rightNumber, err2 := strconv.Atoi(splitNumbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("errors in parseRulesText: ", err1, err2)
			return nil
		}
		rulesMap[leftNumber] = append(rulesMap[leftNumber], rightNumber)
	}

	return rulesMap
}

func parseUpdateText(updatesText []string) [][]int {
	var updateArray [][]int
	for _, updateLine := range updatesText {
		splitPageStrings := strings.Split(updateLine, ",")
		var update []int
		for _, numberString := range splitPageStrings {
			if num, err := strconv.Atoi(numberString); err == nil {
				update = append(update, num)
			} else {
				fmt.Println("error in parseUpdateText: ", err)
			}
		}
		updateArray = append(updateArray, update)
	}
	return updateArray
}

func parseInput(filename string) (map[int][]int, [][]int, error) {
	rulesText, updateText, err := ReadInput(filename)
	if err != nil {
		return nil, nil, err
	}

	rules := parseRulesText(rulesText)
	updates := parseUpdateText(updateText)

	return rules, updates, nil
}

func findWrongPlacement(update []int, rules map[int][]int) (int, int) {
	swapIndex1 := -1
	swapIndex2 := -1

	for index, page := range update {
		prevPages := update[:index]
		for _, prevPage := range prevPages {
			for _, rule := range rules[page] {
				if rule == prevPage {
					swapIndex1 = index
					swapIndex2 = slices.Index(update, prevPage)
				}
			}

		}
	}
	return swapIndex1, swapIndex2
}

func isUpdateCorrect(update []int, rules map[int][]int) bool {
	index, _ := findWrongPlacement(update, rules)
	if index != -1 {
		return false
	} else {
		return true
	}
}

func rearrangeUpdate(update []int, rules map[int][]int) {

}

func solvePart1(filename string) (int, error) {
	rules, updates, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	//fmt.Println(rules)
	//fmt.Println(updates)

	sumOfMiddlePages := 0

	for _, update := range updates {
		if isUpdateCorrect(update, rules) {
			sumOfMiddlePages += update[len(update)/2]
		}
	}
	return sumOfMiddlePages, nil
}

func solvePart2(filename string) (int, error) {
	rules, updates, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	sumOfMiddlePages := 0

	for _, update := range updates {
		if !isUpdateCorrect(update, rules) {
			rearrangeUpdate(update, rules)
			sumOfMiddlePages += update[len(update)/2]
		}
	}
	return sumOfMiddlePages, nil
}
