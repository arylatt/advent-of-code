package aoc202206

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
)

const (
	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "Fill_In_Sample_2_Answer"
)

func TestPart1Sample(t *testing.T) {
	td := elves.TestData{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "7",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "5",
		"nppdvjthqldpwncqszvftbrmjlhg":      "6",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "10",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "11",
	}

	elves.TestSample(t, td, Part1)
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, "2022", "6")
}

func TestPart2Sample(t *testing.T) {
	td := elves.TestData{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "19",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "23",
		"nppdvjthqldpwncqszvftbrmjlhg":      "23",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "29",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "26",
	}

	elves.TestSample(t, td, Part2)
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, "2022", "6")
}
