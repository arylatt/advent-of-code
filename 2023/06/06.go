package aoc202306

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

type (
	Race struct {
		Time, Distance int
	}
)

const (
	ExprNum = `\d+`
)

var (
	exprNum = regexp.MustCompile(ExprNum)
)

func (r Race) checkWin(hold int) bool {
	return r.Distance < (hold * (r.Time - hold))
}

func (r Race) LazyWin() int {
	winConditions := 0

	hold := r.Time - 1
	for hold != 0 {
		win := r.checkWin(hold)
		if !win && winConditions > 0 {
			return winConditions
		}

		if win {
			winConditions++
		}

		hold--
	}

	return winConditions
}

func parseRaces(input []string) []Race {
	races := []Race{}

	times := exprNum.FindAllString(input[0], -1)
	distances := exprNum.FindAllString(input[1], -1)

	for i, timeStr := range times {
		time, _ := strconv.Atoi(timeStr)
		distance, _ := strconv.Atoi(distances[i])

		races = append(races, Race{Time: time, Distance: distance})
	}

	return races
}

func parseRacesPart2(input []string) Race {
	times := exprNum.FindAllString(input[0], -1)
	distances := exprNum.FindAllString(input[1], -1)

	time, _ := strconv.Atoi(strings.Join(times, ""))
	distance, _ := strconv.Atoi(strings.Join(distances, ""))

	return Race{Time: time, Distance: distance}
}

func racePart1(input []string) int {
	races := parseRaces(input)

	count := 0
	for _, race := range races {
		winConds := race.LazyWin()
		if winConds == 0 {
			continue
		}

		if count == 0 {
			count = winConds
			continue
		}

		count *= winConds
	}

	return count
}

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	return strconv.Itoa(racePart1(lines))
}

func Part2(input string) (output string) {
	return strconv.Itoa(parseRacesPart2(elves.SplitIntoLines(input)).LazyWin())
}
