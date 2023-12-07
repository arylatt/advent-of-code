package aoc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	// LeaderBoardURLFormat is the format for constructing request URLs for private leaderboards.
	LeaderboardURLFormat = "https://adventofcode.com/%s/leaderboard/private/view/%s.json"

	// defaultUserAgent is the default UA header sent if a custom one is not set.
	defaultUserAgent = "Go-aoc-client/1.0 (https://github.com/arylatt/advent-of-code)"
)

var (
	// ErrBadResponseCode is returned when the call to AoC received a non-200 HTTP response code.
	ErrBadResponseCode = errors.New("non-200 http response received")

	// ErrSessionTokenMustBeProvided is returned if the client does not have a session token
	// after all options have been applied.
	ErrSessionTokenMustBeProvided = errors.New("a non-empty session token must be provided in the options")

	// ErrRedirectBlocked is returned if the call to AoC returns a redirect response.
	ErrRedirectBlocked = errors.New("redirect is blocked (invalid token?)")
)

// Client is the caching AoC client.
type Client struct {
	cacheTimeout          time.Duration
	httpClient            *http.Client
	leaderboardCache      map[string]CacheEntry
	leaderboardCacheMutex *sync.RWMutex
	leaderboardURLFormat  string
	sessionCookie         *http.Cookie
	userAgent             string
}

// New creates a new Client. Can be customized with the With* functions.
// At minimum, WithSessionToken must be provided.
func New(opts ...clientOpt) (*Client, error) {
	c := &Client{
		cacheTimeout:          DefaultCacheTime,
		httpClient:            http.DefaultClient,
		leaderboardCache:      map[string]CacheEntry{},
		leaderboardCacheMutex: &sync.RWMutex{},
		leaderboardURLFormat:  LeaderboardURLFormat,
		userAgent:             defaultUserAgent,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	if c.sessionCookie == nil {
		return nil, ErrSessionTokenMustBeProvided
	}

	c.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return ErrRedirectBlocked
	}

	return c, nil
}

func cacheKey(event, owner string) string {
	return fmt.Sprintf("%s/%s", event, owner)
}

func (c *Client) writeCache(event, owner string, leaderboard Leaderboard, expires time.Time) {
	key := cacheKey(event, owner)
	entry := CacheEntry{
		Expires:     expires,
		Leaderboard: leaderboard,
	}

	c.leaderboardCacheMutex.Lock()
	c.leaderboardCache[key] = entry
	c.leaderboardCacheMutex.Unlock()
}

// LeaderboardFromCache returns a leaderboard from the cache, if a cached version is present.
func (c *Client) LeaderboardFromCache(event, owner string) (CacheEntry, error) {
	key := cacheKey(event, owner)

	c.leaderboardCacheMutex.RLock()
	entry, ok := c.leaderboardCache[key]
	c.leaderboardCacheMutex.RUnlock()

	if !ok {
		return CacheEntry{}, fmt.Errorf("%w, key=%s", ErrCacheNoValue, key)
	}

	return entry, nil
}

// Leaderboard returns a leaderboard. If there is a cached, non-expired version then it will be returned.
// If there is not a cached value, or the cached value has expired, a new version is fetched.
func (c *Client) Leaderboard(event, owner string) (Leaderboard, error) {
	if cachedEntry, err := c.LeaderboardFromCache(event, owner); err == nil && !cachedEntry.Expired() {
		return cachedEntry.Leaderboard, nil
	}

	url := fmt.Sprintf(c.leaderboardURLFormat, event, owner)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Leaderboard{}, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.AddCookie(c.sessionCookie)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Leaderboard{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Leaderboard{}, fmt.Errorf("%w, statusCode=%d", ErrBadResponseCode, resp.StatusCode)
	}

	leaderboard := Leaderboard{}
	err = json.NewDecoder(resp.Body).Decode(&leaderboard)
	if err != nil {
		return leaderboard, err
	}

	go c.writeCache(event, owner, leaderboard, time.Now().Add(c.cacheTimeout))
	return leaderboard, nil
}
