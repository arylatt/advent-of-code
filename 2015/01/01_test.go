package aoc201501

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
)

func TestPart1Sample(t *testing.T) {
	td := elves.TestData{
		"(())":    "0",
		"()()":    "0",
		"(((":     "3",
		"(()(()(": "3",
		"))(((((": "3",
		"())":     "-1",
		"))(":     "-1",
		")))":     "-3",
		")())())": "-3",
	}

	elves.TestSample(t, td, Part1)
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, "2015", "1")
}

func TestPart2Sample(t *testing.T) {
	td := elves.TestData{
		")":     "1",
		"()())": "5",
	}

	elves.TestSample(t, td, Part2)
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, "2015", "1")
}
