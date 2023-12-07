package aoc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemberLastStarTime(t *testing.T) {
	tNow := time.Now().Truncate(time.Second)

	member := Member{LastStarTimestamp: tNow.Unix()}

	assert.Equal(t, tNow, member.LastStarTime())
}

func TestPartCompletionCompleted(t *testing.T) {
	tests := []struct {
		ts       int64
		idx      int
		expected bool
	}{
		{0, 0, false},
		{1, 0, false},
		{0, 1, false},
		{1, 1, true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, (PartCompletion{test.ts, test.idx}).Completed())
	}
}

func TestPartCompletionGetStarTime(t *testing.T) {
	tNow := time.Now().Truncate(time.Second)

	part := PartCompletion{GetStarTimestamp: tNow.Unix()}

	assert.Equal(t, tNow, part.GetStarTime())
}
