package main

import (
	"strings"

	"github.com/arylatt/advent-of-code/sub"
)

func ParseVentPositions(input string) (start, end sub.Position) {
	data := strings.Split(input, " -> ")
	start = sub.ParsePosition(data[0])
	end = sub.ParsePosition(data[1])
	return
}

func GeneratePositionsAndDiagram(path string) (positions []sub.Position, diagram [][]int) {
	inputs := ParseInputFile(path)

	highX, highY := 0, 0

	for _, input := range inputs {
		p1, p2 := ParseVentPositions(input)

		if p1.X >= highX {
			highX = p1.X
		}
		if p2.X >= highX {
			highX = p2.X
		}
		if p1.Y >= highY {
			highY = p1.Y
		}
		if p2.Y >= highY {
			highY = p2.Y
		}

		positions = append(positions, p1, p2)
	}

	diagram = make([][]int, highY+1)
	for i := range diagram {
		diagram[i] = make([]int, highX+1)
	}

	return
}

func CalculateHorizontals(path string) (positions []sub.Position, diagram [][]int) {
	positions, diagram = GeneratePositionsAndDiagram(path)

	for i := 0; i < len(positions); i += 2 {
		start := positions[i]
		end := positions[i+1]

		if start.X != end.X && start.Y != end.Y {
			// Ignoring diagonals
			continue
		}

		if start.X != end.X {
			for x := start.X; x <= end.X; x++ {
				diagram[start.Y][x]++
			}
			for x := start.X; x >= end.X; x-- {
				diagram[start.Y][x]++
			}
		}

		if start.Y != end.Y {
			for y := start.Y; y <= end.Y; y++ {
				diagram[y][start.X]++
			}
			for y := start.Y; y >= end.Y; y-- {
				diagram[y][start.X]++
			}
		}
	}

	return
}

func CalculateDiagonals(positions []sub.Position, in [][]int) (out [][]int) {
	out = in

	for i := 0; i < len(positions); i += 2 {
		start := positions[i]
		end := positions[i+1]

		if start.X == end.X || start.Y == end.Y {
			// Ignoring horizontals + verticals
			continue
		}

		if start.X > end.X && start.Y > end.Y {
			for diff := 0; (start.X - diff) >= (end.X); diff++ {
				out[start.Y-diff][start.X-diff]++
			}
		}
		if start.X < end.X && start.Y < end.Y {
			for diff := 0; (start.X + diff) <= (end.X); diff++ {
				out[start.Y+diff][start.X+diff]++
			}
		}
		if start.X > end.X && start.Y < end.Y {
			for diff := 0; (start.X - diff) >= (end.X); diff++ {
				out[start.Y+diff][start.X-diff]++
			}
		}
		if start.X < end.X && start.Y > end.Y {
			for diff := 0; (start.X + diff) <= (end.X); diff++ {
				out[start.Y-diff][start.X+diff]++
			}
		}
	}

	return
}

func CalculateAnswer(diagram [][]int) (answer int) {
	for _, row := range diagram {
		for y := range row {
			if row[y] >= 2 {
				answer++
			}
		}
	}

	return
}

func Day5Exec(path string) (answer int) {
	_, diagram := CalculateHorizontals(path)
	answer = CalculateAnswer(diagram)

	return
}

func Day5ExecII(path string) (answer int) {
	positions, diagram := CalculateHorizontals(path)
	diagram = CalculateDiagonals(positions, diagram)
	answer = CalculateAnswer(diagram)

	return
}
