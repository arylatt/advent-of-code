package bingo

import (
	"reflect"
	"testing"
)

func TestGenerateCards(t *testing.T) {
	expected := []*Card{
		{
			Rows: [5]Row{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
		},
		{
			Rows: [5]Row{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
		},
	}

	inputs := []string{
		" 1 2 3 4 5",
		"6 7 8 9 10",
		"11 12 13 14 15",
		"16 17 18 19 20",
		"21 22 23 24 25",
		"",
		"1 2 3 4 5",
		"6 7 8 9 10",
		"11 12 13 14 15",
		"16 17 18 19 20",
		"21 22 23 24 25",
	}

	actual := GenerateCards(inputs)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestMarkCards(t *testing.T) {
	expected := &Card{
		Rows: [5]Row{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		marks: [5]mark{
			{true},
		},
	}

	test := &Card{
		Rows: [5]Row{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}}

	actual_1 := MarkCards([]*Card{test}, 1)
	actual_26 := MarkCards([]*Card{test}, 26)

	if len(actual_1) != 1 {
		t.Fatalf("Mark 1 did not return true")
	}

	if len(actual_26) != 0 {
		t.Fatalf("Mark 26 did not return false")
	}

	if !reflect.DeepEqual(expected, test) {
		t.Fatalf("Expected %v did not match actual %v", expected, test)
	}
}

func TestGetWinningCards(t *testing.T) {
	expected := []*Card{
		{
			Rows: [5]Row{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			marks: [5]mark{
				{true, true, true, true, true},
			},
		},
		{
			Rows: [5]Row{
				{1, 6, 11, 16, 21},
				{2, 7, 12, 17, 22},
				{3, 8, 13, 18, 23},
				{4, 9, 14, 19, 24},
				{5, 10, 15, 20, 25},
			},
			marks: [5]mark{
				{true},
				{true},
				{true},
				{true},
				{true},
			},
		},
	}

	input := []*Card{
		{
			Rows: [5]Row{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			marks: [5]mark{
				{true, true, true, true, true},
			},
		},
		{
			Rows: [5]Row{
				{1, 6, 11, 16, 21},
				{2, 7, 12, 17, 22},
				{3, 8, 13, 18, 23},
				{4, 9, 14, 19, 24},
				{5, 10, 15, 20, 25},
			},
			marks: [5]mark{
				{true},
				{true},
				{true},
				{true},
				{true},
			},
		},
		{},
	}

	actual := GetWinningCards(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestSumOfWinningNumbers(t *testing.T) {
	expected := 310

	input := &Card{
		Rows: [5]Row{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		marks: [5]mark{
			{true, true, true, true, true},
		},
	}

	actual := input.SumOfUnmarkedNumbers()
	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestGetOrderedWinningCards(t *testing.T) {
	expected := []*Card{
		{
			Rows: [5]Row{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			marks: [5]mark{
				{true, true, true, true, true},
			},
		},
		nil,
	}

	input := []*Card{
		{
			Rows: [5]Row{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			marks: [5]mark{
				{true, true, true, true, true},
			},
		},
		{},
	}

	actual := GetOrderedWinningCards(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
