package main

import "testing"

func TestDay5Exec(t *testing.T) {
	expected := 5
	actual := Day5Exec("inputs/5_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay5ExecII(t *testing.T) {
	expected := 12
	actual := Day5ExecII("inputs/5_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
