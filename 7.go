package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day7Exec(path string) (answer int) {
	input := ParseInputFile(path)
	input = strings.Split(input[0], ",")

	crabs := []int{}
	for _, in := range input {
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
	answer = fuelConsumption[0]

	return
}

func Day7ExecII(path string) (answer int) {
	input := ParseInputFile(path)
	input = strings.Split(input[0], ",")

	crabs := []int{}
	for _, in := range input {
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
	answer = fuelConsumption[0]

	return
}
