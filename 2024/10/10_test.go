package aoc202410

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "36"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "81"
)

func TestPart1Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample1Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2024", "10")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.TestInputsToTestData([]elves.TestInput{
		{FileName: "testdata/sample2.txt", Answer: "3"},
		{FileName: elves.SampleFileName, Answer: Sample2Answer},
	})

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2024", "10")
}
