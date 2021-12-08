package main

import "testing"

func TestDay7Exec(t *testing.T) {
	expected := 37
	actual := Day7Exec("inputs/7_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay7ExecII(t *testing.T) {
	expected := 168
	actual := Day7ExecII("inputs/7_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
