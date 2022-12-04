package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func day4Part1() {
	trimmed := strings.TrimSuffix(day4Input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	var score int

	for _, line := range strings.Split(strings.TrimSuffix(trimmed, "\n"), "\n") {

		pairs := strings.Split(line, ",")
		first := pairs[0]
		firstArray := strings.Split(first, "-")
		firstExploded := explode(firstArray[0], firstArray[1])

		second := pairs[1]
		secondArray := strings.Split(second, "-")
		secondExploded := explode(secondArray[0], secondArray[1])

		intersection := intersect(firstExploded, secondExploded)

		if slices.Compare(intersection, firstExploded) == 0 || slices.Compare(intersection, secondExploded) == 0 {
			score += 1
		}
	}

	fmt.Println(fmt.Sprintf("Day 4 part 1: %v", score))
}

func day4Part2() {

	trimmed := strings.TrimSuffix(day4Input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	var score int

	for _, line := range strings.Split(strings.TrimSuffix(trimmed, "\n"), "\n") {

		pairs := strings.Split(line, ",")
		first := pairs[0]
		firstArray := strings.Split(first, "-")
		firstExploded := explode(firstArray[0], firstArray[1])

		second := pairs[1]
		secondArray := strings.Split(second, "-")
		secondExploded := explode(secondArray[0], secondArray[1])

		intersection := intersect(firstExploded, secondExploded)

		if len(intersection) > 0 {
			score += 1
		}
	}

	fmt.Println(fmt.Sprintf("Day 4 part 2: %v", score))
}

func explode(start, end string) []int {
	var s, e int

	s, _ = strconv.Atoi(start)
	e, _ = strconv.Atoi(end)

	var result []int
	for i := s; i <= e; i++ {
		result = append(result, i)
	}
	return result
}
