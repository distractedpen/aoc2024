package day5

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc2024backend/utils"
)

func Part1(filename string) int {
	rulesTable, updateLists := parseFile(filename)

	total := 0
	for _, list := range updateLists {
		if validateList(rulesTable, list) {
			total += getMiddleValue(list)
		}
	}

	return total
}

func validateList(rulesTable map[int][]int, updateList []int) bool {
	for ind, page := range updateList[:len(updateList)-1] {
		preReqs := rulesTable[page]
		for _, nextPage := range updateList[ind+1:] {
			if !slices.Contains(preReqs, nextPage) {
				return false
			}
		}
	}
	return true
}

func getMiddleValue(updateList []int) int {
	midPoint := (len(updateList) / 2)
	return updateList[midPoint]
}

func parseFile(filename string) (map[int][]int, [][]int) {
	rulesTable := make(map[int][]int)
	updateLists := make([][]int, 0)

	data, err := os.ReadFile(filename)
	utils.CheckFatal(err)

	parsingRules := true
	for _, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		if len(line) == 0 {
			parsingRules = false
			continue
		}

		if parsingRules {
			// parsing Rules Table
			rawStringVals := strings.Split(line, "|")
			page, err := strconv.Atoi(rawStringVals[0])
			utils.CheckFatal(err)
			predPage, err := strconv.Atoi(rawStringVals[1])
			utils.CheckFatal(err)

			_, ok := rulesTable[page]
			if !ok {
				rulesTable[page] = make([]int, 0)
			}
			rulesTable[page] = append(rulesTable[page], predPage)
		} else {
			// parsing Update List
			pages := make([]int, 0)
			for _, val := range strings.Split(line, ",") {
				iVal, err := strconv.Atoi(val)
				utils.CheckFatal(err)
				pages = append(pages, iVal)
			}
			updateLists = append(updateLists, pages)
		}
	}

	return rulesTable, updateLists
}
