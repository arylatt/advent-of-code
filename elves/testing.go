package elves

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const SampleFileName = "testdata/sample.txt"

var isSample = false

type TestData map[string]string

type TestInput struct {
	FileName string
	Answer   string
}

func (t TestData) Join(td ...TestData) TestData {
	for _, test := range td {
		for input, answer := range test {
			t[input] = answer
		}
	}

	return t
}

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

func TestInputsToTestData(inputs []TestInput) (TestData, error) {
	testDatas := []TestData{}
	for _, input := range inputs {
		td, err := SampleFileToTestData(input.Answer, input.FileName)
		if err != nil {
			return TestData{}, err
		}

		testDatas = append(testDatas, td)
	}

	return testDatas[0].Join(testDatas[1:]...), nil
}

func TestSample(t *testing.T, td TestData, f func(string) string) {
	t.Helper()
	isSample = true

	for input, output := range td {
		assert.Equal(t, output, f(input))
	}
}

func TestReal(t *testing.T, f func(string) string, level int, year, day string) {
	t.Helper()
	isSample = false

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

func IsSample() bool {
	return isSample
}
