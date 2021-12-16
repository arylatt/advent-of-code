package main

import "testing"

func TestDay13Exec(t *testing.T) {
	expected := 17
	actual := Day13Exec("inputs/13_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay13ExecII(t *testing.T) {
	expected := 17
	actual := Day13ExecII("inputs/13_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
