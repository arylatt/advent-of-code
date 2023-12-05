package aoc202305

import (
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type (
	Item string

	AlmanacRange struct {
		Source, Destination, Length int
	}

	AlmanacEntry struct {
		Previous, Source, Destination Item

		Ranges []AlmanacRange

		min int
		max int
	}
)

const (
	ItemSeed        Item = "seed"
	ItemSoil        Item = "soil"
	ItemFertilizer  Item = "fertilizer"
	ItemWater       Item = "water"
	ItemLight       Item = "light"
	ItemTemperature Item = "temperature"
	ItemHumidity    Item = "humidity"
	ItemLocation    Item = "location"

	ExprMapIdentifier = `(\w+)-to-(\w+) map:`
	ExprRanges        = `(\d+) (\d+) (\d+)`
	ExprNum           = `\d+`
)

var (
	exprMapIdentifier = regexp.MustCompile(ExprMapIdentifier)
	exprRanges        = regexp.MustCompile(ExprRanges)
	exprNum           = regexp.MustCompile(ExprNum)
)

func (e AlmanacEntry) Lookup(item int) (Item, int) {
	for _, r := range e.Ranges {
		max := r.Source + r.Length - 1
		if item >= r.Source && item <= max {
			return e.Destination, r.Destination + (item - r.Source)
		}
	}

	return e.Destination, item
}

func (e AlmanacEntry) InverseLookup(item int) (Item, int) {
	for _, r := range e.Ranges {
		max := r.Destination + r.Length - 1
		if item >= r.Destination && item <= max {
			return e.Previous, r.Source + (item - r.Destination)
		}
	}

	return e.Previous, item
}

func parseMap(input []string, previous Item) (AlmanacEntry, []string) {
	entry := AlmanacEntry{
		Previous: previous,
		Ranges:   []AlmanacRange{},
	}

	for lineNo, line := range input {
		if strings.TrimSpace(line) == "" {
			sort.Slice(entry.Ranges, func(i, j int) bool {
				return entry.Ranges[i].Destination < entry.Ranges[j].Destination
			})

			return entry, input[lineNo:]
		}

		if matches := exprMapIdentifier.FindAllStringSubmatch(line, -1); len(matches) != 0 {
			entry.Source, entry.Destination = Item(matches[0][1]), Item(matches[0][2])
			continue
		}

		matches := exprRanges.FindAllStringSubmatch(line, -1)

		destStart, _ := strconv.Atoi(matches[0][1])
		sourceStart, _ := strconv.Atoi(matches[0][2])
		iters, _ := strconv.Atoi(matches[0][3])

		entry.Ranges = append(entry.Ranges, AlmanacRange{Source: sourceStart, Destination: destStart, Length: iters})

		if entry.min == 0 || sourceStart < entry.min {
			entry.min = sourceStart
		}

		iters--

		if sourceStart+iters > entry.max {
			entry.max = sourceStart + iters
		}
	}

	panic("parseMap hit the final return")
}

func parseEntries(lines []string) map[Item]AlmanacEntry {
	lines = lines[2:]

	previousEntry := Item("")
	entries := map[Item]AlmanacEntry{}
	for len(lines) != 0 {
		entry, newLines := parseMap(lines, previousEntry)

		previousEntry = entry.Source

		entries[entry.Source] = entry
		if len(newLines) <= 1 {
			break
		}

		if strings.TrimSpace(newLines[0]) == "" {
			newLines = newLines[1:]
		}

		lines = newLines
	}

	return entries
}

func parseInputPart1(input string) (map[Item]AlmanacEntry, []int) {
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")

	seedMatches := exprNum.FindAllString(lines[0], -1)
	seeds := []int{}
	for _, seed := range seedMatches {
		s, _ := strconv.Atoi(seed)
		seeds = append(seeds, s)
	}

	return parseEntries(lines), seeds
}

func calculateLocation(entries map[Item]AlmanacEntry, pos int) int {
	next := ItemSeed

	for next != ItemLocation {
		next, pos = entries[next].Lookup(pos)
	}

	return pos
}

func Part1(input string) (output string) {
	entries, seeds := parseInputPart1(input)

	locations := []int{}
	for _, seed := range seeds {
		locations = append(locations, calculateLocation(entries, seed))
	}

	slices.Sort[[]int](locations)

	return strconv.Itoa(locations[0])
}

func parseInputPart2(input string) (map[Item]AlmanacEntry, []AlmanacRange) {
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")

	seedMatches := exprNum.FindAllString(lines[0], -1)
	seeds := []AlmanacRange{}
	for i := 0; i < len(seedMatches); i += 2 {
		start, _ := strconv.Atoi(seedMatches[i])
		length, _ := strconv.Atoi(seedMatches[i+1])
		seeds = append(seeds, AlmanacRange{Source: start, Length: length})
	}

	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].Source < seeds[j].Source
	})

	return parseEntries(lines), seeds
}

func validSeed(seed int, seeds []AlmanacRange) bool {
	for _, seedRange := range seeds {
		if seed >= seedRange.Source && seed <= seedRange.Source+(seedRange.Length-1) {
			return true
		}
	}

	return false
}

func inverseCalculateLocation(entries map[Item]AlmanacEntry, pos int) int {
	next := ItemHumidity

	for next != Item("") {
		next, pos = entries[next].InverseLookup(pos)
	}

	return pos
}

func Part2(input string) (output string) {
	entries, seeds := parseInputPart2(input)

	i := 0
	for {
		seed := inverseCalculateLocation(entries, i)
		if validSeed(seed, seeds) {
			return strconv.Itoa(i)
		}

		i++
	}
}
