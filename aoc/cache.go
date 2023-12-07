package aoc

import (
	"errors"
	"time"
)

// DefaultCacheTime is the default time for returning cached results
// instead of calling AoC again, as per the guidelines for their API.
const DefaultCacheTime = time.Minute * 15

// ErrCacheNoValue is returned if no value is found in the cache
// for the given key.
var ErrCacheNoValue = errors.New("no value found in cache")

// CacheEntry represents a cache entry with expiration.
type CacheEntry struct {
	Expires     time.Time
	Leaderboard Leaderboard
}

// Expired indicates if the given cache entry is expired.
func (c CacheEntry) Expired() bool {
	return c.Expires.Before(time.Now())
}
