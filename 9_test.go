package main

import "testing"

func TestDay9Exec(t *testing.T) {
	expected := 15
	actual := Day9Exec("inputs/9_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay9ExecII(t *testing.T) {
	expected := 1134
	actual := Day9ExecII("inputs/9_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
