package main

import "testing"

func TestDay10Exec(t *testing.T) {
	expected := 26397
	actual := Day10Exec("inputs/10_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay10ExecII(t *testing.T) {
	expected := 288957
	actual := Day10ExecII("inputs/10_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
