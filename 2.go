package main

import "github.com/arylatt/advent-of-code/sub"

func Day2Answer(s sub.Sub) int {
	return s.Position.X * s.Aim
}

func Day2AnswerII(s sub.Sub) int {
	return s.Position.X * s.Position.Y
}

func Day2ParseInputFile(path string) []sub.Command {
	result := []sub.Command{}

	lines := ParseInputFile(path)

	for _, line := range lines {
		cmd := sub.Command{}
		cmd.Parse(line)
		result = append(result, cmd)
	}

	return result
}

func Day2Exec(path string) int {
	cmds := Day2ParseInputFile(path)

	sub := &sub.Sub{}

	sub.ExecuteCommands(cmds)
	return Day2Answer(*sub)
}

func Day2ExecII(path string) int {
	cmds := Day2ParseInputFile(path)

	sub := &sub.Sub{}

	sub.ExecuteCommands(cmds)
	return Day2AnswerII(*sub)
}
