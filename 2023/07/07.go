package aoc202307

import (
	"sort"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/elves"
)

var jValue = 11

type (
	Card rune

	Hand struct {
		Cards []Card

		counter map[Card]int
		str     string
	}
)

func (c Card) Value() int {
	switch c {
	case 'T':
		return 10
	case 'J':
		return jValue
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}

	val, _ := strconv.Atoi(string(c))
	return val
}

func (h Hand) String() string {
	return h.str
}

func (h Hand) jokers() int {
	if !IsJokerTime() {
		return 0
	}

	j := h.counter[Card('J')]
	return j
}

func (h Hand) FiveOfAKind() bool {
	return len(h.counter) == 1 || len(h.counter) == 2 && h.jokers() != 0
}

func (h Hand) FourOfAKind() bool {
	jokers := h.jokers()
	if len(h.counter) != 2 && !(len(h.counter) == 3 && jokers != 0) {
		return false
	}

	handCounter := h.mutate(jokers)

	for _, val := range handCounter {
		if val != 1 && val != 4 {
			return false
		}
	}

	return true
}

func (h Hand) FullHouse() bool {
	jokers := h.jokers()
	if len(h.counter) != 2 && !(len(h.counter) == 3 && jokers != 0) {
		return false
	}

	handCounter := h.mutate(jokers)

	for _, val := range handCounter {
		if val != 2 && val != 3 {
			return false
		}
	}

	return true
}

func (h Hand) ThreeOfAKind() bool {
	jokers := h.jokers()
	if len(h.counter) != 3 && !(len(h.counter) == 4 && jokers != 0) {
		return false
	}

	handCounter := h.mutate(jokers)

	for _, val := range handCounter {
		if val == 3 {
			return true
		}
	}

	return false
}

func (h Hand) TwoPair() bool {
	jokers := h.jokers()
	if len(h.counter) != 3 && !(len(h.counter) == 4 && jokers != 0) {
		return false
	}

	handCounter := h.mutate(jokers)

	for _, val := range handCounter {
		if val != 2 && val != 1 {
			return false
		}
	}

	return true
}

func (h Hand) OnePair() bool {
	return len(h.counter) == 4 || (len(h.counter) == 5 && h.jokers() != 0)
}

func (h Hand) Strength() int {
	if h.FiveOfAKind() {
		return 6
	}

	if h.FourOfAKind() {
		return 5
	}

	if h.FullHouse() {
		return 4
	}

	if h.ThreeOfAKind() {
		return 3
	}

	if h.TwoPair() {
		return 2
	}

	if h.OnePair() {
		return 1
	}

	return 0
}

func parseHand(input string) (Hand, int) {
	handAndBid := strings.Split(input, " ")
	bid, _ := strconv.Atoi(handAndBid[1])

	hand := Hand{Cards: []Card{}, counter: map[Card]int{}, str: handAndBid[0]}
	for _, card := range handAndBid[0] {
		c := Card(card)
		hand.Cards = append(hand.Cards, c)

		if _, ok := hand.counter[c]; ok {
			hand.counter[c]++
			continue
		}

		hand.counter[c] = 1
	}

	return hand, bid
}

func handLess(h1, h2 Hand) bool {
	strength1, strength2 := h1.Strength(), h2.Strength()
	if strength1 != strength2 {
		return strength1 < strength2
	}

	for i := 0; i < 5; i++ {
		cVal1, cVal2 := h1.Cards[i].Value(), h2.Cards[i].Value()
		if cVal1 != cVal2 {
			return cVal1 < cVal2
		}
	}

	panic("at the disco")
}

func handLessSort(hands []Hand) ([]Hand, func(i, j int) bool) {
	return hands, func(i, j int) bool {
		return handLess(hands[i], hands[j])
	}
}

func calculate(input string) string {
	lines := elves.SplitIntoLines(input)

	bids := map[string]int{}
	hands := []Hand{}

	for _, line := range lines {
		hand, bid := parseHand(line)

		hands = append(hands, hand)
		bids[hand.String()] = bid
	}

	sort.Slice(handLessSort(hands))

	count := 0

	for i, hand := range hands {
		count += ((i + 1) * bids[hand.String()])
	}

	return strconv.Itoa(count)
}

func Part1(input string) (output string) {
	return calculate(input)
}

func ActivateJoker() {
	jValue = 1
}

func IsJokerTime() bool {
	return jValue == 1
}

func (h Hand) mutate(jokers int) map[Card]int {
	handCounter := h.counter

	if jokers == 0 {
		return handCounter
	}

	delete(handCounter, Card('J'))

	cardToIncrease := Card('0')

	for c := range handCounter {
		if cardToIncrease == Card('0') {
			cardToIncrease = c
			continue
		}

		if handCounter[c] > handCounter[cardToIncrease] {
			cardToIncrease = c
		}
	}

	handCounter[cardToIncrease] += jokers

	return handCounter
}

func Part2(input string) (output string) {
	ActivateJoker()

	return calculate(input)
}
