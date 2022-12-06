package elves

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// AOCInputURLFormat is the templated address for AOC inputs
	AOCInputURLFormat = "https://adventofcode.com/%s/day/%s/input"
)

var (
	// ErrNoSessionCookie is returned when the AOC_SESSION_COOKIE env var is empty
	ErrNoSessionCookie = errors.New("no session cookie value found in env var aoc_session_cookie")

	// ErrInvalidDay is returned when we get a 404 because the day is not valid
	ErrInvalidDay = errors.New("invalid day - server returned http 404")
)

func GetAOCInput(year, day string) (input string, err error) {
	url := fmt.Sprintf(AOCInputURLFormat, year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	sessionToken := os.Getenv("AOC_SESSION_COOKIE")
	if sessionToken == "" {
		err = ErrNoSessionCookie
		return
	}

	aocSession := &http.Cookie{
		Name:     "session",
		Value:    sessionToken,
		Path:     "/",
		Domain:   ".adventofcode.com",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	req.AddCookie(aocSession)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		err = ErrInvalidDay
		return
	}

	byteArr, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	input = string(byteArr)
	return
}
