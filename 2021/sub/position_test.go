package sub

import (
	"reflect"
	"testing"
)

func TestParsePosition(t *testing.T) {
	input := "12,99"
	expected := Position{X: 12, Y: 99}
	actual := ParsePosition(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v does not match actual %v", expected, actual)
	}
}
