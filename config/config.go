package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

// Config dupa
type Config struct {
	Discord struct {
		Token  string
		Prefix string
	}
}

// GetConfig dupa
func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Unsuccessful loading .env, ", err)
	}

	token, discordTokenExists := os.LookupEnv("DISCORD_TOKEN")
	prefix, prefixExists := os.LookupEnv("DISCORD_PREFIX")
	if !tokenExists {
		log.Fatal("No discord token detected")
	}
	if !prefixExists {
		log.Warn("No discord prefix detected, default '>' will be used")
		prefix = ">"
	}

	config := &Config{
		Discord: struct {
			Token  string
			Prefix string
		}{
			Token:  token,
			Prefix: prefix,
		}
	}
	return config
}

// InitConfig dupa
func InitConfig() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	log.SetOutput(os.Stdout)
}
