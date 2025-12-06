package aoc202506

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	sums := [][]string{}
	answer := 0

	expr := regexp.MustCompile(`[^ ]+`)

	for _, line := range lines {
		matches := expr.FindAllString(line, -1)
		sums = append(sums, matches)
	}

	for i := range len(sums[0]) {
		operator := sums[len(sums)-1][i]
		sum := 0

		for j := range len(sums) - 1 {
			if j == 0 {
				sum = elves.Atoi(sums[j][i])
				continue
			}

			switch operator {
			case "+":
				sum += elves.Atoi(sums[j][i])
			case "*":
				sum *= elves.Atoi(sums[j][i])
			}
		}

		answer += sum
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	allChars := [][]string{}
	answer := 0

	for _, line := range lines {
		chars := strings.Split(line, "")
		allChars = append(allChars, chars)
	}

	bottomRow := len(allChars) - 1

	nums := []int{}

	for i := range len(allChars[0]) {
		i = len(allChars[0]) - 1 - i

		num := ""

		for j := range len(allChars) - 1 {
			if val := allChars[j][i]; val != " " {
				num += val
			}
		}

		if numInt := elves.Atoi(num); numInt != 0 {
			nums = append(nums, numInt)
		}

		if len(allChars[bottomRow]) > i && allChars[bottomRow][i] != " " {
			operator := allChars[bottomRow][i]
			sum := nums[0]

			for k := 1; k < len(nums); k++ {
				switch operator {
				case "+":
					sum += nums[k]
				case "*":
					sum *= nums[k]
				}
			}
			answer += sum
			nums = []int{}
		}
	}

	return strconv.Itoa(answer)
}
