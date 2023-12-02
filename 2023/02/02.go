package aoc202302

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

const (
	exprGame     = `Game (?P<GameID>\d+)`
	exprCubeData = `(?P<Count>\d+) (?P<Color>(?:red|green|blue))`
)

var (
	ExprGame     = regexp.MustCompile(exprGame)
	ExprCubeData = regexp.MustCompile(exprCubeData)
)

type Game struct {
	ID     int
	Rounds []map[string]int
	Max    map[string]int
}

func (g Game) Valid(maxRed, maxGreen, maxBlue int) bool {
	if g.Max["red"] > maxRed || g.Max["blue"] > maxBlue || g.Max["green"] > maxGreen {
		return false
	}

	return true
}

func parseLine(input string) Game {
	g := Game{
		ID:     0,
		Rounds: []map[string]int{},
		Max:    map[string]int{},
	}

	gameData := strings.Split(input, ":")

	matches := ExprGame.FindAllStringSubmatch(gameData[0], 1)
	g.ID, _ = strconv.Atoi(matches[0][1])

	for _, round := range strings.Split(strings.TrimSpace(gameData[1]), ";") {
		roundData := map[string]int{}

		for _, draw := range strings.Split(strings.TrimSpace(round), ",") {
			matches := ExprCubeData.FindAllStringSubmatch(draw, 1)

			for _, match := range matches {
				roundData[match[2]], _ = strconv.Atoi(match[1])

				max, ok := g.Max[match[2]]
				if !ok || max < roundData[match[2]] {
					g.Max[match[2]] = roundData[match[2]]
				}
			}
		}

		g.Rounds = append(g.Rounds, roundData)
	}

	return g
}

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	count := 0

	for _, line := range lines {
		if g := parseLine(line); g.Valid(12, 13, 14) {
			count += g.ID
		}
	}

	return strconv.Itoa(count)
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	count := 0

	for _, line := range lines {
		g := parseLine(line)

		count += (g.Max["red"] * g.Max["blue"] * g.Max["green"])
	}

	return strconv.Itoa(count)
}
