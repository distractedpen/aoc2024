package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc2024backend/utils"
)

// 64984403 too low

func Part1(filename string) int {
	input := utils.Readfile(filename)
	input = strings.ReplaceAll(input, "\n", "")
	re := regexp.MustCompile(`mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)`)
	results := re.FindAllString(input, -1)
	fmt.Println(results)

	total := 0
	for _, entry := range results {
		fmt.Print(entry)
		re2 := regexp.MustCompile(`[0-9][0-9]*`)
		pairStr := re2.FindAllString(entry, -1)
		if len(pairStr) == 1 {
			fmt.Println("\t\tMismatch!!!", pairStr)
		}
		fmt.Print(" | ")
		fmt.Println(pairStr)
		val1, err := strconv.Atoi(pairStr[0])
		utils.CheckFatal(err)
		val2, err := strconv.Atoi(pairStr[1])
		utils.CheckFatal(err)
		total += val1 * val2
	}
	return total
}
