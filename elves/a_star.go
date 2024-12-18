package elves

import (
	"math"
	"slices"
)

func AStarSolve(start, end Point, heuristic func(Point) int, weight func(Point, Point) int, directions []Point, valid func(Point) bool) []Point {
	open := []Point{start}
	cameFrom := map[Point]Point{}

	gScore := map[Point]int{start: 0}
	fScore := map[Point]int{start: heuristic(start)}

	for len(open) > 0 {
		current := getLowestFScore(open, fScore)
		if current.Equals(end) {
			return reconstructPath(cameFrom, current)
		}

		open = slices.DeleteFunc(open, func(p Point) bool { return p.Equals(current) })

		for _, neighbour := range generateNeighbours(current, directions, valid) {
			tentativeGScore := gScore[current] + weight(current, neighbour)

			if tentativeGScore < getGScore(gScore, neighbour) {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = gScore[neighbour] + heuristic(neighbour)

				open = addNeighbour(open, neighbour)
			}
		}
	}

	return []Point{}
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
