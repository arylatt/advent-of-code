package aoc202415

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "10092"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "9021"
)

func TestPart1Sample(t *testing.T) {
	td, err := elves.TestInputsToTestData([]elves.TestInput{
		{FileName: "testdata/sample2.txt", Answer: "2028"},
		{FileName: elves.SampleFileName, Answer: Sample1Answer},
	})

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2024", "15")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.TestInputsToTestData([]elves.TestInput{
		{FileName: "testdata/sample3.txt", Answer: "618"},
		{FileName: elves.SampleFileName, Answer: Sample2Answer},
	})

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2024", "15")
}
