package day1

import (
	"slices"
	"strconv"
	"strings"

	"aoc2024backend/utils"
)

func Part1(filename string) int {
	// pair up all numbers
	// smallest to smallest
	// second-smallest to second-smallest
	// etc.
	// calc different of pairs
	// sum the differences

	input := utils.Readfile(filename)

	list1, list2 := buildLists(input)
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for index, val1 := range list1 {
		val2 := list2[index]
		diff := val1 - val2
		if diff > 0 {
			sum += diff
		} else {
			sum += -diff
		}
	}
	return sum
}

func buildLists(input string) ([]int, []int) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, "   ")

		v1, err := strconv.Atoi(values[0])
		utils.CheckFatal(err)
		v2, err := strconv.Atoi(values[1])
		utils.CheckFatal(err)

		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}

	return list1, list2
}
