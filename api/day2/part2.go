package day2

import (
	"fmt"
	"slices"

	"aoc2024backend/utils"
)

func Part2(filename string) int {
	input := utils.Readfile(filename)
	df := parseLines(input)

	safeCount := 0
	usedProblemDampener := false
	for _, line := range df {
		if checkProblemDampener(line, usedProblemDampener) {
			safeCount++
		} else {
			fmt.Println(line)
		}
	}

	return safeCount
}

func checkProblemDampener(line []int, usedProblemDampener bool) bool {
	// try running the line
	if isMonotonic(line) && isSafeDistance(line) {
		return true
	}

	if !usedProblemDampener {
		// remove each element from the list of entries
		for rmvIndex := 0; rmvIndex < len(line); rmvIndex++ {
			newLine := slices.Clone(line)
			swapToEnd(&newLine, rmvIndex)
			newLine = newLine[0 : len(newLine)-1]
			// with removing a single element, allow it to pass
			if isMonotonic(newLine) && isSafeDistance(newLine) {
				usedProblemDampener = true
				return true
			}
		}
	}
	return false
}

func swapToEnd(line *[]int, index int) {
	for i := index; i < len(*line)-1; i++ {
		tmp := (*line)[i+1]
		(*line)[i+1] = (*line)[i]
		(*line)[i] = tmp
	}
}
