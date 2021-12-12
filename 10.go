package main

import (
	"sort"
	"strings"
)

func Day10Exec(path string) (answer int) {
	inputs := ParseInputFile(path)

	points := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

	for _, line := range inputs {
		expected := []string{}
	char:
		for _, char := range strings.Split(line, "") {
			switch char {
			case "<":
				expected = append([]string{">"}, expected...)
			case "{":
				expected = append([]string{"}"}, expected...)
			case "[":
				expected = append([]string{"]"}, expected...)
			case "(":
				expected = append([]string{")"}, expected...)
			case expected[0]:
				expected = expected[1:]
			default:
				answer += points[char]
				break char
			}
		}
	}
	return
}

func Day10ExecII(path string) (answer int) {
	inputs := ParseInputFile(path)

	points := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	scores := []int{}

line:
	for _, line := range inputs {
		expected := []string{}
		for _, char := range strings.Split(line, "") {
			switch char {
			case "<":
				expected = append([]string{">"}, expected...)
			case "{":
				expected = append([]string{"}"}, expected...)
			case "[":
				expected = append([]string{"]"}, expected...)
			case "(":
				expected = append([]string{")"}, expected...)
			case expected[0]:
				expected = expected[1:]
			default:
				continue line
			}
		}

		if len(expected) != 0 {
			score := 0
			for _, char := range expected {
				score = (score * 5) + points[char]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	answer = scores[(len(scores)-1)/2]
	return
}
