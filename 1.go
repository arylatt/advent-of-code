package main

import (
	"strconv"
)

func Day1ParseInputFile(path string) []int {
	result := []int{}

	lines := ParseInputFile(path)

	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		result = append(result, val)
	}

	return result
}

func Day1Exec(path string) int {
	inputs := Day1ParseInputFile(path)

	result := 0
	previous := inputs[0]

	for _, val := range inputs[1:] {
		if val > previous {
			result++
		}

		previous = val
	}

	return result
}

func Day1GenerateWindows(inputs []int) []int {
	result := []int{}

	for i := 0; i < len(inputs)-2; i++ {
		result = append(result, (inputs[i] + inputs[i+1] + inputs[i+2]))
	}

	return result
}

func Day1ExecII(path string) int {
	inputs := Day1ParseInputFile(path)

	inputs = Day1GenerateWindows(inputs)

	result := 0
	previous := inputs[0]

	for _, val := range inputs[1:] {
		if val > previous {
			result++
		}

		previous = val
	}

	return result
}
