package sub

import (
	"reflect"
	"testing"
)

func TestCommandParse(t *testing.T) {
	inputs := []string{"forward 1", "up 1", "down 1"}
	expected := []Command{
		{
			Direction: Forward,
			Units:     1,
		},
		{
			Direction: Up,
			Units:     1,
		},
		{
			Direction: Down,
			Units:     1,
		},
	}

	actual := []Command{}

	for _, test := range inputs {
		cmd := Command{}
		cmd.Parse(test)
		actual = append(actual, cmd)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
