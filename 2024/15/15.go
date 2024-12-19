package aoc202415

import (
	"slices"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	walls, boxes, start, instructions := buildGrid(input)

	for _, instruction := range instructions {
		direction := elves.Point{}
		switch instruction {
		case '<':
			direction.X = -1
		case '^':
			direction.Y = -1
		case '>':
			direction.X = 1
		case 'v':
			direction.Y = 1
		}

		nextEmpty := elves.Point{}
		next := start.ShiftPos(direction)
		boxesInWay := []int{}
		for {
			if slices.IndexFunc(walls, func(p elves.Point) bool {
				return p.Equals(next)
			}) != -1 {
				break
			}

			if boxIdx := slices.IndexFunc(boxes, func(p elves.Point) bool {
				return p.Equals(next)
			}); boxIdx != -1 {
				boxesInWay = append(boxesInWay, boxIdx)
				next = next.ShiftPos(direction)
				continue
			}

			nextEmpty = next
			break
		}

		if nextEmpty.Equals(elves.Point{}) {
			// blocked
			continue
		}

		for _, boxIdx := range boxesInWay {
			boxes[boxIdx] = boxes[boxIdx].ShiftPos(direction)
		}

		start = start.ShiftPos(direction)
	}

	answer := 0

	for _, box := range boxes {
		answer += box.X + (box.Y * 100)
	}

	return strconv.Itoa(answer)
}

func buildGrid(input string) (walls, boxes []elves.Point, start elves.Point, instructions string) {
	lines := elves.SplitIntoLines(input)
	gridBuilt := false

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '#':
				walls = append(walls, elves.Point{X: x, Y: y})
			case 'O':
				boxes = append(boxes, elves.Point{X: x, Y: y})
			case '@':
				start = elves.Point{X: x, Y: y}
			case '.':
				// do nothing.
			default:
				gridBuilt = true
			}

			if gridBuilt {
				break
			}
		}

		if gridBuilt {
			instructions = strings.Join(lines[y:], "")
			break
		}
	}

	return
}

func Part2(input string) (output string) {
	walls, boxes, start, instructions := buildWideGrid(input)

	for _, instruction := range instructions {
		direction := elves.Point{}
		switch instruction {
		case '<':
			direction.X = -1
		case '^':
			direction.Y = -1
		case '>':
			direction.X = 1
		case 'v':
			direction.Y = 1
		}

		nextEmpty := elves.Point{}
		next := []elves.Point{start.ShiftPos(direction)}
		boxesInWay := []int{}
		for {
			if slices.IndexFunc(walls, func(p [2]elves.Point) bool {
				for _, n := range next {
					if p[0].Equals(n) || p[1].Equals(n) {
						return true
					}
				}

				return false
			}) != -1 {
				break
			}

			moreBoxLenBefore := len(boxesInWay)

			for nextIdx, n := range next {
				for boxIdx, box := range boxes {
					if box[0].Equals(n) || box[1].Equals(n) {
						boxesInWay = append(boxesInWay, boxIdx)
					} else {
						continue
					}

					next[nextIdx] = n.ShiftPos(direction)

					if direction.X != 0 {
						next[nextIdx] = next[nextIdx].ShiftPos(direction)
						continue
					}

					if box[0].Equals(n) {
						next = append(next, box[1].ShiftPos(direction))
						continue
					}

					next = append(next, box[0].ShiftPos(direction))
				}
			}

			boxesInWay = slices.Compact(boxesInWay)
			if len(boxesInWay) > moreBoxLenBefore {
				continue
			}

			nextEmpty = next[0]
			break
		}

		if nextEmpty.Equals(elves.Point{}) {
			// blocked
			continue
		}

		for _, boxIdx := range boxesInWay {
			boxes[boxIdx][0] = boxes[boxIdx][0].ShiftPos(direction)
			boxes[boxIdx][1] = boxes[boxIdx][1].ShiftPos(direction)
		}

		start = start.ShiftPos(direction)
	}

	answer := 0

	for _, box := range boxes {
		answer += box[0].X + (box[0].Y * 100)
	}

	return strconv.Itoa(answer)
}

func buildWideGrid(input string) (walls, boxes [][2]elves.Point, start elves.Point, instructions string) {
	lines := elves.SplitIntoLines(input)
	gridBuilt := false

	for y, line := range lines {
		x := 0
		for _, char := range line {
			switch char {
			case '#':
				walls = append(walls, [2]elves.Point{{X: x, Y: y}, {X: x + 1, Y: y}})
			case 'O':
				boxes = append(boxes, [2]elves.Point{{X: x, Y: y}, {X: x + 1, Y: y}})
			case '@':
				start = elves.Point{X: x, Y: y}
			case '.':
				// do nothing.
			default:
				gridBuilt = true
			}

			x += 2

			if gridBuilt {
				break
			}
		}

		if gridBuilt {
			instructions = strings.Join(lines[y:], "")
			break
		}
	}

	return
}
