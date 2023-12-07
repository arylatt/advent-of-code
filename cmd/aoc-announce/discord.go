package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	discordUsername           = "Advent of Code Announcer"
	discordAvatarURL          = "https://adventofcode.com/favicon.png"
	discordUserAgent          = "aoc-announce/1.0 (https://github.com/arylatt/advent-of-code)"
	discordPayloadContentType = "application/json"
)

type (
	Discord interface {
		SendMessage(message string) error
	}

	DiscordClient struct {
		httpClient *http.Client
		webhookURL string
	}

	discordMessage struct {
		AvatarURL string `json:"avatar_url"`
		Content   string `json:"content"`
		Username  string `json:"username"`
	}
)

var (
	discordClient Discord
)

func (d *DiscordClient) SendMessage(message string) error {
	msg, err := newMessage(message)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}

	err = json.NewEncoder(buf).Encode(&msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, d.webhookURL, buf)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", discordUserAgent)
	req.Header.Set("Content-Type", discordPayloadContentType)

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected http response code from discord: %q", resp.Status)
	}

	return nil
}

func newMessage(message string) (discordMessage, error) {
	if strings.TrimSpace(message) == "" {
		return discordMessage{}, errors.New("cannot construct discord message, message is empty")
	}

	return discordMessage{AvatarURL: discordAvatarURL, Content: message, Username: discordUsername}, nil
}

// https://canary.discord.com/api/webhooks/944287669385588766/1nOkzcex8P_oBXRBeUcsFZTJ6X-Kz-4rI_gW0POx_Vk1SOjuLD1WxWTdByo7H563A8Oy
