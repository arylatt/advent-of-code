package elves

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	// AOCInputURLFormat is the templated address for AOC inputs
	AOCInputURLFormat = "https://adventofcode.com/%s/day/%s/input"

	// AOCAnswerURLFormat is the templated address for AOC inputs
	AOCAnswerURLFormat = "https://adventofcode.com/%s/day/%s/answer"

	ResponseIncorrect = "That's not the right answer"
	ResponseCooldown  = "You gave an answer too recently"
)

var (
	// ErrNoSessionCookie is returned when the AOC_SESSION_COOKIE env var is empty
	ErrNoSessionCookie = errors.New("no session cookie value found in env var aoc_session_cookie")

	// ErrInvalidDay is returned when we get a 404 because the day is not valid
	ErrInvalidDay = errors.New("invalid day - server returned http 404")

	// ErrUnexpectedResponseCode is returned when we get a non-200 response for answer submit
	ErrUnexpectedResponseCode = errors.New("unexpected response code from answer submission. expected http 200")

	// ErrUnexpectedResponsePage is returned when we can't match our expected strings in the response for answer submit
	ErrUnexpectedResponsePage = errors.New("unexpected response page from answer submission")
)

func GetAOCInput(year, day string) (input string, err error) {
	url := fmt.Sprintf(AOCInputURLFormat, year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	if err = aocSessionInject(req); err != nil {
		return
	}

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

func PostAOCAnswer(year, day string, level int, answer string) (ok bool, err error) {
	ok = true
	answerUrl := fmt.Sprintf(AOCAnswerURLFormat, year, day)

	if os.Getenv("AOC_SUBMIT_ANSWERS") == "" {
		return
	}

	body := url.Values{}
	body.Add("answer", answer)
	body.Add("level", strconv.Itoa(level))

	req, err := http.NewRequest(http.MethodPost, answerUrl, strings.NewReader(body.Encode()))
	if err != nil {
		return
	}

	if err = aocSessionInject(req); err != nil {
		return
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%w; response code %d", ErrUnexpectedResponseCode, resp.StatusCode)
		return
	}

	byteArr, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response := string(byteArr)
	if strings.Contains(response, "You don't seem to be solving the right level.") || strings.Contains(response, "That's the right answer!") {
		return
	} else if strings.Contains(response, ResponseIncorrect) {
		indexIn := strings.Index(response, ResponseIncorrect)
		indexOut := strings.Index(response[indexIn:], "<")
		err = errors.New(response[indexIn : indexIn+indexOut])
		ok = false
		return
	} else if strings.Contains(response, ResponseCooldown) {
		indexIn := strings.Index(response, ResponseCooldown)
		indexOut := strings.Index(response[indexIn:], "<")
		err = errors.New(response[indexIn : indexIn+indexOut])
		ok = false
		return
	} else {
		err = ErrUnexpectedResponsePage
		return
	}
}

func aocSessionInject(req *http.Request) error {
	sessionToken := os.Getenv("AOC_SESSION_COOKIE")
	if sessionToken == "" {
		return ErrNoSessionCookie
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
	return nil
}
