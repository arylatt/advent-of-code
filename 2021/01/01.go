package aoc202101

import (
	"strconv"
	"strings"
)

func Day1ParseInputFile(input string) []int {
	result := []int{}

	lines := strings.Split(input, "\r\n")

	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		result = append(result, val)
	}

	return result
}

func Part1(input string) (output string) {
	inputs := Day1ParseInputFile(input)

	result := 0
	previous := inputs[0]

	for _, val := range inputs[1:] {
		if val > previous {
			result++
		}

		previous = val
	}

	return strconv.Itoa(result)
}

func Day1GenerateWindows(inputs []int) []int {
	result := []int{}

	for i := 0; i < len(inputs)-2; i++ {
		result = append(result, (inputs[i] + inputs[i+1] + inputs[i+2]))
	}

	return result
}

func Part2(input string) (output string) {
	inputs := Day1ParseInputFile(input)

	inputs = Day1GenerateWindows(inputs)

	result := 0
	previous := inputs[0]

	for _, val := range inputs[1:] {
		if val > previous {
			result++
		}

		previous = val
	}

	return strconv.Itoa(result)
}
