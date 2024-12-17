package aoc202417

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

const (
	ExprRegister = `Register ([ABC]): (\d+)`
)

var (
	exprRegister = regexp.MustCompile(ExprRegister)
)

func Part1(input string) (output string) {
	registers, program := parseProgram(input)

	return runProgram(registers, program)
}

func parseProgram(input string) (map[string]int, []string) {
	lines := elves.SplitIntoLines(input)
	registers := map[string]int{}
	instructions := []string{}

	for _, line := range lines {
		if matches := exprRegister.FindStringSubmatch(line); matches != nil {
			registers[matches[1]] = elves.Atoi(matches[2])
			continue
		}

		programParts := strings.Split(line, " ")
		instructions = strings.Split(programParts[1], ",")
	}

	return registers, instructions
}

func runProgram(registers map[string]int, program []string) string {
	instructionPointer := 0
	outputs := []string{}

	for instructionPointer < len(program) {
		switch program[instructionPointer] {
		case "0":
			numerator := registers["A"]
			denominator := getComboOperandValue(registers, program[instructionPointer+1])

			denominator = int(math.Pow(2, float64(denominator)))

			registers["A"] = numerator / denominator
			instructionPointer += 2
		case "1":
			registers["B"] = registers["B"] ^ elves.Atoi(program[instructionPointer+1])
			instructionPointer += 2
		case "2":
			registers["B"] = getComboOperandValue(registers, program[instructionPointer+1]) % 8
			instructionPointer += 2
		case "3":
			if registers["A"] != 0 {
				instructionPointer = elves.Atoi(program[instructionPointer+1])
			} else {
				instructionPointer += 2
			}
		case "4":
			registers["B"] = registers["B"] ^ registers["C"]
			instructionPointer += 2
		case "5":
			outputs = append(outputs, strconv.Itoa(getComboOperandValue(registers, program[instructionPointer+1])%8))
			instructionPointer += 2
		case "6":
			numerator := registers["A"]
			denominator := getComboOperandValue(registers, program[instructionPointer+1])

			denominator = int(math.Pow(2, float64(denominator)))

			registers["B"] = numerator / denominator
			instructionPointer += 2
		case "7":
			numerator := registers["A"]
			denominator := getComboOperandValue(registers, program[instructionPointer+1])

			denominator = int(math.Pow(2, float64(denominator)))

			registers["C"] = numerator / denominator
			instructionPointer += 2
		}
	}

	return strings.Join(outputs, ",")
}

func getComboOperandValue(registers map[string]int, operand string) int {
	switch operand {
	case "4":
		return registers["A"]
	case "5":
		return registers["B"]
	case "6":
		return registers["C"]
	default:
		return elves.Atoi(operand)
	}
}

func Part2(input string) string {
	_, program := parseProgram(input)

	aVal := 0
	for i := len(program) - 1; i >= 0; i-- {
		aVal <<= 3

		for !slices.Equal(strings.Split(runProgram(map[string]int{"A": aVal, "B": 0, "C": 0}, program), ","), program[i:]) {
			aVal++
		}
	}

	return strconv.Itoa(aVal)
}
