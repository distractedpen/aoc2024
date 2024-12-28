package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc2024backend/utils"
)

func Part1(filename string) int {
	input := utils.Readfile(filename)
	df := parseLines(input)

	safeCount := 0
	for _, line := range df {
		if isMonotonic(line) && isSafeDistance(line) {
			safeCount++
		} else {
			fmt.Println(line)
		}
	}

	return safeCount
}

func isMonotonic(values []int) bool {
	return slices.IsSorted(values) || slices.IsSortedFunc(values, func(a, b int) int {
		return b - a
	})
}

func isSafeDistance(values []int) bool {
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
	}

	return true
}

func parseLines(input string) [][]int {
	levelLinesString := strings.Split(strings.Trim(input, "\n"), "\n")

	levelLinesInt := make([][]int, 0)

	for _, line := range levelLinesString {
		intList := make([]int, 0)
		for _, strVal := range strings.Split(line, " ") {
			intVal, err := strconv.Atoi(strVal)
			utils.CheckFatal(err)
			intList = append(intList, intVal)
		}
		levelLinesInt = append(levelLinesInt, intList)

	}

	return levelLinesInt
}
