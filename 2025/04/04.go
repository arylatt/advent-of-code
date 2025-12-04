package aoc202504

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	answer := 0
	paperRolls := []elves.Point{}

	floor := elves.SplitInto2DArrayWithFunc(input, func(p elves.Point, r rune) {
		if r == '@' {
			paperRolls = append(paperRolls, p)
		}
	})

	maxX, maxY := len(floor[0])-1, len(floor)-1

	for _, roll := range paperRolls {
		blocks := 0

		for i, dir := range elves.DirectionAll {
			if (len(elves.DirectionAll) - 3) == (i - blocks) {
				break
			}

			check := roll.ShiftPos(dir)

			if !check.Valid(maxX, maxY) {
				continue
			}

			if floor[check.Y][check.X] == '@' {
				blocks++
			}

			if blocks > 3 {
				break
			}
		}

		if blocks < 4 {
			answer++
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	answer := 0
	paperRolls := []elves.Point{}

	floor := elves.SplitInto2DArrayWithFunc(input, func(p elves.Point, r rune) {
		if r == '@' {
			paperRolls = append(paperRolls, p)
		}
	})

	maxX, maxY := len(floor[0])-1, len(floor)-1

	for {
		removeIndexes := []int{}

		for i, roll := range paperRolls {
			blocks := 0

			for i, dir := range elves.DirectionAll {
				if (len(elves.DirectionAll) - 3) == (i - blocks) {
					break
				}

				check := roll.ShiftPos(dir)

				if !check.Valid(maxX, maxY) {
					continue
				}

				if floor[check.Y][check.X] == '@' {
					blocks++
				}

				if blocks > 3 {
					break
				}
			}

			if blocks < 4 {
				answer++
				removeIndexes = append([]int{i}, removeIndexes...)
				floor[roll.Y][roll.X] = '.'
			}
		}

		if len(removeIndexes) == 0 {
			break
		}

		for _, ridx := range removeIndexes {
			paperRolls = append(paperRolls[:ridx], paperRolls[ridx+1:]...)
		}
	}

	return strconv.Itoa(answer)
}
