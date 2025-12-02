package aoc202502

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "1227775554"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "4174379265"
)

func TestCalculateNextInvalidId(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"10", "11"},
		{"95", "99"},
		{"998", "1010"},
	}

	for _, test := range tests {
		result := calculateNextInvalidId(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestPart1Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample1Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2025", "2")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2025", "2")
}
