package aoc

import "time"

// Leaderboard represents a private AoC leaderboard.
type Leaderboard struct {
	// Members is a map of member IDs to member objects.
	Members map[int]Member `json:"members"`
	// OwnerID is the owner (and leaderboard) ID.
	OwnerID int `json:"owner_id"`
	// Event is the event the leaderboard is showing data for.
	Event string `json:"event"`
}

// Member represents a member on the leaderboard.
type Member struct {
	// Name is the members name.
	// Will be an empty string if the member is appearing anonymously.
	Name string `json:"name"`
	// LastStarTimestamp indicates when the member received their most
	// recent star.
	LastStarTimestamp int64 `json:"last_star_ts"`
	// Stars is the number of total stars the member has achieved.
	Stars int `json:"stars"`
	// ID is the member ID (also the key of the map entry).
	ID int `json:"id"`
	// GlobalScore is the members global leaderboard score.
	GlobalScore int `json:"global_score"`
	// LocalScore is the members score for this private leaderboard.
	LocalScore int `json:"local_score"`
	// CompletionDayLevel represents which days and parts the member has completed.
	CompletionDayLevel map[int]DayCompletion `json:"completion_day_level"`
}

// LastStarTime returns the conversion of LastStarTimestamp to time.Time.
func (m Member) LastStarTime() time.Time {
	return timestamp(m.LastStarTimestamp)
}

// DayCompletion represents a given days completion state.
type DayCompletion struct {
	// Part1 is the members completion state of the current day part 1.
	Part1 PartCompletion `json:"1"`
	// Part2 is the members completion state of the current day part 2.
	Part2 PartCompletion `json:"2"`
}

// PartCompletion represents the completion state. Values will be empty if
// the part is not completed.
type PartCompletion struct {
	// GetStarTimestamp represents when the member completed the part.
	GetStarTimestamp int64 `json:"get_star_ts"`
	// StarIndex is the unique index of the star.
	StarIndex int `json:"star_index"`
}

// Completed returns if the part is completed.
func (p PartCompletion) Completed() bool {
	return p.GetStarTimestamp != 0 && p.StarIndex != 0
}

// GetStarTime returns the conversion of GetStarTimestamp to time.Time.
func (p PartCompletion) GetStarTime() time.Time {
	return timestamp(p.GetStarTimestamp)
}

func timestamp(unix int64) time.Time {
	return time.Unix(unix, 0)
}
