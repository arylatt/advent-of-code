package elves

import (
	"math"
	"slices"
)

func AStarSolve(start, end Point, heuristic func(Point) int, weight func(Point, Point, map[Point]Point) int, directions []Point, valid func(Point) bool) ([]Point, map[Point]int) {
	solves, gScore := AStarSolveMulti(start, end, heuristic, weight, directions, valid, 1)

	return solves[0], gScore
}

func AStarSolveMulti(start, end Point, heuristic func(Point) int, weight func(Point, Point, map[Point]Point) int, directions []Point, valid func(Point) bool, limitSolves int) ([][]Point, map[Point]int) {
	open := []Point{start}
	cameFrom := map[Point]Point{}

	gScore := map[Point]int{start: 0}
	fScore := map[Point]int{start: heuristic(start)}

	solves := [][]Point{}

	for len(open) > 0 {
		current := getLowestFScore(open, fScore)
		if current.Equals(end) {
			solves = append(solves, reconstructPath(cameFrom, current))

			if len(solves) == limitSolves {
				return solves, gScore
			}
		}

		open = slices.DeleteFunc(open, func(p Point) bool { return p.Equals(current) })

		for _, neighbour := range generateNeighbours(current, directions, valid) {
			tentativeGScore := gScore[current] + weight(current, neighbour, cameFrom)

			if tentativeGScore < getGScore(gScore, neighbour) {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = gScore[neighbour] + heuristic(neighbour)

				open = addNeighbour(open, neighbour)
			}
		}
	}

	return solves, gScore
}

func getLowestFScore(open []Point, fScore map[Point]int) Point {
	lowest := open[0]
	lowestScore := fScore[lowest]

	for _, p := range open {
		if fScore[p] < lowestScore {
			lowest = p
			lowestScore = fScore[p]
		}
	}

	return lowest
}

func reconstructPath(cameFrom map[Point]Point, current Point) []Point {
	totalPath := []Point{current}

	current, ok := cameFrom[current]
	for ok {
		totalPath = append([]Point{current}, totalPath...)
		current, ok = cameFrom[current]
	}

	return totalPath
}

func generateNeighbours(p Point, directions []Point, valid func(Point) bool) []Point {
	neighbours := []Point{}

	for _, d := range directions {
		neighbour := p.ShiftPos(d)
		if !valid(neighbour) {
			continue
		}

		neighbours = append(neighbours, neighbour)
	}

	return neighbours
}

func addNeighbour(open []Point, neighbour Point) []Point {
	for _, p := range open {
		if p.Equals(neighbour) {
			return open
		}
	}

	return append(open, neighbour)
}

func getGScore(gScore map[Point]int, p Point) int {
	if score, ok := gScore[p]; ok {
		return score
	}

	return math.MaxInt
}

func HeuristicManhattanDistance(end Point) func(Point) int {
	return func(p Point) int {
		return p.ManhattanDistance(end)
	}
}
