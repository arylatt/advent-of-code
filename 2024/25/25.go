package aoc202425

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	locks, keys := parseInput(input)

	answer := 0

	for _, lock := range locks {
		for _, key := range keys {
			if keyFitsLock(key, lock) {
				answer++
			}
		}
	}

	return strconv.Itoa(answer)
}

func parseInput(input string) (locks, keys [][5]int) {
	lines := elves.SplitIntoLines(input)

	for i := 0; i < len(lines); i += 7 {
		if lines[i] == "#####" {
			locks = append(locks, buildLock(lines[i+1:i+7]))
		} else {
			keys = append(keys, buildKey(lines[i:i+6]))
		}
	}

	return
}

func buildLock(lines []string) (lock [5]int) {
	for i := range 5 {
		for j := range 7 {
			if lines[j][i] == '#' {
				lock[i]++
			} else {
				break
			}
		}
	}

	return
}

func buildKey(lines []string) (key [5]int) {
	for i := range 5 {
		for j := 5; j >= 0; j-- {
			if lines[j][i] == '#' {
				key[i]++
			} else {
				break
			}
		}
	}

	return
}

func keyFitsLock(key, lock [5]int) bool {
	for i := range 5 {
		if key[i]+lock[i] > 5 {
			return false
		}
	}

	return true
}

func Part2(input string) (output string) {
	return
}
