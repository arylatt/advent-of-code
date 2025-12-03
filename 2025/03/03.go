package aoc202503

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func solveGeneric(input string, batteries int) string {
	batteryBanks := elves.SplitIntoLines(input)
	answer := 0

	for _, bank := range batteryBanks {
		num := ""
		startIdx := 0

		for b := range batteries {
			for i := range 9 {
				if i == 9 {
					continue
				}

				idx := strings.Index(bank[startIdx:], strconv.Itoa(9-i))

				if idx == -1 {
					continue
				}

				idx += startIdx

				if idx >= len(bank)-(batteries-1-b) {
					continue
				}

				startIdx = idx + 1
				num += string(bank[idx])
				break
			}
		}

		answer += elves.Atoi(num)
	}

	return strconv.Itoa(answer)
}

func Part1(input string) (output string) {
	return solveGeneric(input, 2)
}

func Part2(input string) (output string) {
	return solveGeneric(input, 12)
}
