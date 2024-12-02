package aoc202402

import (
	"math"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	answer := 0

	for _, line := range lines {
		if markSafe(line, -1) {
			answer++
		}
	}

	return strconv.Itoa(answer)
}

func markSafe(line string, skipIndex int) bool {
	stringVals := strings.Split(line, " ")
	vals := []int{}
	incrementing := false

	if skipIndex != -1 {
		stringVals = append(stringVals[:skipIndex], stringVals[skipIndex+1:]...)
	}

	for i, stringVal := range stringVals {
		val, _ := strconv.Atoi(stringVal)
		vals = append(vals, val)

		if i == 0 {
			continue
		}

		if i == 1 {
			incrementing = vals[i] > vals[i-1]
		}

		if vals[i] == vals[i-1] {
			// unsafe, same value
			return false
		}

		if incrementing && vals[i] < vals[i-1] {
			// unsafe, decrementing when should be incrementing
			return false
		}

		if !incrementing && vals[i] > vals[i-1] {
			// unsafe, incrementing when should be decrementing
			return false
		}

		if math.Abs(float64(vals[i]-vals[i-1])) > 3 {
			// unsafe, difference greater than 3
			return false
		}
	}

	return true
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	answer := 0

	for _, line := range lines {
		for skipIndex := -1; skipIndex < len(strings.Split(line, " ")); skipIndex++ {
			if markSafe(line, skipIndex) {
				answer++
				break
			}
		}
	}

	return strconv.Itoa(answer)
}
