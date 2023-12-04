package aoc202304

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

const ExprNum = `\d+`

var exprNum = regexp.MustCompile(ExprNum)

type card struct {
	ID      int
	Winning map[int]bool
	My      []int

	SubScore int
}

func (c card) Score() int {
	score := 0

	for _, num := range c.My {
		_, ok := c.Winning[num]
		if !ok {
			continue
		}

		if score == 0 {
			score = 1
			continue
		}

		score *= 2
	}

	return score
}

func (c *card) Score2(idx int, cards []*card) {
	if c.SubScore != 0 {
		return
	}

	c.SubScore = 1
	wins := 0

	for _, num := range c.My {
		_, ok := c.Winning[num]
		if !ok {
			continue
		}

		cards[idx-1-wins].Score2(idx-1-wins, cards)
		c.SubScore += cards[idx-1-wins].SubScore
		wins++
	}
}

func parseCard(input string) card {
	idDataParts := strings.Split(input, ":")
	numberParts := strings.Split(idDataParts[1], "|")

	id, _ := strconv.Atoi(exprNum.FindString(idDataParts[0]))
	winningNumStrs := exprNum.FindAllString(numberParts[0], -1)
	myNumStrs := exprNum.FindAllString(numberParts[1], -1)

	winningNums := map[int]bool{}
	for _, str := range winningNumStrs {
		num, _ := strconv.Atoi(str)
		winningNums[num] = true
	}

	myNums := []int{}
	for _, str := range myNumStrs {
		num, _ := strconv.Atoi(str)
		myNums = append(myNums, num)
	}

	c := card{
		ID:      id,
		Winning: winningNums,
		My:      myNums,
	}

	return c
}

func parseCardP1(input string, score int) (card, int) {
	c := parseCard(input)
	return c, (score + c.Score())
}

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	count := 0

	for _, line := range lines {
		_, count = parseCardP1(line, count)
	}

	return strconv.Itoa(count)
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	cards := []*card{}

	for _, line := range lines {
		c := parseCard(line)
		cards = append(cards, &c)
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].ID > cards[j].ID
	})

	score := 0
	for i := 0; i < len(cards); i++ {
		cards[i].Score2(i, cards)
		score += cards[i].SubScore
	}

	return strconv.Itoa(score)
}
