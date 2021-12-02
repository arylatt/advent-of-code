package main

import (
	"reflect"
	"testing"
)

func TestParseInputFile(t *testing.T) {
	expected := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	actual := ParseInputFile("inputs/1_sample.txt")

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
