package aoc202403

import (
	"regexp"
	"strconv"
)

const (
	ExprMul       = `mul\((\d{1,3})\,(\d{1,3})\)`
	ExprMulDoDont = `mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)`
)

var (
	exprMul       = regexp.MustCompile(ExprMul)
	exprMulDoDont = regexp.MustCompile(ExprMulDoDont)
)

func Part1(input string) (output string) {
	return strconv.Itoa(extractInstructionsAndCalculate(input))
}

func extractInstructionsAndCalculate(input string) int {
	matches := exprMul.FindAllStringSubmatch(input, -1)

	answer := 0

	for _, match := range matches {
		calc := [2]int{}

		calc[0], _ = strconv.Atoi(match[1])
		calc[1], _ = strconv.Atoi(match[2])

		answer += calc[0] * calc[1]
	}

	return answer
}

func extractInstructionsAndCalculateDoDont(input string) int {
	matches := exprMulDoDont.FindAllStringSubmatch(input, -1)

	answer := 0
	skip := false

	for _, match := range matches {
		if match[0] == "do()" {
			skip = false
			continue
		}

		if match[0] == "don't()" {
			skip = true
			continue
		}

		if skip {
			continue
		}

		calc := [2]int{}

		calc[0], _ = strconv.Atoi(match[1])
		calc[1], _ = strconv.Atoi(match[2])

		answer += calc[0] * calc[1]
	}

	return answer
}

func Part2(input string) (output string) {
	return strconv.Itoa(extractInstructionsAndCalculateDoDont(input))
}
