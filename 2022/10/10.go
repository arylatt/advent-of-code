package aoc202210

import (
	"strconv"
	"strings"
)

func Part1(input string) (output string) {
	instructions := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	cycles := []int{1}
	results := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}

	for _, instruction := range instructions {
		if _, ok := results[len(cycles)]; ok {
			results[len(cycles)] = cycles[len(cycles)-1] * len(cycles)
		}

		cycles = append(cycles, cycles[len(cycles)-1])

		if _, ok := results[len(cycles)]; ok {
			results[len(cycles)] = cycles[len(cycles)-1] * len(cycles)
		}

		if strings.HasPrefix(instruction, "noop") {
			continue
		}

		add, _ := strconv.Atoi(instruction[5:])
		cycles = append(cycles, cycles[len(cycles)-1]+add)

		if len(cycles) > 221 {
			break
		}
	}

	total := 0
	for _, v := range results {
		total += v
	}

	return strconv.Itoa(total)
}

func Part2(input string) (output string) {
	return
}
