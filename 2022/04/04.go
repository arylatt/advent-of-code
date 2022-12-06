package aoc202204

import (
	"strconv"
	"strings"
)

func Part1(input string) (output string) {
	elvePairs := strings.Split(strings.TrimSpace(input), "\n")
	overlaps := 0

	for _, pair := range elvePairs {
		elves := strings.Split(pair, ",")
		elf1, elf2 := strings.Split(elves[0], "-"), strings.Split(elves[1], "-")

		e1s, _ := strconv.Atoi(strings.TrimSpace(elf1[0]))
		e1e, _ := strconv.Atoi(strings.TrimSpace(elf1[1]))
		e2s, _ := strconv.Atoi(strings.TrimSpace(elf2[0]))
		e2e, _ := strconv.Atoi(strings.TrimSpace(elf2[1]))

		if e1s <= e2s && e2e <= e1e || e2s <= e1s && e1e <= e2e {
			overlaps++
		}
	}

	return strconv.Itoa(overlaps)
}

func Part2(input string) (output string) {
	elvePairs := strings.Split(strings.TrimSpace(input), "\n")
	overlaps := 0

	for _, pair := range elvePairs {
		elves := strings.Split(pair, ",")
		elf1, elf2 := strings.Split(elves[0], "-"), strings.Split(elves[1], "-")

		e1s, _ := strconv.Atoi(strings.TrimSpace(elf1[0]))
		e1e, _ := strconv.Atoi(strings.TrimSpace(elf1[1]))
		e2s, _ := strconv.Atoi(strings.TrimSpace(elf2[0]))
		e2e, _ := strconv.Atoi(strings.TrimSpace(elf2[1]))

		if e1s <= e2s && e2e <= e1e || e2s <= e1s && e1e <= e2e || e1s <= e2s && e1e <= e2e && e1e >= e2s || e2s <= e1s && e2e <= e1e && e2e >= e1s {
			overlaps++
		}
	}

	return strconv.Itoa(overlaps)
}
