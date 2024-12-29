package day4

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"aoc2024backend/utils"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Add(q Point) Point {
	return Point{
		p.X + q.X,
		p.Y + q.Y,
	}
}

func (p *Point) Sub(q Point) Point {
	return Point{
		p.X - q.X,
		p.Y - q.Y,
	}
}

type Size = Point // size and point are similar enough

func Part1(filename string) int {
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
            xmasFound := search(table, tableSize, startLoc, "XMAS")
            count += xmasFound
		}
	}

	return count
}

func printTable(table [][]string) {
    for _, r := range table {
        for _, val := range r {
            fmt.Printf("%s ", val)
        }
        fmt.Println()
    }
}

func makeTable(input string) [][]string {
	table := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		table = append(table, strings.Split(line, ""))
	}
	return table
}

func search(table [][]string, tableSize Size, startLoc Point, searchString string) int {
	// check eight directions around startLoc
    xmasCount := 0
	if table[startLoc.Y][startLoc.X] == "X" {
		for dY := -1; dY <= 1; dY++ {
			for dX := -1; dX <= 1; dX++ {
				if dY == 0 && dX == 0 {
					continue
				}

				// check the next three characters in the same direction
				nextCharX := startLoc.X + dX
				nextCharY := startLoc.Y + dY
				var sb strings.Builder
				for k := 0; k < utf8.RuneCountInString(searchString[1:]); k++ {
					if nextCharX < 0 || nextCharX >= tableSize.X {
						break
					}
					if nextCharY < 0 || nextCharY >= tableSize.Y {
						break
					}
					sb.WriteString(table[nextCharY][nextCharX])
                    nextCharX += dX
                    nextCharY += dY
				}
                if sb.String() == "MAS" {
                    fmt.Println(startLoc, Point{nextCharX, nextCharY})
                    xmasCount++
                }
			}
		}
	}
	return xmasCount
}
