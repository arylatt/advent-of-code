package main

import "strings"

func Day8Exec(path string) (answer int) {
	inputs := ParseInputFile(path)

	stringyBois1 := []string{}
	for _, input := range inputs {
		vals := strings.Split(input, " | ")[1]
		stringyBois1 = append(stringyBois1, strings.Split(vals, " ")...)
	}

	for _, val := range stringyBois1 {
		switch len(val) {
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 7:
			answer++
		}
	}

	return
}

func Day8ExecII(path string) (answer int) {
	return
}
