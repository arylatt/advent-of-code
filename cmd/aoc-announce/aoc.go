package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/arylatt/advent-of-code/aoc"
)

var (
	aocClient    AOCClient
	tokenFlagged = false
	leaderboard  aoc.Leaderboard
	aocClientMu  = sync.RWMutex{}
)

type AOCClient interface {
	Leaderboard(event, owner string) (aoc.Leaderboard, error)
}

func loop(ctx context.Context, loopTime time.Duration, event, owner string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("[loop] context done.")
			return
		case <-time.After(loopTime):
			getLeaderboard(event, owner)
		}
	}
}

func getLeaderboard(event, owner string) {
	log.Println("[getLeaderboard] fetching leaderboard")

	aocClientMu.RLock()

	lb, err := aocClient.Leaderboard(event, owner)
	if errors.Is(err, aoc.ErrRedirectBlocked) && !tokenFlagged {
		aocClientMu.RUnlock()
		flagToken(err)
		return
	}

	aocClientMu.RUnlock()

	if err != nil {
		log.Printf("[getLeaderboard] error fetching leaderboard: %s\n", err)
		return
	}

	tokenFlagged = false

	for memberID, member := range lb.Members {
		if strings.TrimSpace(member.Name) == "" {
			log.Printf("[getLeaderboard] skipping member %d with no name\n", memberID)
			continue
		}

		previousEntry, ok := leaderboard.Members[memberID]
		if !ok {
			processNewMember(member)
		} else {
			processMember(member, previousEntry)
		}

		log.Printf("[getLeaderboard] processed member %d\n", memberID)
	}

	leaderboard = lb
}

func processNewMember(member aoc.Member) {
	discordClientMu.RLock()

	err := discordClient.SendMessage(fmt.Sprintf(":star: %s has joined the leaderboard!", member.Name))

	discordClientMu.RUnlock()

	if err != nil {
		log.Printf("[processNewMember] error announcing new member %d: %s\n", member.ID, err)
	} else {
		log.Printf("[processNewMember] member %d joins the party\n", member.ID)
	}
}

func generateStarText(day int, part int, starTime time.Time) string {
	return fmt.Sprintf("day %d part %d at %s", day, part, starTime.Format("2006-01-02 15:04:05"))
}

func generateStarMessage(name string, starTexts []string) string {
	msg := ""
	for i := 0; i < len(starTexts); i++ {
		msg += ":star2:"
	}

	msg += fmt.Sprintf(" %s completed ", name)

	if len(starTexts) >= 2 {
		lastStar := len(starTexts) - 1
		starTexts[lastStar] = "and " + starTexts[lastStar]
	}

	separator := ", "
	if len(starTexts) < 2 {
		separator = " "
	}

	msg += strings.Join(starTexts, separator)
	return msg + "."
}

func processMember(current, previous aoc.Member) {
	if current.LastStarTimestamp == previous.LastStarTimestamp {
		log.Printf("[processMember] member %d is unchanged\n", current.ID)
		return
	}

	newStars := []string{}

	for day, parts := range current.CompletionDayLevel {
		if parts.Part1.GetStarTimestamp > previous.LastStarTimestamp {
			newStars = append(newStars, generateStarText(day, 1, parts.Part1.GetStarTime()))
		}
		if parts.Part2.GetStarTimestamp > previous.LastStarTimestamp {
			newStars = append(newStars, generateStarText(day, 2, parts.Part2.GetStarTime()))
		}
	}

	if len(newStars) == 0 {
		log.Printf("[processMember] member %d is unchanged (but their timestamp changed???)\n", current.ID)
		return
	}

	discordClientMu.RLock()

	err := discordClient.SendMessage(generateStarMessage(current.Name, newStars))

	discordClientMu.RUnlock()

	if err != nil {
		log.Printf("[processMember] error announcing member update %d: %s\n", current.ID, err)
	} else {
		log.Printf("[processMember] member %d did some things\n", current.ID)
	}
}

func flagToken(err error) {
	tokenFlagged = true

	discordClientMu.RLock()

	sendErr := discordClient.SendMessage(fmt.Sprintf(":x: AOC token appears to have expired: %s", err))

	discordClientMu.RUnlock()

	if sendErr != nil {
		log.Printf("[flagToken] error flagging expired token: %s\n", sendErr)
	} else {
		log.Printf("[flagToken] token flagged\n")
	}
}
