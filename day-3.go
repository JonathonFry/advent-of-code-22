package main

import (
	"fmt"
	"strings"
	"unicode"
)


/*
The ASCII value of the lowercase alphabet is from 97 to 122.
And, the ASCII value of the uppercase alphabet is from 65 to 90.
-96
*/

func day3Part1() {
	var score int

	for _, line := range readInputAsLines(day3Input) {
		if line == "" {
			continue
		}

		chars := strings.Split(line, "")
		size := len(line) / 2
		first := chars[0:size]
		second := chars[size:len(line)]
		values := intersect(first, second)

		value := charCodeAt(values[0], 0)
		if unicode.IsLower(value) {
			value -= 96
		} else {
			value -= 38
		}

		score += int(value)
	}
	println(fmt.Sprintf("Day 3 part 1: %v", score))
}
func day3Part2() {
	var score int

	backpacks := readInputAsLines(day3Input)

	for i := 0; i < len(backpacks); i += 3 {
		intersection := intersectThree(strings.Split(backpacks[i], ""), strings.Split(backpacks[i+1], ""), strings.Split(backpacks[i+2], ""))
		value := charCodeAt(intersection[0], 0)
		if unicode.IsLower(value) {
			value -= 96
		} else {
			value -= 38
		}
		score += int(value)
	}
	println(fmt.Sprintf("Day 3 part 2: %v", score))
}
