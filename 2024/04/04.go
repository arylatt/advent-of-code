package aoc202404

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	lines, maxX, maxY := setupGrid(input)

	answer := 0

	for y, line := range lines {
		for x := range line {
			letter := string(line[x])

			if letter != "X" {
				continue
			}

			point := elves.Point{x, y}

			directions := [][]elves.Point{
				{point.Shift(0, 1), point.Shift(0, 2), point.Shift(0, 3)},
				{point.Shift(1, 0), point.Shift(2, 0), point.Shift(3, 0)},
				{point.Shift(1, 1), point.Shift(2, 2), point.Shift(3, 3)},
				{point.Shift(-1, 1), point.Shift(-2, 2), point.Shift(-3, 3)},
				{point.Shift(1, -1), point.Shift(2, -2), point.Shift(3, -3)},
				{point.Shift(-1, -1), point.Shift(-2, -2), point.Shift(-3, -3)},
				{point.Shift(-1, 0), point.Shift(-2, 0), point.Shift(-3, 0)},
				{point.Shift(0, -1), point.Shift(0, -2), point.Shift(0, -3)},
			}

			for _, direction := range directions {
				if buildWord(direction, lines, maxX, maxY, "X") == "XMAS" {
					answer++
				}
			}
		}
	}

	return strconv.Itoa(answer)
}

func buildWord(direction []elves.Point, lines []string, maxX, maxY int, start string) string {
	word := start

	for _, point := range direction {
		if !point.Valid(maxX, maxY) {
			return word
		}

		word += string(lines[point.Y][point.X])
	}

	return word
}

func setupGrid(input string) ([]string, int, int) {
	lines := elves.SplitIntoLines(input)
	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	return lines, maxX, maxY
}

func Part2(input string) (output string) {
	lines, maxX, maxY := setupGrid(input)

	answer := 0
	masMap := map[elves.Point]int{}

	for y, line := range lines {
		for x := range line {
			letter := string(line[x])

			if letter != "M" {
				continue
			}

			point := elves.Point{x, y}

			directions := [][]elves.Point{
				{point.Shift(1, 1), point.Shift(2, 2)},
				{point.Shift(-1, 1), point.Shift(-2, 2)},
				{point.Shift(1, -1), point.Shift(2, -2)},
				{point.Shift(-1, -1), point.Shift(-2, -2)},
			}

			for _, direction := range directions {
				if buildWord(direction, lines, maxX, maxY, "M") == "MAS" {
					if _, ok := masMap[direction[0]]; !ok {
						masMap[direction[0]] = 0
					}

					masMap[direction[0]]++

					if masMap[direction[0]] == 2 {
						answer++
					}
				}
			}
		}
	}

	return strconv.Itoa(answer)
}
