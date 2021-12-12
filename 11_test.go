package main

import "testing"

func TestDay11Exec(t *testing.T) {
	expected := 1656
	actual := Day11Exec("inputs/11_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}

func TestDay11ExecII(t *testing.T) {
	expected := 195
	actual := Day11ExecII("inputs/11_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
