package main

import (
	"bufio"
	"os"
	"strconv"
)

func Day1ParseInputFile(path string) ([]int, error) {
	result := []int{}

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		result = append(result, val)
	}

	return result, nil
}

func Day1Exec(path string) int {
	inputs, err := Day1ParseInputFile(path)
	if err != nil {
		panic(err)
	}

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
	inputs, err := Day1ParseInputFile(path)
	if err != nil {
		panic(err)
	}

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
