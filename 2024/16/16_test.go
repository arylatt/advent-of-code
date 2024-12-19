package aoc202416

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "7036"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "45"
)

func TestPart1Sample(t *testing.T) {
	td, err := elves.TestInputsToTestData([]elves.TestInput{
		{
			FileName: elves.SampleFileName,
			Answer:   Sample1Answer,
		},
		{
			FileName: "testdata/sample2.txt",
			Answer:   "11048",
		},
	})

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2024", "16")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.TestInputsToTestData([]elves.TestInput{
		{
			FileName: elves.SampleFileName,
			Answer:   Sample2Answer,
		},
		{
			FileName: "testdata/sample2.txt",
			Answer:   "64",
		},
	})

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2024", "16")
}
