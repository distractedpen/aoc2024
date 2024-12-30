package day5

import (
	"fmt"
	"slices"
)

func Part2(filename string) int {
	rulesTable, updateLists := parseFile(filename)
	printRulesTable(rulesTable)
	invalidLists := make([][]int, 0)
	for _, list := range updateLists {
		if !validateList(rulesTable, list) {
			invalidLists = append(invalidLists, list)
		}
	}

	total := 0
	for _, invalidList := range invalidLists {
		// fmt.Print("Before: ", invalidList, " | ")
		validList := correctInvalidList(rulesTable, invalidList)
		// fmt.Print("After: ", validList, "\n")
		total += getMiddleValue(validList)
	}

	return total
}

func printRulesTable(rulesTable map[int][]int) {
	for key, value := range rulesTable {
		fmt.Println(key, ": ", value)
	}
}

func correctInvalidList(rulesTable map[int][]int, invalidList []int) []int {
	validList := make([]int, len(invalidList))
	for i, c := range invalidList {
		rulesList := rulesTable[c]
		count := 0
		for j, v := range invalidList {
			if i == j {
				continue
			}
			if slices.Contains(rulesList, v) {
				count++
			}
		}
		validList[len(invalidList)-1-count] = c
	}

	return validList
}
