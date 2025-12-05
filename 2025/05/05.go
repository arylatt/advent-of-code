package aoc202505

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	ranges := [][2]int{}
	answer := 0

	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, end := elves.Atoi(parts[0]), elves.Atoi(parts[1])

			ranges = append(ranges, [2]int{start, end})

			continue
		}

		num := elves.Atoi(line)
		for _, freshRange := range ranges {
			if num >= freshRange[0] && num <= freshRange[1] {
				answer++
				break
			}
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	return
}
