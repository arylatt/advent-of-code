package aoc202406

import (
	"fmt"
	"strconv"
	"sync"

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
	up, right, down, left := elves.Point{X: 0, Y: -1}, elves.Point{X: 1, Y: 0}, elves.Point{X: 0, Y: 1}, elves.Point{X: -1, Y: 0}
	grid, maxX, maxY, start := getGrid(input)
	answer := 0
	wg := &sync.WaitGroup{}

	for y, line := range grid {
		for x, char := range line {
			if char == '#' || char == '^' {
				continue
			}

			wg.Add(1)

			go func(x, y int) {
				points := map[string]int{}
				current := start
				direction, directionStr := up, "up"

				for {
					pointDir := fmt.Sprintf("%s-%s", current, directionStr)

					if _, ok := points[pointDir]; ok {
						answer++
						wg.Done()
						break
					}

					points[pointDir] = 1

					current = current.Shift(direction.X, direction.Y)
					if !current.Valid(maxX, maxY) {
						wg.Done()
						break
					}

					if grid[current.Y][current.X] == '#' || current.X == x && current.Y == y {
						// let's rotate the board.
						current = current.Shift(-direction.X, -direction.Y)

						if direction.Equals(up) {
							direction = right
							directionStr = "right"
						} else if direction.Equals(right) {
							direction = down
							directionStr = "down"
						} else if direction.Equals(down) {
							direction = left
							directionStr = "left"
						} else {
							direction = up
							directionStr = "up"
						}
					}
				}
			}(x, y)
		}
	}

	wg.Wait()

	return strconv.Itoa(answer)
}
