package aoc201503

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
)

func TestPart1Sample(t *testing.T) {
	td := elves.TestData{
		">":          "2",
		"^>v<":       "4",
		"^v^v^v^v^v": "2",
	}

	elves.TestSample(t, td, Part1)
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, "2015", "3")
}

func TestPart2Sample(t *testing.T) {
	td := elves.TestData{
		"^v":         "3",
		"^>v<":       "3",
		"^v^v^v^v^v": "11",
	}

	elves.TestSample(t, td, Part2)
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, "2015", "3")
}
