package aoc202418

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	maxIterations := 1024
	if elves.IsSample() {
		maxIterations = 12
	}

	obstructions, maxX, maxY, start, end := buildGrid(input, maxIterations)

	weight := func(p1, p2 elves.Point) int { return 1 }
	heuristic := func(p elves.Point) int { return p.ManhattanDistance(end) }

	valid := func(p elves.Point) bool {
		if !p.Valid(maxX, maxY) {
			return false
		}

		for _, obstruction := range obstructions {
			if p.Equals(obstruction) {
				return false
			}
		}

		return true
	}

	path := elves.AStarSolve(start, end, heuristic, weight, elves.DirectionCardinals, valid)

	return strconv.Itoa(len(path) - 1)
}

func buildGrid(input string, maxIterations int) ([]elves.Point, int, int, elves.Point, elves.Point) {
	maxX, maxY := 70, 70
	if elves.IsSample() {
		maxX, maxY = 6, 6
	}

	start := elves.Point{}
	end := elves.Point{X: maxX, Y: maxY}

	obstructions := []elves.Point{}

	for i, instruction := range elves.SplitIntoLines(input) {
		if i >= maxIterations {
			break
		}

		instructionParts := strings.Split(instruction, ",")

		obstructions = append(obstructions, elves.Point{X: elves.Atoi(instructionParts[0]), Y: elves.Atoi(instructionParts[1])})
	}

	return obstructions, maxX, maxY, start, end
}

func Part2(input string) (output string) {
	coords := elves.SplitIntoLines(input)
	min, max := 1024, len(coords)-1

	if elves.IsSample() {
		min = 6
	}

	weight := func(p1, p2 elves.Point) int { return 1 }

	for min != max {
		mid := (min + max) / 2

		obstructions, maxX, maxY, start, end := buildGrid(input, mid)

		heuristic := func(p elves.Point) int { return p.ManhattanDistance(end) }

		valid := func(p elves.Point) bool {
			if !p.Valid(maxX, maxY) {
				return false
			}

			for _, obstruction := range obstructions {
				if p.Equals(obstruction) {
					return false
				}
			}

			return true
		}

		if len(elves.AStarSolve(start, end, heuristic, weight, elves.DirectionCardinals, valid)) == 0 {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return coords[min-1]
}
