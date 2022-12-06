package aoc202107

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (output string) {
	inputs := strings.Split(input, "\r\n")
	inputs = strings.Split(inputs[0], ",")

	crabs := []int{}
	for _, in := range inputs {
		crab, _ := strconv.Atoi(in)
		crabs = append(crabs, crab)
	}

	sort.Ints(crabs)

	low := crabs[0]
	high := crabs[len(crabs)-1]

	fuelConsumption := []int{}

	for i := low; i <= high; i++ {
		fuelConsumed := 0
		for _, crab := range crabs {
			fuelConsumed += int(math.Abs(float64(crab - i)))
		}
		fuelConsumption = append(fuelConsumption, fuelConsumed)
	}

	sort.Ints(fuelConsumption)
	output = strconv.Itoa(fuelConsumption[0])

	return
}

func Part2(input string) (output string) {
	inputs := strings.Split(input, "\r\n")
	inputs = strings.Split(inputs[0], ",")

	crabs := []int{}
	for _, in := range inputs {
		crab, _ := strconv.Atoi(in)
		crabs = append(crabs, crab)
	}

	sort.Ints(crabs)

	low := crabs[0]
	high := crabs[len(crabs)-1]

	fuelConsumption := []int{}

	for i := low; i <= high; i++ {
		fuelConsumed := 0
		for _, crab := range crabs {
			for f := int(math.Abs(float64(crab - i))); f >= 1; f-- {
				fuelConsumed += f
			}
		}
		fuelConsumption = append(fuelConsumption, fuelConsumed)
	}

	sort.Ints(fuelConsumption)
	output = strconv.Itoa(fuelConsumption[0])

	return
}
