package aoc202206

import (
	"strconv"
	"strings"
)

func Execute(input string, count int) (output string) {
	inputs, index := strings.Split(input, ""), count
	buffer := inputs[:index]

	for {
		check := map[string]bool{}
		for _, bufItem := range buffer {
			check[bufItem] = true
		}

		if len(check) == count {
			break
		}

		index++
		buffer = inputs[index-count : index]
	}

	return strconv.Itoa(index)
}

func Part1(input string) (output string) {
	return Execute(input, 4)
}

func Part2(input string) (output string) {
	return Execute(input, 14)
}
