package aoc202407

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	equations := elves.SplitIntoLines(input)
	answer := 0

	for _, equation := range equations {
		answerAndComponents := strings.Split(equation, ": ")
		expectedAnswer := elves.Atoi(answerAndComponents[0])
		components := strings.Split(answerAndComponents[1], " ")

		combinations := (1 << (len(components) - 1))
		operators := make([][]string, combinations)

		for i := 0; i < combinations; i++ {
			operators[i] = make([]string, len(components)-1)

			for j := range len(components) - 1 {
				if i&(1<<j) == 0 {
					operators[i][j] = "*"
					continue
				}

				operators[i][j] = "+"
			}
		}

		for _, operatorSet := range operators {
			total := elves.Atoi(components[0])

			for i, operator := range operatorSet {
				if operator == "+" {
					total = total + elves.Atoi(components[i+1])
				}

				if operator == "*" {
					total = total * elves.Atoi(components[i+1])
				}

				if total > expectedAnswer {
					break
				}
			}

			if total == expectedAnswer {
				answer += total
				break
			}
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	return
}
