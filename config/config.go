package config

import (
	"os"

	"github.com/fhodun/stupid-questions/qp"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config dupa
type Config struct {
	DiscordToken string
	Sentences    []qp.Sentence
}

// mustGetEnv dupa
func mustGetEnv(key string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		logrus.Fatalf("'%s' env not present in .env", key)
	}
	return env
}

// Load dupa
func Load() Config {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Fail loading .env file", err)

	}
	// TODO: figure out some better way to load those, maybe JSON file since they're too complex for .env files
	sentences := []qp.Sentence{
		{
			PrimaryWord: "dziala",
			Answer:      "tak kurwa dziala",
			Tags: []qp.SentenceTag{
				{
					Weight: 10,
					Text:   "testportal",
				},
				{
					Weight: 10,
					Text:   "anti testportal",
				},
				{
					Weight: 10,
					Text:   "anty testportal",
				},
				{
					Weight: 10,
					Text:   "anti-testportal",
				},
				{
					Weight: 10,
					Text:   "wtyczka",
				},
				{
					Weight: 10,
					Text:   "nadal",
				},
				{
					Weight: 10,
					Text:   "jeszcze",
				},
			},
		},
	}

	return Config{
		DiscordToken: mustGetEnv("DISCORD_TOKEN"),
		Sentences:    sentences,
	}
}
