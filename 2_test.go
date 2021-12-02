package main

import "testing"

func TestDay2Exec(t *testing.T) {
	expected := 150
	actual := Day2Exec("inputs/2_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}

func TestDay2ExecII(t *testing.T) {
	expected := 900
	actual := Day2ExecII("inputs/2_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}
