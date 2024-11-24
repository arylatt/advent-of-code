package main

import (
	"context"
	"io"
	"log"
	"testing"

	"github.com/arylatt/advent-of-code/aoc"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	MockAOC struct {
		mock.Mock
	}

	MockDiscord struct {
		mock.Mock
	}
)

func (m *MockAOC) Leaderboard(event, owner string) (aoc.Leaderboard, error) {
	args := m.Called(event, owner)

	lb, ok := args.Get(0).(aoc.Leaderboard)
	if !ok {
		lb = aoc.Leaderboard{}
	}

	return lb, args.Error(1)
}

func (m *MockDiscord) SendMessage(msg string) error {
	args := m.Called(msg)

	return args.Error(0)
}

func TestMain(m *testing.M) {
	log.SetOutput(io.MultiWriter())

	m.Run()
}

func TestInitAOCClient(t *testing.T) {
	viper.Reset()
	viper.Set("token", "test")

	c, err := initAOCClient(context.Background())

	if assert.NoError(t, err) {
		assert.IsType(t, &aoc.Client{}, c)
	}
}

func TestInitDiscordClient(t *testing.T) {
	viper.Reset()
	viper.Set("discord-webhook-url", "test")

	c, err := initDiscordClient(context.Background())

	if assert.NoError(t, err) {
		assert.IsType(t, &DiscordClient{}, c)
	}
}
