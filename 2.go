package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CommandDirection int

const (
	Forward CommandDirection = iota
	Up
	Down
)

type Command struct {
	Direction CommandDirection
	Units     int
}

type ModelSub struct {
	Horizontal int
	Depth      int // for part 2, this is treated as 'aim'
	AimedDepth int
}

func (s *ModelSub) ApplyCommand(cmd Command) {
	switch cmd.Direction {
	case Forward:
		s.Horizontal += cmd.Units
		s.AimedDepth += (cmd.Units * s.Depth)
	case Up:
		s.Depth -= cmd.Units
	case Down:
		s.Depth += cmd.Units
	}
}

func (s *ModelSub) ApplyCommands(cmds []Command) {
	for _, cmd := range cmds {
		s.ApplyCommand(cmd)
	}
}

func (s *ModelSub) Day2Answer() int {
	return s.Horizontal * s.Depth
}

func (s *ModelSub) Day2AnswerII() int {
	return s.Horizontal * s.AimedDepth
}

func ParseCommand(input string) Command {
	cmdParts := strings.Split(input, " ")
	if len(cmdParts) != 2 {
		panic("unexpected parts length")
	}

	cmd := Command{}
	switch cmdParts[0] {
	case "forward":
		cmd.Direction = Forward
	case "up":
		cmd.Direction = Up
	case "down":
		cmd.Direction = Down
	}

	val, _ := strconv.Atoi(cmdParts[1])
	cmd.Units = val

	return cmd
}

func Day2ParseInputFile(path string) ([]Command, error) {
	result := []Command{}

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmd := ParseCommand(scanner.Text())
		result = append(result, cmd)
	}

	return result, nil
}

func Day2Exec(path string) int {
	cmds, err := Day2ParseInputFile(path)
	if err != nil {
		panic("error not expected here")
	}

	sub := &ModelSub{}

	sub.ApplyCommands(cmds)
	return sub.Day2Answer()
}

func Day2ExecII(path string) int {
	cmds, err := Day2ParseInputFile(path)
	if err != nil {
		panic("error not expected here")
	}

	sub := &ModelSub{}

	sub.ApplyCommands(cmds)
	return sub.Day2AnswerII()
}
