package aoc202210

import (
	"math"
	"strconv"
	"strings"
)

func Cycle(instructions []string, f func(value, cycle int)) {
	cycles := []int{1}

	for _, instruction := range instructions {
		f(cycles[len(cycles)-1], len(cycles))

		cycles = append(cycles, cycles[len(cycles)-1])

		if strings.HasPrefix(instruction, "noop") {
			continue
		}

		f(cycles[len(cycles)-1], len(cycles))

		add, _ := strconv.Atoi(instruction[5:])
		cycles = append(cycles, cycles[len(cycles)-1]+add)

		if len(cycles) > 241 {
			break
		}
	}
}

func Part1(input string) (output string) {
	instructions := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")

	total := 0
	results := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}

	Cycle(instructions, func(value, cycle int) {
		if _, ok := results[cycle]; ok {
			total += value * cycle
		}
	})

	return strconv.Itoa(total)
}

func Part2(input string) (output string) {
	instructions := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	crt := [6][40]string{}

	Cycle(instructions, func(value, cycle int) {
		row := int(math.Ceil(float64(cycle)/40)) - 1
		column := ((cycle - 1 - (row * 40)) % 40)

		sprite := [3]int{value - 1, value, value + 1}
		for _, val := range sprite {
			if val == column {
				crt[row][column] = "#"
				break
			}
		}

		if crt[row][column] != "#" {
			crt[row][column] = "."
		}
	})

	for _, row := range crt {
		for _, col := range row {
			output += col
		}
		output += "\n"
	}

	return strings.TrimSpace(output)
}
