package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/arylatt/advent-of-code/aoc"
	"github.com/fsnotify/fsnotify"
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
	rootCmd.PersistentFlags().String("token-file", "", "File to watch containing AoC session token to use.")
	rootCmd.PersistentFlags().StringP("discord-webhook-url", "d", "", "Discord webhook URL for posting events to.")
	rootCmd.PersistentFlags().String("discord-webhook-url-file", "", "File to watch containing Discord webhook URL for posting events to.")

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

func initAOCClient(ctx context.Context) (*aoc.Client, error) {
	token, err := getAOCToken(ctx)
	if err != nil {
		return nil, err
	}

	c, err := aoc.New(aoc.WithSessionToken(token))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func initDiscordClient(ctx context.Context) (*DiscordClient, error) {
	webhookURL, err := getDiscordWebhookURL(ctx)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(webhookURL) == "" {
		return nil, errors.New("discord webhook url is missing")
	}

	return &DiscordClient{httpClient: http.DefaultClient, webhookURL: webhookURL}, nil
}

func runE(cmd *cobra.Command, args []string) error {
	var err error

	ctx, cancel := context.WithCancel(cmd.Context())
	defer cancel()

	aocClient, err = initAOCClient(ctx)
	if err != nil {
		return err
	}

	discordClient, err = initDiscordClient(ctx)
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

func getAOCToken(ctx context.Context) (string, error) {
	return getSecret(ctx, "token", func(token string) error {
		newAOCClient, err := aoc.New(aoc.WithSessionToken(token))
		if err != nil {
			return fmt.Errorf("failed to create new aoc client with updated token: %w", err)
		}

		aocClientMu.Lock()
		defer aocClientMu.Unlock()

		aocClient = newAOCClient
		return nil
	})
}

func getDiscordWebhookURL(ctx context.Context) (string, error) {
	return getSecret(ctx, "discord-webhook-url", func(token string) error {
		newDiscordClient := &DiscordClient{discordClient.(*DiscordClient).httpClient, token}

		discordClientMu.Lock()
		defer discordClientMu.Unlock()

		discordClient = newDiscordClient
		return nil
	})
}

func getSecret(ctx context.Context, secretName string, updateFunc func(string) error) (string, error) {
	secret, watcher, err := setupSecretWatcher(secretName)
	if err != nil {
		return "", err
	}

	if watcher == nil {
		return secret, nil
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Name != viper.GetString(secretName+"-file") || !event.Has(fsnotify.Write) {
					continue
				}

				fileBytes, err := os.ReadFile(event.Name)
				if err != nil {
					log.Printf("failed to read %s file: %v\n", secretName, err)
					continue
				}

				newSecret := string(fileBytes)
				if err := updateFunc(newSecret); err != nil {
					log.Printf("failed to update %s: %v\n", secretName, err)
				}
			case err := <-watcher.Errors:
				log.Printf("%s watcher error: %v\n", secretName, err)
			case <-ctx.Done():
				watcher.Close()
				return
			}
		}
	}()

	return secret, nil
}

// setupSecretWatcher sets up a file watcher for a secret file and returns the secret value.
// If the secret is provided directly via a configuration value, it returns the secret without setting up a watcher.
// The function returns the secret value, a file watcher (if applicable), and an error if any.
func setupSecretWatcher(secretName string) (string, *fsnotify.Watcher, error) {
	if secret := viper.GetString(secretName); strings.TrimSpace(secret) != "" {
		return secret, nil, nil
	}

	secretFile := viper.GetString(secretName + "-file")
	if strings.TrimSpace(secretFile) == "" {
		return "", nil, fmt.Errorf("no %s file provided", secretName)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return "", nil, fmt.Errorf("failed to create fsnotify watcher: %w", err)
	}

	dirPath, err := filepath.Abs(path.Dir(secretFile))
	if err != nil {
		return "", nil, fmt.Errorf("failed to get absolute path for %s file: %w", secretName, err)
	}

	err = watcher.Add(dirPath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to watch %s file: %w", secretName, err)
	}

	fileBytes, err := os.ReadFile(secretFile)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read %s file: %w", secretName, err)
	}

	return strings.TrimSpace(string(fileBytes)), watcher, nil
}
