package main

import (
	"bufio"
	"os"
)

// Reusable helpers for AoC

// ParseInputFile will read in a given path and return
// a string array of lines within it
func ParseInputFile(path string) []string {
	output := []string{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}
