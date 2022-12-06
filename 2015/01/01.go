package aoc201501

import "strconv"

func Part1(input string) (output string) {
	floor := 0

	for _, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return strconv.Itoa(floor)
}

func Part2(input string) (output string) {
	floor := 0

	for i, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			return strconv.Itoa(i + 1)
		}
	}

	return strconv.Itoa(floor)
}
