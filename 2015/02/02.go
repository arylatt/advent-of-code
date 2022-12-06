package aoc201502

import (
	"sort"
	"strconv"
	"strings"
)

type Dimension struct {
	Length int
	Width  int
	Height int
}

func ParseDimensions(input string) Dimension {
	if input == "" {
		return Dimension{}
	}

	dimensions := strings.Split(input, "x")
	dimension := Dimension{}

	dimension.Length, _ = strconv.Atoi(dimensions[0])
	dimension.Width, _ = strconv.Atoi(dimensions[1])
	dimension.Height, _ = strconv.Atoi(dimensions[2])

	return dimension
}

func Part1(input string) (output string) {
	sheets := strings.Split(input, "\n")
	total := 0

	for _, sheet := range sheets {
		sheetDimensions := ParseDimensions(sheet)

		lengthWidth := (sheetDimensions.Length * sheetDimensions.Width)
		widthHeight := (sheetDimensions.Width * sheetDimensions.Height)
		heightLength := (sheetDimensions.Height * sheetDimensions.Length)

		areas := sort.IntSlice([]int{lengthWidth, widthHeight, heightLength})
		areas.Sort()

		total += (2*lengthWidth + 2*widthHeight + 2*heightLength) // Surface Area
		total += areas[0]                                         // Slack
	}

	return strconv.Itoa(total)
}

func Part2(input string) (output string) {
	sheets := strings.Split(input, "\n")
	total := 0

	for _, sheet := range sheets {
		sheetDimensions := ParseDimensions(sheet)

		areas := sort.IntSlice([]int{sheetDimensions.Length, sheetDimensions.Width, sheetDimensions.Height})
		areas.Sort()

		total += (sheetDimensions.Length * sheetDimensions.Width * sheetDimensions.Height) // Bow
		total += (2*areas[0] + 2*areas[1])                                                 // Ribbon
	}

	return strconv.Itoa(total)
}
