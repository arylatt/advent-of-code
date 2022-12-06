package aoc202203

import (
	"strconv"
	"strings"
)

func LetterToInt(letter string) int {
	val := int64([]rune(letter)[0])

	if val >= 97 {
		return int(val - 96)
	}

	return int(val - 38)
}

func Part1(input string) (output string) {
	rucksacks := strings.Split(input, "\n")

	uniqueItems := []string{}

	for _, rucksack := range rucksacks {
		rucksack = strings.TrimSpace(rucksack)
		rucksackLen := len(rucksack)
		comp1, comp2 := rucksack[:rucksackLen/2], rucksack[rucksackLen/2:]

		comp1Map := map[string]bool{}
		for _, item := range comp1 {
			comp1Map[string(item)] = true
		}

		for _, item := range comp2 {
			if _, ok := comp1Map[string(item)]; ok {
				uniqueItems = append(uniqueItems, string(item))
				break
			}
		}
	}

	sum := 0
	for _, item := range uniqueItems {
		sum += LetterToInt(item)
	}

	return strconv.Itoa(sum)
}

func Part2(input string) (output string) {
	rucksacks := strings.Split(input, "\n")
	uniqueItems := []string{}

	for i := 0; i < len(rucksacks)-1; i += 3 {
		subRucksacks := rucksacks[i : i+3]
		r1Map, r2Map := map[string]bool{}, map[string]bool{}

		for _, item := range strings.TrimSpace(subRucksacks[0]) {
			r1Map[string(item)] = true
		}

		for _, item := range strings.TrimSpace(subRucksacks[1]) {
			if _, ok := r1Map[string(item)]; ok {
				r2Map[string(item)] = true
			}
		}

		for _, item := range strings.TrimSpace(subRucksacks[2]) {
			if _, ok := r2Map[string(item)]; ok {
				uniqueItems = append(uniqueItems, string(item))
				break
			}
		}
	}

	sum := 0
	for _, item := range uniqueItems {
		sum += LetterToInt(item)
	}

	return strconv.Itoa(sum)
}
