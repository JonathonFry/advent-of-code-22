package main

import (
	"fmt"
	"strings"
)

type Move struct {
	name     string
	strategy string
	score    int
}

var rock = Move{
	name:     "A",
	strategy: "X",
	score:    1,
}

var paper = Move{
	name:     "B",
	strategy: "Y",
	score:    2,
}

var scissors = Move{
	name:     "C",
	strategy: "Z",
	score:    3,
}

func day2Part1() {
	var score int
	for _, line := range readInputAsLines(day2Input) {
		if line == "" {
			continue
		}
		moves := strings.Split(line, " ")
		move := moveFromString(moves[0])
		strategy := strategyFromString(move, moves[1])

		score += calculateScore(move, strategy)
	}
	println(fmt.Sprintf("Day 2 part 1: %v", score))
}

func day2Part2() {
	var score int
	for _, line := range readInputAsLines(day2Input) {
		if line == "" {
			continue
		}
		moves := strings.Split(line, " ")
		move := moveFromString(moves[0])
		strategy := strategyFromString(move, moves[1])

		score += calculateScore(move, strategy)
	}
	println(fmt.Sprintf("Day 2 part 2: %v", score))
}

func moveFromString(value string) Move {
	switch {
	case value == rock.name || value == rock.strategy:
		return rock
	case value == paper.name || value == paper.strategy:
		return paper
	case value == scissors.name || value == scissors.strategy:
		return scissors
	}
	panic("error parsing move")
}

func strategyFromString(move Move, value string) Move {
	switch {
	case value == "X":
		// lose
		switch {
		case move == rock:
			return scissors
		case move == paper:
			return rock
		case move == scissors:
			return paper

		}
	case value == "Y":
		// draw
		return move
	case value == "Z":
		// win
		switch {
		case move == rock:
			return paper
		case move == paper:
			return scissors
		case move == scissors:
			return rock
		}
	}
	panic("error handling move")
}

func calculateScore(move, strategy Move) int {
	return result(move, strategy) + strategy.score
}

func result(move, strategy Move) int {
	// paper beats rock
	// rock beats scissors
	// scissors beats paper
	if move == rock {
		if strategy == scissors {
			return 0
		}
		if strategy == paper {
			return 6
		}
		if strategy == rock {
			return 3
		}
	}
	if move == paper {
		if strategy == scissors {
			return 6
		}
		if strategy == paper {
			return 3
		}
		if strategy == rock {
			return 0
		}
	}

	if move == scissors {
		if strategy == scissors {
			return 3
		}
		if strategy == paper {
			return 0
		}
		if strategy == rock {
			return 6
		}
	}

	return 0
}
