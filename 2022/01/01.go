package aoc202201

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (output string) {
	cumulative := 0

	for _, val := range strings.Split(input, "\n") {
		if strings.TrimSpace(val) == "" {
			if outputInt, _ := strconv.Atoi(output); outputInt < cumulative {
				output = strconv.Itoa(cumulative)
			}

			cumulative = 0
		}

		valInt, _ := strconv.Atoi(strings.TrimSpace(val))
		cumulative += valInt
	}

	return
}

func Part2(input string) (output string) {
	cumulative, totals := 0, []int{}

	for _, val := range strings.Split(input, "\n") {
		if strings.TrimSpace(val) == "" {
			totals = append(totals, cumulative)
			cumulative = 0
		}

		valInt, _ := strconv.Atoi(strings.TrimSpace(val))
		cumulative += valInt
	}

	sort.Ints(totals)

	outputCumulative := 0
	for _, i := range []int{1, 2, 3} {
		outputCumulative += totals[len(totals)-i]
	}

	return strconv.Itoa(outputCumulative)
}
