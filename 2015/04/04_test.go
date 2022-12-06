package aoc201504

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "Fill_In_Sample_2_Answer"
)

func TestPart1Sample(t *testing.T) {
	td := elves.TestData{
		"abcdef":  "609043",
		"pqrstuv": "1048970",
	}

	elves.TestSample(t, td, Part1)
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, "2015", "4")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, "2015", "4")
}
