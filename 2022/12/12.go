package aoc202212

import (
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

type Grid struct {
	Grid       [][]string
	CurrentPos Point
	Moves      []Point
}

func Parse(input string) (g Grid) {
	g = Grid{CurrentPos: Point{0, 0}}

	rows := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	g.Grid = make([][]string, len(rows))

	for i, row := range rows {
		cols := strings.Split(row, "")
		g.Grid[i] = append(g.Grid[i], cols...)
	}

	return
}

func (g Grid) Walk(direction ...Point) Grid {
	if len(direction) != 0 {
		g.Moves = append(g.Moves, g.CurrentPos)
		g.CurrentPos = direction[0]
	}

	if g.Grid[g.CurrentPos.Y][g.CurrentPos.X] == "E" {
		return g
	}

	subGrids := []Grid{}
	for _, move := range g.Movements() {
		subGrids = append(subGrids, g.Walk(move))
	}

	sort.SliceStable(subGrids, func(i, j int) bool {
		return len(subGrids[i].Moves) < len(subGrids[j].Moves)
	})

	if len(subGrids) == 0 {
		return Grid{}
	}

	for _, grid := range subGrids {
		if len(grid.Moves) != 0 {
			return grid
		}
	}

	return Grid{}
}

func (g Grid) Movements() (moves []Point) {
	if p := (Point{g.CurrentPos.X, g.CurrentPos.Y - 1}); g.CurrentPos.Y != 0 && g.CanMove(p) {
		moves = append(moves, p)
	}

	if p := (Point{g.CurrentPos.X, g.CurrentPos.Y + 1}); g.CurrentPos.Y != len(g.Grid)-1 && g.CanMove(p) {
		moves = append(moves, p)
	}

	if p := (Point{g.CurrentPos.X - 1, g.CurrentPos.Y}); g.CurrentPos.X != 0 && g.CanMove(p) {
		moves = append(moves, p)
	}

	if p := (Point{g.CurrentPos.X + 1, g.CurrentPos.Y}); g.CurrentPos.X != len(g.Grid[0])-1 && g.CanMove(p) {
		moves = append(moves, p)
	}

	return
}

func (g Grid) CanMove(to Point) bool {
	for _, move := range g.Moves {
		if move.X == to.X && move.Y == to.Y {
			return false
		}
	}

	if g.Grid[g.CurrentPos.Y][g.CurrentPos.X] == "S" {
		return true
	}

	if g.Grid[g.CurrentPos.Y][g.CurrentPos.X] != "z" && g.Grid[to.Y][to.X] == "E" {
		return false
	}

	return int(g.Grid[g.CurrentPos.Y][g.CurrentPos.X][0])+1 >= int(g.Grid[to.Y][to.X][0])
}

func Part1(input string) (output string) {
	return strconv.Itoa(len(Parse(input).Walk().Moves))
}

func Part2(input string) (output string) {
	return
}
