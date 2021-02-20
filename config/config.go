package config

import (
	"os"

	"github.com/fhodun/stupid-questions/qp"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config dupa
type Config struct {
	DiscordToken string
	Sentences    []qp.Sentence
}

func mustGetEnv(key string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("'%s' env not present in .env", key)
	}
	return env
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Fail loading .env file", err)
	}

	// Figure out some better way to load those, maybe JSON file since they're too complex for .env files
	sentences := []qp.Sentence{
		{
			PrimaryWord: "anti-testportal",
			Answer:      "tak kurwa dziala spierdalaj",
			Tags: []qp.SentenceTag{
				{
					Weight: 10,
					Text:   "dziala",
				},
			},
		},
	}

	return Config{
		DiscordToken: mustGetEnv("DISCORD_TOKEN"),
		Sentences:    sentences,
	}
}
