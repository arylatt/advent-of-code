package aoc202301

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func numFromLine(line string) int {
	str := ""

	for _, c := range line {
		if c >= '0' && c <= '9' {
			str += string(c)
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			str += string(line[i])
			break
		}
	}

	num, _ := strconv.Atoi(str)
	return num
}

func stringReplace(line string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							strings.ReplaceAll(
								strings.ReplaceAll(
									strings.ReplaceAll(line,
										"one", "o1ne"),
									"two", "t2wo"),
								"three", "t3hree"),
							"four", "f4our"),
						"five", "f5ive"),
					"six", "s6ix"),
				"seven", "s7even"),
			"eight", "e8ight"),
		"nine", "n9ine")
}

func Part1(input string) (output string) {
	count := 0

	for _, line := range elves.SplitIntoLines(input) {
		count += numFromLine(line)
	}

	return strconv.Itoa(count)
}

func Part2(input string) (output string) {
	count := 0

	for _, line := range elves.SplitIntoLines(input) {
		count += numFromLine(stringReplace(line))
	}

	return strconv.Itoa(count)
}
