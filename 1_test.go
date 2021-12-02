package main

import (
	"reflect"
	"testing"
)

func TestDay1ParseInputFile(t *testing.T) {
	expected := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	actual, err := Day1ParseInputFile("inputs/1_sample.txt")
	if err != nil {
		t.Fatalf("Got an error: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}

func TestDay1Exec(t *testing.T) {
	expected := 7
	actual := Day1Exec("inputs/1_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}

func TestDay1ExecII(t *testing.T) {
	expected := 5
	actual := Day1ExecII("inputs/1_sample.txt")

	if expected != actual {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}

func TestDay1GenerateWindows(t *testing.T) {
	expected := []int{607, 618, 618, 617, 647, 716, 769, 792}
	actual := Day1GenerateWindows([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v did not match actual %v", expected, actual)
	}
}
