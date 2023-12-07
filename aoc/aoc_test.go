package aoc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testSrv(validate func(r *http.Request)) *httptest.Server {
	leaderboard, err := os.ReadFile("testdata/leaderboard.json")
	if err != nil {
		panic(err)
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if validate != nil {
			validate(r)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(leaderboard)
	}))
}

func TestNew(t *testing.T) {
	srv := testSrv(nil)
	defer srv.Close()

	token := "token"
	c, err := New(WithHTTPClient(srv.Client()), WithSessionToken(token))

	if assert.NoError(t, err) {
		assert.Equal(t, DefaultCacheTime, c.cacheTimeout)
		assert.Equal(t, srv.Client(), c.httpClient)
		assert.NotNil(t, c.leaderboardCache)
		assert.NotNil(t, c.leaderboardCacheMutex)
		assert.Equal(t, LeaderboardURLFormat, c.leaderboardURLFormat)
		assert.Equal(t, token, c.sessionCookie.Value)
		assert.Equal(t, defaultUserAgent, c.userAgent)
	}
}

func TestCacheKey(t *testing.T) {
	event, owner := "event", "owner"
	expected := fmt.Sprintf("%s/%s", event, owner)

	assert.Equal(t, expected, cacheKey(event, owner))
}

func TestClientWriteCache(t *testing.T) {
	event, owner, expires := "event", "owner", time.Now().Add(time.Minute)
	leaderboard := Leaderboard{OwnerID: 1}

	c := &Client{
		leaderboardCache:      map[string]CacheEntry{},
		leaderboardCacheMutex: &sync.RWMutex{},
	}

	expected := CacheEntry{
		Expires:     expires,
		Leaderboard: leaderboard,
	}

	c.writeCache(event, owner, leaderboard, expires)

	assert.Equal(t, expected, c.leaderboardCache[cacheKey(event, owner)])
}

func TestClientLeaderboardFromCache(t *testing.T) {
	event, owner, expires := "event", "owner", time.Now().Add(time.Minute)
	leaderboard := Leaderboard{OwnerID: 1}

	expected := CacheEntry{
		Expires:     expires,
		Leaderboard: leaderboard,
	}

	c := &Client{
		leaderboardCache: map[string]CacheEntry{
			cacheKey(event, owner): expected,
		},
		leaderboardCacheMutex: &sync.RWMutex{},
	}

	actual, err := c.LeaderboardFromCache(event, owner)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestClientLeaderboard_UsesCache(t *testing.T) {
	event, owner, token := "event", "owner", "token"

	serverCalls := 0

	srv := testSrv(func(r *http.Request) {
		serverCalls++
	})

	defer srv.Close()

	expected := Leaderboard{
		OwnerID: 1,
	}

	c, _ := New(WithHTTPClient(srv.Client()), WithSessionToken(token))
	c.leaderboardURLFormat = fmt.Sprintf("%s/%%s/%%s", srv.URL)
	c.leaderboardCache[cacheKey(event, owner)] = CacheEntry{
		Expires:     time.Now().Add(time.Hour),
		Leaderboard: expected,
	}

	leaderboard, err := c.Leaderboard(event, owner)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, leaderboard)
		assert.Equal(t, 0, serverCalls)
	}
}

func TestClientLeaderboard_CacheExpiry(t *testing.T) {
	event, owner, token := "event", "owner", "token"

	serverCalls := 0

	srv := testSrv(func(r *http.Request) {
		serverCalls++

		cookie, err := r.Cookie("session")
		if assert.NoError(t, err) {
			assert.Equal(t, token, cookie.Value)
		}
	})

	defer srv.Close()

	c, _ := New(WithHTTPClient(srv.Client()), WithSessionToken(token))
	c.leaderboardURLFormat = fmt.Sprintf("%s/%%s/%%s", srv.URL)
	c.leaderboardCache[cacheKey(event, owner)] = CacheEntry{Expires: time.Now().Add(-time.Hour)}

	_, err := c.Leaderboard(event, owner)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, serverCalls)
	}
}
