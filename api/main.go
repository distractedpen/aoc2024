package main

import (
	"log"
	"os"
	"strconv"

	"aoc2024backend/day1"
	"aoc2024backend/day2"
	"aoc2024backend/day3"
	"aoc2024backend/utils"
)

func main() {
	dayArg := os.Args[1]
	partArg := os.Args[2]
	fileInputArg := os.Args[3]

	day, err := strconv.Atoi(dayArg)
	utils.CheckFatal(err)
	part, err := strconv.Atoi(partArg)
	utils.CheckFatal(err)

	if len(fileInputArg) == 0 {
		log.Fatal("Provide input file")
	}

	programRunner(day, part, fileInputArg)
}

func programRunner(day int, part int, filename string) {
	switch day {
	case 1:
		switch part {
		case 1:
			res := day1.Part1(filename)
			log.Println(res)
		case 2:
			res := day1.Part2(filename)
			log.Println(res)
		}
	case 2:
		switch part {
		case 1:
			res := day2.Part1(filename)
			log.Println(res)
		case 2:
			res := day2.Part2(filename)
			log.Println(res)
		}
	case 3:
		switch part {
		case 1:
			res := day3.Part1(filename)
			log.Println(res)
		case 2:
			res := day3.Part2(filename)
			log.Println(res)
		}
	}
}
