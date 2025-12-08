package aoc202507

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func parseMap(input string) (elves.Point, map[elves.Point]int, [][]rune) {
	source := elves.Point{}
	splitters := map[elves.Point]int{}

	tachMap := elves.SplitInto2DArrayWithFunc(input, func(p elves.Point, r rune) {
		switch r {
		case 'S':
			source = p
		case '^':
			splitters[p] = 0
		}
	})

	return source, splitters, tachMap
}

func Part1(input string) (output string) {
	answer := 0
	source, splitters, tachMap := parseMap(input)
	maxX, maxY := len(tachMap[0])-1, len(tachMap)-1
	next := []elves.Point{source.Shift(0, 1)}

	for len(next) > 0 {
		nextPos := next[0]
		next = next[1:]

		if !nextPos.Valid(maxX, maxY) {
			continue
		}

		nextVal := tachMap[nextPos.Y][nextPos.X]

		if nextVal == '.' {
			next = append([]elves.Point{nextPos.Shift(0, 1)}, next...)
			continue
		}

		if nextVal != '^' {
			panic("we've hit something mental")
		}

		if splitters[nextPos] == 1 {
			continue // already processed
		}

		answer++
		splitters[nextPos] = 1

		next = append([]elves.Point{nextPos.Shift(-1, 0)}, append(next, nextPos.Shift(1, 0))...)
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	return
}
