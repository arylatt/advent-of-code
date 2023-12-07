package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/arylatt/advent-of-code/aoc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:  "aoc-announce [-e EVENT] [-l LEADERBOARD] [-t AOC_TOKEN] [-d DISCORD_WEBHOOK_URL]",
	RunE: runE,
}

func init() {
	rootCmd.PersistentFlags().StringP("event", "e", "", "AoC event to monitor.")
	rootCmd.PersistentFlags().StringP("leaderboard", "l", "", "AoC leaderboard to monitor.")
	rootCmd.PersistentFlags().StringP("token", "t", "", "AoC session token to use.")
	rootCmd.PersistentFlags().StringP("discord-webhook-url", "d", "", "Discord webhook URL for posting events to.")

	viper.BindPFlags(rootCmd.PersistentFlags())

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("aoc_announce")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func initAOCClient() (*aoc.Client, error) {
	c, err := aoc.New(aoc.WithSessionToken(viper.GetString("token")))
	if err != nil {
		return nil, err
	}

	aocClient = c
	return c, nil
}

func initDiscordClient() (*DiscordClient, error) {
	webhookURL := viper.GetString("discord-webhook-url")
	if strings.TrimSpace(webhookURL) == "" {
		return nil, errors.New("discord webhook url is missing")
	}

	return &DiscordClient{httpClient: http.DefaultClient, webhookURL: webhookURL}, nil
}

func runE(cmd *cobra.Command, args []string) error {
	var err error

	aocClient, err = initAOCClient()
	if err != nil {
		return err
	}

	discordClient, err = initDiscordClient()
	if err != nil {
		return err
	}

	leaderboard = aoc.Leaderboard{}

	event, owner := viper.GetString("event"), viper.GetString("leaderboard")
	if strings.TrimSpace(event) == "" || strings.TrimSpace(owner) == "" {
		return fmt.Errorf("event and leaderboard must be non-empty. event=%q, leaderboard=%q", event, owner)
	}

	for leaderboard.OwnerID == 0 && leaderboard.Event == "" {
		leaderboard, err = aocClient.Leaderboard(event, owner)
		if err != nil {
			log.Printf("Failed to populate initial leaderboard: %q, retrying in 15 seconds\n", err.Error())

			<-time.After(time.Second * 15)
		}
	}

	log.Println("Populated initial leaderboard data!")

	ctx, cancel := context.WithCancel(cmd.Context())

	go loop(ctx, aoc.DefaultCacheTime, event, owner)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	log.Println("aoc-announce is running!")

	<-c

	log.Println("Caught signal")

	cancel()

	<-ctx.Done()

	return nil
}
