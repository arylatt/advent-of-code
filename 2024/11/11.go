package aoc202411

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	stones := strings.Split(elves.SplitIntoLines(input)[0], " ")

	for range 25 {
		stones = applyRules(stones)
	}

	return strconv.Itoa(len(stones))
}

func applyRules(originalStones []string) (newStones []string) {
	for _, stone := range originalStones {
		if stone == "0" {
			newStones = append(newStones, "1")
			continue
		}

		if len(stone)%2 == 0 {
			half := len(stone) / 2

			newStones = append(newStones, strconv.Itoa(elves.Atoi(stone[:half])))
			newStones = append(newStones, strconv.Itoa(elves.Atoi(stone[half:])))
			continue
		}

		newStones = append(newStones, strconv.Itoa(elves.Atoi(stone)*2024))
	}

	return newStones
}

func Part2(input string) (output string) {
	return
}
