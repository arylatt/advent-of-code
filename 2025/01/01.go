package aoc202501

import (
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func crackSafe(instructions []string, answerFunc func(int), part2Func func(int, int)) {
	position := 50

	if part2Func == nil {
		part2Func = func(int, int) {}
	}

	for _, instruction := range instructions {
		value := elves.Atoi(instruction[1:])
		origVal := value

		for value >= 100 {
			value -= 100
		}

		switch instruction[0] {
		case 'L':
			part2Func(position, -origVal)
			position -= value
		case 'R':
			part2Func(position, origVal)
			position += value
		}

		if position < 0 {
			position = 100 + position
		} else if position > 99 {
			position = position - 100
		}

		answerFunc(position)
	}
}

func Part1(input string) (output string) {
	answer := 0

	crackSafe(elves.SplitIntoLines(input), func(pos int) {
		if pos == 0 {
			answer++
		}
	}, nil)

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	answer := 0

	crackSafe(elves.SplitIntoLines(input), func(pos int) {
		if pos == 0 {
			answer++
		}
	}, func(start int, move int) {
		absMove := elves.Abs(move)

		hundreds := int(absMove / 100)
		if hundreds > 0 {
			answer += hundreds
			move = move % 100
		}

		if start == 0 {
			return
		}

		if start+move > 100 || start+move < 0 {
			answer++
		}
	})

	return strconv.Itoa(answer)
}
