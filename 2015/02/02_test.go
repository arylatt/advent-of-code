package aoc201502

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
)

func TestPart1Sample(t *testing.T) {
	td := elves.TestData{
		"2x3x4":  "58",
		"1x1x10": "43",
	}

	elves.TestSample(t, td, Part1)
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2015", "3")
}

func TestPart2Sample(t *testing.T) {
	td := elves.TestData{
		"2x3x4":  "34",
		"1x1x10": "14",
	}

	elves.TestSample(t, td, Part2)
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2015", "3")
}
