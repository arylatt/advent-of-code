package aoc202104

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/2021/bingo"
)

func Part1(input string) (output string) {
	inputs := strings.Split(input, "\r\n")

	cards := bingo.GenerateCards(inputs[2:])
	numbers := strings.Split(inputs[0], ",")

	for _, val := range numbers {
		num, _ := strconv.Atoi(val)

		bingo.MarkCards(cards, num)

		winningCards := bingo.GetWinningCards(cards)
		if len(winningCards) != 0 {
			output = strconv.Itoa(num * winningCards[0].SumOfUnmarkedNumbers())
			return
		}
	}

	return
}

func Part2(input string) (output string) {
	inputs := strings.Split(input, "\r\n")

	cards := bingo.GenerateCards(inputs[2:])
	numbers := strings.Split(inputs[0], ",")

	lastNumberChange := 0
	lastWinningCard := []*bingo.Card{{}}
	winningCardIndexes := []int{}

	for _, val := range numbers {
		num, _ := strconv.Atoi(val)

		bingo.MarkCards(cards, num)

		winningCards := bingo.GetOrderedWinningCards(cards)
		for i, card := range winningCards {
			if card == nil {
				continue
			}

			winnerIndexed := false
			for _, winner := range winningCardIndexes {
				if i == winner {
					winnerIndexed = true
					break
				}
			}

			if winnerIndexed {
				continue
			}

			winningCardIndexes = append(winningCardIndexes, i)
			lastNumberChange = num

			copy(lastWinningCard, []*bingo.Card{card})
		}

		if len(cards) == len(bingo.GetWinningCards(cards)) {
			break
		}
	}

	output = strconv.Itoa(lastNumberChange * lastWinningCard[0].SumOfUnmarkedNumbers())

	return
}
