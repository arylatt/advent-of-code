package aoc

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ErrInvalidCacheTimeout is returned if a cache
// timeout less than the DefaultCacheTime is specified.
var ErrInvalidCacheTimeout = fmt.Errorf("invalid cache time specified, must be at least %s seconds", DefaultCacheTime/time.Second)

type clientOpt func(c *Client) error

// WithCacheTimeout allows to customize the cache timeout for the client.
// It cannot be lower than DefaultCacheTimeout.
func WithCacheTimeout(timeout time.Duration) clientOpt {
	return func(c *Client) error {
		if timeout < DefaultCacheTime {
			return ErrInvalidCacheTimeout
		}

		c.cacheTimeout = timeout
		return nil
	}
}

// WithHTTPClient allows to specify a custom HTTP client to use.
func WithHTTPClient(client *http.Client) clientOpt {
	return func(c *Client) error {
		if client != nil {
			c.httpClient = client
		}

		return nil
	}
}

func aocCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		Domain:   ".adventofcode.com",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
}

// WithSessionToken allows to specify the token for the requests.
func WithSessionToken(token string) clientOpt {
	return func(c *Client) error {
		if strings.TrimSpace(token) != "" {
			c.sessionCookie = aocCookie(token)
		}

		return nil
	}
}

// WithUserAgent allows customizing the User-Agent header sent in
// HTTP requests.
func WithUserAgent(userAgent string) clientOpt {
	return func(c *Client) error {
		if strings.TrimSpace(userAgent) != "" {
			c.userAgent = userAgent
		}

		return nil
	}
}
