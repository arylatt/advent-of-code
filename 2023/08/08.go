package aoc202308

import (
	"regexp"
	"slices"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

const ExprDirections = `(\w+) = \((\w+)\, (\w+)\)`

var exprDirections = regexp.MustCompile(ExprDirections)

type Direction struct {
	Current string
	Left    string
	Right   string
}

func (d Direction) Next(directions map[string]Direction, choice byte) Direction {
	switch choice {
	case 'R':
		return directions[d.Right]
	case 'L':
		return directions[d.Left]
	}

	panic("at the disco")
}

func parseDirection(input string) Direction {
	matches := exprDirections.FindAllStringSubmatch(input, -1)

	return Direction{
		Current: matches[0][1],
		Left:    matches[0][2],
		Right:   matches[0][3],
	}
}

func parseInput(input string) (string, map[string]Direction) {
	lines := elves.SplitIntoLines(input)

	route := lines[0]
	directions := map[string]Direction{}

	for i := 1; i < len(lines); i++ {
		direction := parseDirection(lines[i])
		directions[direction.Current] = direction
	}

	return route, directions
}

func calculate(input string) string {
	count := 0
	route, directions := parseInput(input)
	current := directions["AAA"]

	for i := 0; i < len(route); i++ {
		count++

		current = current.Next(directions, route[i])

		if current.Current == "ZZZ" {
			return strconv.Itoa(count)
		}

		if i == len(route)-1 {
			i = -1
		}
	}

	panic("at the disco")
}

func calculate2(input string) string {
	route, directions := parseInput(input)
	current := []Direction{}

	for _, direction := range directions {
		if direction.Current[2] == 'A' {
			current = append(current, direction)
		}
	}

	counts := []int{}

	for _, me := range current {
		count := 0

		for i := 0; i < len(route); i++ {
			count++

			me = me.Next(directions, route[i])

			if me.Current[2] == 'Z' {
				counts = append(counts, int(count))
				break
			}

			if i == len(route)-1 {
				i = -1
			}
		}
	}

	slices.Sort[[]int](counts)

	steps, stepVal := counts[len(counts)-1], counts[len(counts)-1]

	for {
		found := true

		for _, num := range counts[0 : len(counts)-1] {
			if steps%num != 0 {
				found = false
				break
			}
		}

		if found {
			return strconv.Itoa(steps)
		}

		steps += stepVal
	}
}

func Part1(input string) (output string) {
	return calculate(input)
}

func Part2(input string) (output string) {
	return calculate2(input)
}
