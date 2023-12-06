package elves

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const SampleFileName = "testdata/sample.txt"

type TestData map[string]string

func SampleFileToTestData(answer string, fileName ...string) (TestData, error) {
	if len(fileName) == 0 {
		fileName = append(fileName, SampleFileName)
	}

	file, err := os.ReadFile(fileName[0])
	if err != nil {
		return TestData{}, err
	}

	return TestData{string(file): answer}, nil
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
		start := time.Now()
		result := f(input)

		t.Logf("Result: %q in %s", result, time.Since(start).String())

		// ok, err := PostAOCAnswer(year, day, level, result)
		ok, err := PostAOCAnswer(year, day, level, result)
		assert.True(t, ok, "If this is false, AoC said no...")
		assert.NoError(t, err, "AoC submit error")
	}
}
