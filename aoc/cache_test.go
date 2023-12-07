package aoc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheEntryExpired(t *testing.T) {
	tests := []struct {
		time     time.Time
		expected bool
	}{
		{
			time.Now().Add(-(time.Hour * 1)),
			true,
		},
		{
			time.Now().Add(time.Hour * 1),
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, (CacheEntry{Expires: test.time}).Expired())
	}
}
