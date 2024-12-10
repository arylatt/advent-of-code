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

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ParseGrid(input string) ([]string, int, int) {
	lines := SplitIntoLines(input)
	return lines, len(lines[0]) - 1, len(lines) - 1
}
