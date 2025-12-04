package elves

import (
	"strconv"
	"strings"
)

func SplitIntoLines(input string) []string {
	inputs := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	ret := []string{}

	for _, input := range inputs {
		if strings.TrimSpace(input) != "" {
			ret = append(ret, input)
		}
	}

	return ret
}

func SplitInto2DArray(input string) [][]rune {
	return SplitInto2DArrayWithFunc(input, nil)
}

func SplitInto2DArrayWithFunc(input string, f func(Point, rune)) [][]rune {
	if f == nil {
		f = func(p Point, r rune) {}
	}

	inputs := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	ret := [][]rune{}

	for i, input := range inputs {
		for j, char := range input {
			if i >= len(ret) {
				ret = append(ret, []rune{})
			}

			ret[i] = append(ret[i], char)

			f(Point{X: j, Y: i}, char)
		}
	}

	return ret
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ParseGrid(input string) ([]string, int, int) {
	lines := SplitIntoLines(input)
	return lines, len(lines[0]) - 1, len(lines) - 1
}
