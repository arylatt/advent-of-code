package aoc202309

import (
	"regexp"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

const ExprNum = `-?\d+`

var exprNum = regexp.MustCompile(ExprNum)

func allZeroes(input []int) bool {
	for _, v := range input {
		if v != 0 {
			return false
		}
	}

	return true
}

func parseLine(input string) [][]int {
	nums := [][]int{}
	matches := exprNum.FindAllString(input, -1)

	nums = append(nums, []int{})

	for i, match := range matches {
		num, _ := strconv.Atoi(match)

		nums[0] = append(nums[0], num)

		if i == 0 {
			continue
		}

		for mod := 0; mod < len(nums); mod++ {
			if allZeroes(nums[mod]) {
				break
			}

			if mod+1 == len(nums) {
				nums = append(nums, []int{})

				for i := 0; i < len(nums[mod])-2; i++ {
					nums[mod+1] = append(nums[mod+1], 0)
				}
			}

			if i-mod-1 < 0 {
				break
			}

			diff := nums[mod][i-mod] - nums[mod][i-mod-1]
			nums[mod+1] = append(nums[mod+1], diff)
		}
	}

	return nums
}

func extrapolateForward(input string) [][]int {
	nums := parseLine(input)

	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] = append(nums[i], (nums[i][len(nums[i])-1] + nums[i+1][len(nums[i+1])-1]))
	}

	return nums
}

func extrapolateReverse(input string) [][]int {
	nums := parseLine(input)

	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] = append([]int{(nums[i][0] - nums[i+1][0])}, nums[i]...)
	}

	return nums
}

func Part1(input string) (output string) {
	count := 0
	lines := elves.SplitIntoLines(input)

	for _, line := range lines {
		parsed := extrapolateForward(line)
		count += parsed[0][len(parsed[0])-1]
	}

	return strconv.Itoa(count)
}

func Part2(input string) (output string) {
	count := 0
	lines := elves.SplitIntoLines(input)

	for _, line := range lines {
		parsed := extrapolateReverse(line)
		count += parsed[0][0]
	}

	return strconv.Itoa(count)
}
