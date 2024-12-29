package day3

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"aoc2024backend/utils"
)

func Part2(filename string) int {
	input := utils.Readfile(filename)
	input = strings.ReplaceAll(input, "\n", "")
	re := regexp.MustCompile(`mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)`)
	reDo := regexp.MustCompile(`do[^n't]`)
	reDont := regexp.MustCompile(`don't`)
	resultIndPairs := re.FindAllStringIndex(input, -1)
	results := re.FindAllString(input, -1)
	doIndPairs := reDo.FindAllStringIndex(input, -1)
	dontIndPairs := reDont.FindAllStringIndex(input, -1)

	resultLocs := make([]int, len(resultIndPairs))
	doLocs := make([]int, len(doIndPairs))
	dontLocs := make([]int, len(dontIndPairs))

	for ind, pair := range resultIndPairs {
		resultLocs[ind] = pair[0]
	}
	for ind, pair := range doIndPairs {
		doLocs[ind] = pair[0]
	}
	for ind, pair := range dontIndPairs {
		dontLocs[ind] = pair[0]
	}

	allInd := slices.Concat(resultLocs, doLocs, dontLocs)
	slices.Sort(allInd)

	tape := make(map[int]string)

	// add instructions
	for ind, loc := range resultIndPairs {
		res := results[ind]
		tape[loc[0]] = res
	}
	// add do instructions
	for _, loc := range doIndPairs {
		tape[loc[0]] = "do"
	}

	for _, loc := range dontIndPairs {
		tape[loc[0]] = "don't"
	}

	total := 0
	enableFlag := true
	for _, loc := range allInd {
		tapeVal := tape[loc]
		switch tapeVal {
		case "do":
            enableFlag = true   
		case "don't":
            enableFlag = false
		default:
			if enableFlag {
				re2 := regexp.MustCompile(`[0-9][0-9]*`)
				pairStr := re2.FindAllString(tapeVal, -1)
				if len(pairStr) == 1 {
					fmt.Println("\t\tMismatch!!!", pairStr)
				}
				val1, err := strconv.Atoi(pairStr[0])
				utils.CheckFatal(err)
				val2, err := strconv.Atoi(pairStr[1])
				utils.CheckFatal(err)
				total += val1 * val2
			}
		}
	}
	return total
}
