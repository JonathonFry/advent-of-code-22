package main

import (
	"fmt"
	"strings"
)

func day5Part1() {
	trimmed := strings.TrimSuffix(Day5Input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	var score int

	for _, line := range strings.Split(strings.TrimSuffix(trimmed, "\n"), "\n") {
		score += 1
		println(fmt.Sprintf("%v", line))
	}

	fmt.Println(fmt.Sprintf("Day 5 part 1: %v", score))
}

func day5Part2() {

	trimmed := strings.TrimSuffix(Day5Input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	var score int

	for _, line := range strings.Split(strings.TrimSuffix(trimmed, "\n"), "\n") {
		score += 1
		println(fmt.Sprintf("%v", line))
	}

	fmt.Println(fmt.Sprintf("Day 5 part 2: %v", score))
}
