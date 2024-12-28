package day1

import "aoc2024backend/utils"

func Part2(filename string) int {
	// pair up all numbers
	// smallest to smallest
	// second-smallest to second-smallest
	// etc.
	// calc different of pairs
	// sum the differences

	input := utils.Readfile(filename)

	list1, list2 := buildLists(input) // in part1

	countMap := make(map[int]int)

	for _, val2 := range list2 {
		_, ok := countMap[val2]
		if !ok {
			countMap[val2] = 1
		} else {
			countMap[val2]++
		}
	}

	sim := 0
	for _, val1 := range list1 {
		c, ok := countMap[val1]
		if ok {
			sim += val1 * c
		}
	}

	return sim
}
