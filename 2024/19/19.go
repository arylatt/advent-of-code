package aoc202419

import (
	"slices"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	answer := 0

	towels, displays := getTowels(input)
	impossible := []string{}

	for _, display := range displays {
		slices.SortStableFunc(towels, func(a, b string) int {
			return len(b) - len(a)
		})

		slices.SortStableFunc(impossible, func(a, b string) int {
			return len(b) - len(a)
		})

		ok := false
		ok, impossible = canMakeDisplay(display, towels, impossible)
		if ok {
			towels = append(towels, display)
			answer++
		}
	}

	return strconv.Itoa(answer)
}

func getTowels(input string) ([]string, []string) {
	lines := elves.SplitIntoLines(input)

	towels := strings.Split(lines[0], ", ")
	displays := lines[1:]

	return towels, displays
}

func canMakeDisplay(display string, towels, impossible []string) (bool, []string) {
	if len(display) == 0 {
		return true, impossible
	}

	for _, imp := range impossible {
		if display == imp {
			return false, impossible
		}
	}

	possible := false

	for _, towel := range towels {
		if strings.HasPrefix(display, towel) {
			possible, impossible = canMakeDisplay(display[len(towel):], towels, impossible)
			if possible {
				return true, impossible
			} else {
				impossible = append(impossible, display)
			}
		}
	}

	return false, impossible
}

func canMakeDisplays(display string, towels, impossible []string) (int, []string) {
	if len(display) == 0 {
		return 1, impossible
	}

	for _, imp := range impossible {
		if display == imp {
			return 0, impossible
		}
	}

	displays, newDisplays := 0, 0

	for _, towel := range towels {
		if strings.HasPrefix(display, towel) {
			newDisplays, impossible = canMakeDisplays(display[len(towel):], towels, impossible)
			displays += newDisplays

			if newDisplays == 0 {
				impossible = append(impossible, display)
			}
		}
	}

	return displays, impossible
}

func Part2(input string) (output string) {
	answer := 0

	towels, displays := getTowels(input)
	impossible := []string{}

	slices.SortStableFunc(towels, func(a, b string) int {
		return len(b) - len(a)
	})

	for _, display := range displays {
		slices.SortStableFunc(impossible, func(a, b string) int {
			return len(b) - len(a)
		})

		newAnswer := 0
		newAnswer, impossible = canMakeDisplays(display, towels, impossible)
		answer += newAnswer
	}

	return strconv.Itoa(answer)
}
