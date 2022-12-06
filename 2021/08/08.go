package aoc202108

import (
	"strconv"
	"strings"
)

func Part1(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")

	stringyBois1 := []string{}
	for _, input := range inputs {
		vals := strings.Split(input, " | ")[1]
		stringyBois1 = append(stringyBois1, strings.Split(vals, " ")...)
	}

	for _, val := range stringyBois1 {
		switch len(val) {
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 7:
			answer++
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	return
}
