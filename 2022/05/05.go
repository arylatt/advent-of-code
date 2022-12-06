package aoc202205

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	RegexpNumberList       = `(\s*?(?P<Digit>\d+)\s*?)`
	RegexpMoveInstructions = `move (?P<Count>\d+) from (?P<Start>\d+) to (?P<End>\d+)`
)

type Stacks map[string][]string

func Move(s Stacks, start, end string, count int) {
	for i := 0; i < count; i++ {
		moving := s[start][len(s[start])-1]
		s[start] = s[start][:len(s[start])-1]
		s[end] = append(s[end], moving)
	}
}

func MoveAdvanced(s Stacks, start, end string, count int) {
	moving := []string{}

	for i := 0; i < count; i++ {
		moving = append([]string{s[start][len(s[start])-1]}, moving...)
		s[start] = s[start][:len(s[start])-1]
	}

	s[end] = append(s[end], moving...)
}

func ParseMap(input []string) (stacks Stacks) {
	stacks = Stacks{}
	stackNumbersRegex := regexp.MustCompile(RegexpNumberList)
	stackNumbers := input[len(input)-1]
	matches := stackNumbersRegex.FindAllStringSubmatch(stackNumbers, -1)
	indexes := map[string]int{}

	for _, match := range matches {
		stacks[match[2]] = []string{}
		indexes[match[2]] = strings.Index(stackNumbers, match[2])
	}

	for _, crateRows := range input[:len(input)-1] {
		for stackNum, index := range indexes {
			if string(crateRows[index]) != " " {
				stacks[stackNum] = append([]string{string(crateRows[index])}, stacks[stackNum]...)
			}
		}
	}

	return
}

func ParseInput(input string) (crateMap []string, instructions []string) {
	input = strings.ReplaceAll(input, "\r", "")
	inputs := strings.Split(input, "\n")

	instruct := false
	for _, input := range inputs {
		if input == "" {
			instruct = true
			continue
		}

		if instruct {
			instructions = append(instructions, input)
		} else {
			crateMap = append(crateMap, input)
		}
	}

	return
}

func Execute(input string, moveFunc func(Stacks, string, string, int)) (output string) {
	crateMap, instructions := ParseInput(input)
	stacks := ParseMap(crateMap)
	instructionsRegexp := regexp.MustCompile(RegexpMoveInstructions)

	for _, instruction := range instructions {
		matches := instructionsRegexp.FindAllStringSubmatch(instruction, -1)
		count, _ := strconv.Atoi(matches[0][1])

		moveFunc(stacks, matches[0][2], matches[0][3], count)
	}

	for i := 1; i <= len(stacks); i++ {
		output += stacks[strconv.Itoa(i)][len(stacks[strconv.Itoa(i)])-1]
	}

	return
}

func Part1(input string) (output string) {
	return Execute(input, Move)
}

func Part2(input string) (output string) {
	return Execute(input, MoveAdvanced)
}
