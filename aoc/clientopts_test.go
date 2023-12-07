package aoc

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithCacheTimeout_Valid(t *testing.T) {
	c := &Client{}

	if assert.NoError(t, WithCacheTimeout(DefaultCacheTime)(c)) {
		assert.Equal(t, DefaultCacheTime, c.cacheTimeout)
	}
}

func TestWithCacheTimeout_Invalid(t *testing.T) {
	c := &Client{}

	if assert.ErrorIs(t, ErrInvalidCacheTimeout, WithCacheTimeout(time.Second)(c)) {
		assert.Zero(t, c.cacheTimeout)
	}
}

func TestWithHTTPClient(t *testing.T) {
	tests := []*http.Client{nil, http.DefaultClient, {}}

	for _, test := range tests {
		c := &Client{}

		if assert.NoError(t, WithHTTPClient(test)(c)) {
			assert.Equal(t, test, c.httpClient)
		}
	}
}

func TestWithSessionToken(t *testing.T) {
	tests := []struct {
		token    string
		expected *http.Cookie
	}{{"", nil},
		{"   ", nil},
		{"token", aocCookie("token")},
	}

	for _, test := range tests {
		c := &Client{}

		if assert.NoError(t, WithSessionToken(test.token)(c)) {
			assert.Equal(t, test.expected, c.sessionCookie)
		}
	}
}

func TestWithUserAgent(t *testing.T) {
	tests := []string{"", "  ", "ua"}

	for _, test := range tests {
		c := &Client{}

		if assert.NoError(t, WithUserAgent(test)(c)) {
			assert.Equal(t, strings.TrimSpace(test), c.userAgent)
		}
	}
}
