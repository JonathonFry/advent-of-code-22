package main

import (
	"fmt"
	"sort"
	"strconv"
)

func day1Part1() {
	var total int
	var totals []int

	for _, line := range readInputAsLines(day1Input) {
		if line == "" || line == "\n" {
			totals = append(totals, total)
			total = 0
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			println(fmt.Sprintf("error: %v", err))
		}

		total += value
	}

	sort.Ints(totals)

	max := totals[len(totals)-1]

	println(fmt.Sprintf("Day 1 part 1: %v", max))
}

func day1Part2() {
	var total int
	var totals []int

	for _, line := range readInputAsLines(day1Input) {
		if line == "" || line == "\n" {
			totals = append(totals, total)
			total = 0
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			println(fmt.Sprintf("error: %v", err))
		}

		total += value
	}

	sort.Ints(totals)

	max := totals[len(totals)-1]
	max1 := totals[len(totals)-2]
	max2 := totals[len(totals)-3]
	println(fmt.Sprintf("Day 1 part 2: %v", max+max1+max2))
}
