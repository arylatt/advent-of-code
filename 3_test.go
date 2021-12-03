package main

import "testing"

func TestDay3Exec(t *testing.T) {
	expected := 198
	actual := Day3Exec("inputs/3_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}

func TestDay3ExecII(t *testing.T) {
	expected := 230
	actual := Day3ExecII("inputs/3_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}
