package aoc202401

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

const (
	ListSeparator = "   "
)

func Part1(input string) (output string) {
	left, right := getSortedLists(input)
	answer := 0

	for i := range left {
		answer += int(math.Abs(float64(left[i] - right[i])))
	}

	return strconv.Itoa(answer)
}

func getSortedLists(input string) ([]int, []int) {
	left, right := []int{}, []int{}
	lines := elves.SplitIntoLines(input)

	for _, line := range lines {
		vals := strings.Split(line, ListSeparator)

		leftVal, _ := strconv.Atoi(vals[0])
		rightVal, _ := strconv.Atoi(vals[1])

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func Part2(input string) (output string) {
	left, right := getSortedLists(input)
	similarityMap := map[int]int{}

	for _, val := range left {
		similarityMap[val] = 0
	}

	for _, val := range right {
		if _, ok := similarityMap[val]; ok {
			similarityMap[val]++
		}
	}

	answer := 0

	for _, val := range left {
		answer += val * similarityMap[val]
	}

	return strconv.Itoa(answer)
}
