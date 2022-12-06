package aoc202202

import (
	"strconv"
	"strings"
)

var (
	ShapeRock     = 1
	ShapePaper    = 2
	ShapeScissors = 3

	OutcomeLoss = 0
	OutcomeDraw = 3
	OutcomeWin  = 6
)

func LetterToShape(letter string) int {
	switch letter {
	case "A":
		fallthrough
	case "X":
		return ShapeRock
	case "B":
		fallthrough
	case "Y":
		return ShapePaper
	case "C":
		fallthrough
	case "Z":
		return ShapeScissors
	}

	return 0
}

func LetterToOutcome(letter string) int {
	switch letter {
	case "X":
		return OutcomeLoss
	case "Y":
		return OutcomeDraw
	case "Z":
		return OutcomeWin
	}

	return 0
}

func GameResult(opponent, self int) int {
	if opponent == self {
		return OutcomeDraw
	}

	if opponent == ShapeRock && self == ShapePaper {
		return OutcomeWin
	}

	if opponent == ShapePaper && self == ShapeScissors {
		return OutcomeWin
	}

	if opponent == ShapeScissors && self == ShapeRock {
		return OutcomeWin
	}

	return OutcomeLoss
}

func Predict(opponent, outcome int) int {
	if outcome == OutcomeDraw {
		return opponent
	}

	if outcome == OutcomeWin {
		if opponent == ShapeRock {
			return ShapePaper
		}

		if opponent == ShapePaper {
			return ShapeScissors
		}

		if opponent == ShapeScissors {
			return ShapeRock
		}
	}

	if opponent == ShapeRock {
		return ShapeScissors
	}

	if opponent == ShapePaper {
		return ShapeRock
	}

	if opponent == ShapeScissors {
		return ShapePaper
	}

	return 0
}

func Part1(input string) (output string) {
	score := 0
	inputs := strings.Split(input, "\n")

	for _, row := range inputs {
		if strings.TrimSpace(row) == "" {
			break
		}

		opponent, self := LetterToShape(string(row[0])), LetterToShape(string(row[2]))
		score += (self + GameResult(opponent, self))
	}

	return strconv.Itoa(score)
}

func Part2(input string) (output string) {
	score := 0
	inputs := strings.Split(input, "\n")

	for _, row := range inputs {
		if strings.TrimSpace(row) == "" {
			break
		}

		opponent, outcome := LetterToShape(string(row[0])), LetterToOutcome(string(row[2]))
		score += (outcome + Predict(opponent, outcome))
	}

	return strconv.Itoa(score)
}
