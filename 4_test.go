package main

import "testing"

func TestDay4Exec(t *testing.T) {
	expected := 4512
	actual := Day4Exec("inputs/4_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}

func TestDay4ExecII(t *testing.T) {
	expected := 1924
	actual := Day4ExecII("inputs/4_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}
