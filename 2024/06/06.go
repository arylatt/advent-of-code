package aoc202406

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	points := map[elves.Point]int{}
	up, right, down, left := elves.Point{X: 0, Y: -1}, elves.Point{X: 1, Y: 0}, elves.Point{X: 0, Y: 1}, elves.Point{X: -1, Y: 0}
	direction := up
	grid, maxX, maxY, start := getGrid(input)
	current := start

	for {
		points[current] = 1

		current = current.Shift(direction.X, direction.Y)
		if !current.Valid(maxX, maxY) {
			return strconv.Itoa(len(points))
		}

		if grid[current.Y][current.X] == '#' {
			// let's rotate the board.
			current = current.Shift(-direction.X, -direction.Y)

			if direction.Equals(up) {
				direction = right
			} else if direction.Equals(right) {
				direction = down
			} else if direction.Equals(down) {
				direction = left
			} else {
				direction = up
			}
		}
	}
}

func getGrid(input string) ([]string, int, int, elves.Point) {
	lines := elves.SplitIntoLines(input)

	maxX, maxY := len(lines[0])-1, len(lines)-1

	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				return lines, maxX, maxY, elves.Point{X: x, Y: y}
			}
		}
	}

	panic("things have gone very wrong indeed")
}

func Part2(input string) (output string) {
	return
}
