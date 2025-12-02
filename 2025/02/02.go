package aoc202502

import (
	"math"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func parseInput(input string) [][2]string {
	output := [][2]string{}

	lines := elves.SplitIntoLines(input)

	for _, vals := range strings.Split(lines[0], ",") {
		pair := strings.Split(vals, "-")

		output = append(output, [2]string{pair[0], pair[1]})
	}

	return output
}

func isInvalidId(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	mid := len(id) / 2
	left := id[:mid]
	right := id[mid:]

	return left == right
}

func calculateNextInvalidId(id string) string {
	if len(id)%2 != 0 {
		id = strconv.Itoa(int(math.Pow(10, float64(len(id)))))

		mid := len(id) / 2
		return id[:mid] + id[:mid]
	}

	mid := len(id) / 2
	left := id[:mid]

	if (left + left) <= id {
		leftInt, _ := strconv.Atoi(left)
		leftInt++
		left = strconv.Itoa(leftInt)
	}

	return left + left
}

func Part1(input string) (output string) {
	inputs := parseInput(input)
	answer := 0

	for _, pair := range inputs {
		upperLimit := elves.Atoi(pair[1])
		nextInvalid := elves.Atoi(calculateNextInvalidId(pair[0]))

		if isInvalidId(pair[0]) {
			answer += elves.Atoi(pair[0])
		}

		for nextInvalid <= upperLimit {
			answer += nextInvalid

			nextVal := nextInvalid + 1
			if nextVal > upperLimit {
				break
			}

			nextInvalid = elves.Atoi(calculateNextInvalidId(strconv.Itoa(nextVal)))
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	return
}
