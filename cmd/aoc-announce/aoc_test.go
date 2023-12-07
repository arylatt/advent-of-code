package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/arylatt/advent-of-code/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoop(t *testing.T) {
	mockAOC := &MockAOC{}
	aocClient = mockAOC

	call := mockAOC.On("Leaderboard", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(aoc.Leaderboard{}, aoc.ErrSessionTokenMustBeProvided).Twice()

	ctx, cancel := context.WithCancel(context.TODO())
	event, owner := "1", "2"

	call.Run(func(args mock.Arguments) {
		assert.Equal(t, event, args.String(0))
		assert.Equal(t, owner, args.String(1))

		if call.Repeatability == -1 {
			cancel()
		}
	})

	loop(ctx, time.Millisecond*50, event, owner)
}

func TestGetLeaderboard(t *testing.T) {
	mockAOC := &MockAOC{}
	mockDiscord := &MockDiscord{}
	aocClient = mockAOC
	discordClient = mockDiscord

	leaderboard = aoc.Leaderboard{
		Members: map[int]aoc.Member{
			1: {
				ID:   1,
				Name: "n1",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
					},
				},
				LastStarTimestamp: 1,
			},
		},
	}

	lb := aoc.Leaderboard{
		Members: map[int]aoc.Member{
			1: {
				ID:   1,
				Name: "n1",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
					},
					2: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 2},
					},
				},
				LastStarTimestamp: 2,
			},
			2: {
				ID:   2,
				Name: "n2",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
					},
				},
				LastStarTimestamp: 1,
			},
			3: {
				ID:   3,
				Name: "",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
					},
				},
				LastStarTimestamp: 1,
			},
		},
	}

	event, owner := "1", "2"

	mockDiscord.On("SendMessage", mock.AnythingOfType("string")).Return(nil).Once().Twice()
	mockAOC.On("Leaderboard", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(lb, nil).Once().Run(func(args mock.Arguments) {
		assert.Equal(t, event, args.String(0))
		assert.Equal(t, owner, args.String(1))
	})

	tokenFlagged = true

	getLeaderboard(event, owner)

	assert.False(t, tokenFlagged)
	assert.Equal(t, lb, leaderboard)

	mockAOC.AssertExpectations(t)
	mockDiscord.AssertExpectations(t)
}

func TestGetLeaderboard_FlagsHTTPError(t *testing.T) {
	mockAOC := &MockAOC{}
	mockDiscord := &MockDiscord{}
	aocClient = mockAOC
	discordClient = mockDiscord

	mockAOC.On("Leaderboard", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(aoc.Leaderboard{}, fmt.Errorf("oh no, %w", aoc.ErrRedirectBlocked)).Twice()
	mockDiscord.On("SendMessage", mock.AnythingOfType("string")).Return(nil).Once()

	tokenFlagged = false

	for i := 0; i < 2; i++ {
		getLeaderboard("1", "2")

		assert.True(t, tokenFlagged)
	}

	mockAOC.AssertExpectations(t)
	mockDiscord.AssertExpectations(t)
}

func TestProcessNewMember(t *testing.T) {
	tests := []struct {
		member   aoc.Member
		expected string
	}{
		{aoc.Member{Name: "n1"}, ":star: n1 has joined the leaderboard!"},
		{aoc.Member{Name: "n2"}, ":star: n2 has joined the leaderboard!"},
	}

	mockDiscord := &MockDiscord{}
	discordClient = mockDiscord

	call := mockDiscord.On("SendMessage", mock.AnythingOfType("string")).Return(nil)
	calls := 0

	for _, test := range tests {
		call.Run(func(args mock.Arguments) {
			assert.Equal(t, test.expected, args.String(0))
		})

		calls++

		processNewMember(test.member)

		mockDiscord.AssertNumberOfCalls(t, "SendMessage", calls)
	}
}

func TestGenerateStarText(t *testing.T) {
	tests := []struct {
		day, part int
		starTime  time.Time
		expected  string
	}{
		{1, 1, time.Date(2023, 12, 4, 12, 23, 1, 0, time.UTC), "day 1 part 1 at 2023-12-04 12:23:01"},
		{2, 1, time.Date(2023, 12, 4, 18, 23, 1, 0, time.UTC), "day 2 part 1 at 2023-12-04 18:23:01"},
		{2, 2, time.Date(2023, 12, 4, 18, 23, 1, 0, time.UTC), "day 2 part 2 at 2023-12-04 18:23:01"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, generateStarText(test.day, test.part, test.starTime))
	}
}

func TestGenerateStarMessage(t *testing.T) {
	tests := []struct {
		name     string
		stars    []string
		expected string
	}{
		{"n1", []string{"star1"}, ":star2: n1 completed star1."},
		{"n2", []string{"star1", "star2"}, ":star2::star2: n2 completed star1, and star2."},
		{"n2", []string{"star1", "star2", "star3"}, ":star2::star2::star2: n2 completed star1, star2, and star3."},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, generateStarMessage(test.name, test.stars))
	}
}

func TestProcessMember(t *testing.T) {
	tests := []struct {
		previous, current aoc.Member
		expected          string
	}{
		{
			aoc.Member{
				Name: "n1",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
					},
				},
				LastStarTimestamp: 1,
			},
			aoc.Member{
				Name: "n1",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
						Part2: aoc.PartCompletion{GetStarTimestamp: 2},
					},
				},
				LastStarTimestamp: 2,
			},
			generateStarMessage("n1", []string{generateStarText(1, 2, time.Unix(2, 0))}),
		},
		{
			aoc.Member{
				Name: "n2",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
						Part2: aoc.PartCompletion{GetStarTimestamp: 2},
					},
				},
				LastStarTimestamp: 2,
			},
			aoc.Member{
				Name: "n2",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
						Part2: aoc.PartCompletion{GetStarTimestamp: 2},
					},
				},
				LastStarTimestamp: 2,
			},
			"",
		},
		{
			aoc.Member{
				Name: "n3",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
					},
				},
				LastStarTimestamp: 1,
			},
			aoc.Member{
				Name: "n3",
				CompletionDayLevel: map[int]aoc.DayCompletion{
					1: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 1},
						Part2: aoc.PartCompletion{GetStarTimestamp: 2},
					},
					2: {
						Part1: aoc.PartCompletion{GetStarTimestamp: 2},
					},
				},
				LastStarTimestamp: 2,
			},
			generateStarMessage("n3", []string{generateStarText(1, 2, time.Unix(2, 0)), generateStarText(2, 1, time.Unix(2, 0))}),
		},
	}

	mockDiscord := &MockDiscord{}
	discordClient = mockDiscord

	call := mockDiscord.On("SendMessage", mock.AnythingOfType("string")).Return(nil)
	calls := 0

	for _, test := range tests {
		call.Run(func(args mock.Arguments) {
			assert.Equal(t, test.expected, args.String(0))
		})

		if test.expected != "" {
			calls++
		}

		processMember(test.current, test.previous)

		mockDiscord.AssertNumberOfCalls(t, "SendMessage", calls)
	}
}

func TestFlagToken(t *testing.T) {
	err := "test error"
	msg := fmt.Sprintf(":x: AOC token appears to have expired: %s", err)

	mockDiscord := &MockDiscord{}
	discordClient = mockDiscord

	mockDiscord.On("SendMessage", mock.AnythingOfType("string")).Run(func(args mock.Arguments) {
		assert.Equal(t, msg, args.String(0))
	}).Return(nil)

	tokenFlagged = false

	flagToken(errors.New(err))

	assert.True(t, tokenFlagged)
}
