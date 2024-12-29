package day4

import (
	"slices"
	"strings"

	"aoc2024backend/utils"
)

func Part2(filename string) int {
	rawInput := utils.Readfile(filename)
	table := makeTable(rawInput)
	printTable(table)
	tableSize := Size{
		len(table[0]),
		len(table),
	}
	count := 0
	for y, row := range table {
		for x := range row {
			startLoc := Point{
				x,
				y,
			}
			if searchXmas(table, tableSize, startLoc) {
                count++
            }
		}
	}

	return count
}

func searchXmas(table [][]string, tableSize Size, startLoc Point) bool {
	// check eight directions around startLoc
	leftFlag := false
	rightFlag := false
	if table[startLoc.Y][startLoc.X] == "A" {
		// check if locs are in bounds
		if (startLoc.Y-1 < 0 || startLoc.Y+1 >= tableSize.Y) || (startLoc.X-1 < 0 || startLoc.X+1 >= tableSize.X) {
			return false
		}
		// pull the cross strings from the table
		var diagLeftSb strings.Builder
		var diagRightSb strings.Builder
		diagLeftSb.WriteString(table[startLoc.Y-1][startLoc.X-1] + table[startLoc.Y][startLoc.X] + table[startLoc.Y+1][startLoc.X+1])
		diagRightSb.WriteString(table[startLoc.Y-1][startLoc.X+1] + table[startLoc.Y][startLoc.X] + table[startLoc.Y+1][startLoc.X-1])
		diagLeft := diagLeftSb.String()
		diagRight := diagRightSb.String()

		// fmt.Println("Diag Left: ")
		// fmt.Print(diagLeft, " ")
		diagLeftS := strings.Split(diagLeft, "")
		slices.Reverse(diagLeftS)
		diagLeftRev := strings.Join(diagLeftS, "")
		// fmt.Println(diagLeftRev)
		if diagLeft == "MAS" || diagLeftRev == "MAS" {
			leftFlag = true
		}

		// fmt.Println("Diag Right: ")
		// fmt.Print(diagRight, " ")
		diagRightS := strings.Split(diagRight, "")
		slices.Reverse(diagRightS)
		diagRightRev := strings.Join(diagRightS, "")
		// fmt.Println(diagRightRev)
		if diagRight == "MAS" || diagRightRev == "MAS" {
			rightFlag = true
		}
		// fmt.Println("----")

	}
	return leftFlag && rightFlag
}
