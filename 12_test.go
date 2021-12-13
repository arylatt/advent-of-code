package main

import "testing"

func TestDay12Exec(t *testing.T) {
	expected := 226
	actual := Day12Exec("inputs/12_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay12ExecII(t *testing.T) {
	expected := 0
	actual := Day12ExecII("inputs/12_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
