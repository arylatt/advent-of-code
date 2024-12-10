package aoc202410

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	grid, maxX, maxY := elves.ParseGrid(input)

	starts := findStarts(grid)

	answer := 0

	for _, start := range starts {
		routes, _ := solveTo9(grid, maxX, maxY, start, -1, []string{})
		answer += routes
	}

	return strconv.Itoa(answer)
}

func findStarts(grid []string) (starts []elves.Point) {
	for y, row := range grid {
		for x, char := range row {
			if char == '0' {
				starts = append(starts, elves.Point{X: x, Y: y})
			}
		}
	}

	return
}

func solveTo9(grid []string, maxX, maxY int, current elves.Point, lastVal int, visited []string) (int, []string) {
	currentStr := fmt.Sprintf("%s-from-%d", current, lastVal)

	if slices.Contains(visited, currentStr) {
		return 0, visited
	}

	visited = append(visited, currentStr)

	if !current.Valid(maxX, maxY) || grid[current.Y][current.X] == '.' {
		return 0, visited
	}

	currentVal := elves.Atoi(string(grid[current.Y][current.X]))
	if currentVal != lastVal+1 {
		return 0, visited
	}

	if grid[current.Y][current.X] == '9' {
		return 1, visited
	}

	routes := 0

	upCount, upVisited := solveTo9(grid, maxX, maxY, current.Shift(0, -1), currentVal, visited)
	routes += upCount
	visited = upVisited

	rightCount, rightVisited := solveTo9(grid, maxX, maxY, current.Shift(1, 0), currentVal, visited)
	routes += rightCount
	visited = rightVisited

	downCount, downVisited := solveTo9(grid, maxX, maxY, current.Shift(0, 1), currentVal, visited)
	routes += downCount
	visited = downVisited

	leftCount, leftVisited := solveTo9(grid, maxX, maxY, current.Shift(-1, 0), currentVal, visited)
	routes += leftCount
	visited = leftVisited

	return routes, visited
}

func solveTo9WithRepeats(grid []string, maxX, maxY int, current elves.Point, lastVal int, visited []string, route [10]string) (int, []string) {
	currentStr := fmt.Sprintf("%s-from-%d", current, lastVal)

	if slices.Contains(visited, currentStr) {
		return 0, visited
	}

	visited = append(visited, currentStr)

	if !current.Valid(maxX, maxY) || grid[current.Y][current.X] == '.' {
		return 0, visited
	}

	currentVal := elves.Atoi(string(grid[current.Y][current.X]))
	if currentVal != lastVal+1 {
		return 0, visited
	}

	route[currentVal] = currentStr

	if grid[current.Y][current.X] == '9' {
		for i, point := range route {
			if i == 0 {
				continue
			}

			idx := slices.Index(visited, point)

			if idx != -1 {
				visited = slices.Delete(visited, idx, idx+1)
			}
		}

		return 1, visited
	}

	routes := 0

	upCount, upVisited := solveTo9WithRepeats(grid, maxX, maxY, current.Shift(0, -1), currentVal, visited, route)
	routes += upCount
	visited = upVisited

	rightCount, rightVisited := solveTo9WithRepeats(grid, maxX, maxY, current.Shift(1, 0), currentVal, visited, route)
	routes += rightCount
	visited = rightVisited

	downCount, downVisited := solveTo9WithRepeats(grid, maxX, maxY, current.Shift(0, 1), currentVal, visited, route)
	routes += downCount
	visited = downVisited

	leftCount, leftVisited := solveTo9WithRepeats(grid, maxX, maxY, current.Shift(-1, 0), currentVal, visited, route)
	routes += leftCount
	visited = leftVisited

	return routes, visited
}

func Part2(input string) (output string) {
	grid, maxX, maxY := elves.ParseGrid(input)

	starts := findStarts(grid)

	answer := 0

	for _, start := range starts {
		routes, _ := solveTo9WithRepeats(grid, maxX, maxY, start, -1, []string{}, [10]string{})
		answer += routes
	}

	return strconv.Itoa(answer)
}
