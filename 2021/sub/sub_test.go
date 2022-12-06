package sub

import (
	"reflect"
	"testing"
)

func TestSubMove(t *testing.T) {
	commands := []Command{
		{
			Direction: Forward,
			Units:     5,
		},
		{
			Direction: Down,
			Units:     5,
		},
		{
			Direction: Up,
			Units:     1,
		},
		{
			Direction: Forward,
			Units:     1,
		},
	}

	expected := Sub{
		Position: Position{
			X: 6,
			Y: 4,
		},
		Aim: 4,
	}

	actual := Sub{}

	actual.ExecuteCommands(commands)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
