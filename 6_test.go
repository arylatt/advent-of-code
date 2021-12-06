package main

import "testing"

func TestDay6Exec80(t *testing.T) {
	expected := 5934
	actual := Day6Exec("inputs/6_sample.txt", 80)

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay6Exec256(t *testing.T) {
	expected := 26984457539
	actual := Day6Exec("inputs/6_sample.txt", 256)

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
