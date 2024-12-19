package aoc202416

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	walls, start, end, maxX, maxY := buildGrid(elves.ParseGrid(input))
	heuristic, valid, weight := setupAStarFuncs(walls, end, maxX, maxY)

	_, gScore := elves.AStarSolve(start, end, heuristic, weight, elves.DirectionCardinals, valid)

	return strconv.Itoa(gScore[end])
}

func setupAStarFuncs(walls []elves.Point, end elves.Point, maxX, maxY int) (func(elves.Point) int, func(elves.Point) bool, func(elves.Point, elves.Point, map[elves.Point]elves.Point) int) {
	heuristic := elves.HeuristicManhattanDistance(end)
	valid := elves.PointValidAndNotInRange(walls, maxX, maxY)

	weight := func(current, next elves.Point, cameFrom map[elves.Point]elves.Point) int {
		previous, ok := cameFrom[current]

		direction := current.ShiftPos(previous.Invert())
		if !ok {
			direction = elves.DirectionRight
		}

		if current.ShiftPos(direction).Equals(next) {
			return 1
		}

		return 1001
	}

	return heuristic, valid, weight
}

func buildGrid(lines []string, maxX, maxY int) ([]elves.Point, elves.Point, elves.Point, int, int) {
	walls := []elves.Point{}
	start, end := elves.Point{}, elves.Point{}

	for y, line := range lines {
		for x, c := range line {
			p := elves.Point{X: x, Y: y}
			switch c {
			case '#':
				walls = append(walls, p)
			case 'S':
				start = p
			case 'E':
				end = p
			}
		}
	}

	return walls, start, end, maxX, maxY
}

func Part2(input string) (output string) {
	walls, start, end, maxX, maxY := buildGrid(elves.ParseGrid(input))
	heuristic, valid, weight := setupAStarFuncs(walls, end, maxX, maxY)

	solves, _ := elves.AStarSolveMulti(start, end, heuristic, weight, elves.DirectionCardinals, valid, -1)

	points := map[elves.Point]bool{}

	for _, solve := range solves {
		for _, p := range solve {
			points[p] = true
		}
	}

	return strconv.Itoa(len(points))
}
