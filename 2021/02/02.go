package aoc202102

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/2021/sub"
)

func Day2Answer(s sub.Sub) int {
	return s.Position.X * s.Aim
}

func Day2AnswerII(s sub.Sub) int {
	return s.Position.X * s.Position.Y
}

func Day2ParseInputFile(input string) []sub.Command {
	result := []sub.Command{}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		cmd := sub.Command{}
		cmd.Parse(line)
		result = append(result, cmd)
	}

	return result
}

func Part1(input string) (output string) {
	cmds := Day2ParseInputFile(input)

	sub := &sub.Sub{}

	sub.ExecuteCommands(cmds)
	return strconv.Itoa(Day2Answer(*sub))
}

func Part2(input string) (output string) {
	cmds := Day2ParseInputFile(input)

	sub := &sub.Sub{}

	sub.ExecuteCommands(cmds)
	return strconv.Itoa(Day2AnswerII(*sub))
}
