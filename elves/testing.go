package elves

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var SampleFileName = "testdata/sample.txt"

type TestData map[string]string

func SampleFileToTestData(answer string) (td TestData, err error) {
	file, err := os.ReadFile(SampleFileName)
	if err != nil {
		return
	}

	td = TestData{
		string(file): answer,
	}

	return
}

func TestSample(t *testing.T, td TestData, f func(string) string) {
	t.Helper()

	for input, output := range td {
		assert.Equal(t, output, f(input))
	}
}

func TestReal(t *testing.T, f func(string) string, level int, year, day string) {
	t.Helper()

	input, err := GetAOCInput(year, day)

	if assert.NoError(t, err) {
		result := f(input)

		t.Logf("Result: '%s'", result)

		// ok, err := PostAOCAnswer(year, day, level, result)
		ok, err := PostAOCAnswer(year, day, level, result)
		assert.True(t, ok, "If this is false, AOC said no...")
		assert.NoError(t, err, "AOC submit error")
	}
}
