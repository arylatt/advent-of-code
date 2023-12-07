package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscordClientSendMessage(t *testing.T) {
	msg := "hello, world"
	expectedBody, _ := newMessage(msg)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualBody := discordMessage{}

		if assert.NoError(t, json.NewDecoder(r.Body).Decode(&actualBody)) {
			assert.Equal(t, expectedBody, actualBody)
		}

		assert.Equal(t, discordUserAgent, r.Header.Get("User-Agent"))

		assert.Equal(t, discordPayloadContentType, r.Header.Get("Content-Type"))

		w.WriteHeader(http.StatusNoContent)
	}))

	defer srv.Close()

	discordClient = &DiscordClient{httpClient: srv.Client(), webhookURL: srv.URL}

	assert.NoError(t, discordClient.SendMessage(msg))
}

func TestNewMessage(t *testing.T) {
	msg := "hello, world"

	expected := discordMessage{AvatarURL: discordAvatarURL, Content: msg, Username: discordUsername}

	actual, err := newMessage(msg)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestNewMessage_Errors(t *testing.T) {
	actual, err := newMessage("  ")

	if assert.Error(t, err) {
		assert.Empty(t, actual)
	}
}
