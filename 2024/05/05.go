package aoc202405

import (
	"slices"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	beforeMap, afterMap := map[string][]string{}, map[string][]string{}
	buildingMaps := true
	answer := 0

	for _, line := range lines {
		if !strings.Contains(line, "|") {
			buildingMaps = false
		}

		if buildingMaps {
			vals := strings.Split(line, "|")

			before, after := vals[0], vals[1]

			if _, ok := beforeMap[before]; !ok {
				beforeMap[before] = []string{}
			}

			if _, ok := afterMap[after]; !ok {
				afterMap[after] = []string{}
			}

			beforeMap[before] = append(beforeMap[before], after)
			afterMap[after] = append(afterMap[after], before)

			continue
		}

		pages := strings.Split(line, ",")
		valid := true

		for i := range pages {
			if !valid {
				break
			}

			for j := range pages {
				if !valid {
					break
				}

				if i == j {
					continue
				}

				before, after := "", ""

				if i < j {
					before, after = pages[i], pages[j]
				} else {
					before, after = pages[j], pages[i]
				}

				if mustBeBefore, ok := beforeMap[after]; ok {
					if slices.Contains(mustBeBefore, before) {
						valid = false
						break
					}
				}

				if mustBeAfter, ok := afterMap[before]; ok {
					if slices.Contains(mustBeAfter, after) {
						valid = false
						break
					}
				}
			}
		}

		if valid {
			answer += elves.Atoi(pages[(len(pages)-1)/2])
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	beforeMap, afterMap := map[string][]string{}, map[string][]string{}
	buildingMaps := true
	invalidPages := [][]string{}
	answer := 0

	for _, line := range lines {
		if !strings.Contains(line, "|") {
			buildingMaps = false
		}

		if buildingMaps {
			vals := strings.Split(line, "|")

			before, after := vals[0], vals[1]

			if _, ok := beforeMap[before]; !ok {
				beforeMap[before] = []string{}
			}

			if _, ok := afterMap[after]; !ok {
				afterMap[after] = []string{}
			}

			beforeMap[before] = append(beforeMap[before], after)
			afterMap[after] = append(afterMap[after], before)

			continue
		}

		pages := strings.Split(line, ",")
		valid := true

		for i := range pages {
			if !valid {
				break
			}

			for j := range pages {
				if !valid {
					break
				}

				if i == j {
					continue
				}

				before, after := "", ""

				if i < j {
					before, after = pages[i], pages[j]
				} else {
					before, after = pages[j], pages[i]
				}

				if mustBeBefore, ok := beforeMap[after]; ok {
					if slices.Contains(mustBeBefore, before) {
						valid = false
						invalidPages = append(invalidPages, pages)
						break
					}
				}

				if mustBeAfter, ok := afterMap[before]; ok {
					if slices.Contains(mustBeAfter, after) {
						valid = false
						invalidPages = append(invalidPages, pages)
						break
					}
				}
			}
		}
	}

	for _, pages := range invalidPages {
		slices.SortStableFunc(pages, func(before, after string) int {
			mustBeBefore, ok := beforeMap[after]
			if !ok {
				return 0
			}

			if slices.Contains(mustBeBefore, before) {
				return 1
			}

			return -1
		})

		slices.SortStableFunc(pages, func(before, after string) int {
			mustBeAfter, ok := afterMap[before]
			if !ok {
				return 0
			}

			if slices.Contains(mustBeAfter, after) {
				return 1
			}

			return -1
		})

		answer += elves.Atoi(pages[(len(pages)-1)/2])
	}

	return strconv.Itoa(answer)
}
