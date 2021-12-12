package main

import "testing"

func TestDay8Exec(t *testing.T) {
	expected := 26
	actual := Day8Exec("inputs/8_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay8ExecII(t *testing.T) {
	expected := 61229
	actual := Day8ExecII("inputs/8_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
