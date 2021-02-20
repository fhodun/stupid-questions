package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config dupa
type Config struct {
	Discord struct {
		Token  string
		Prefix string
	}
	DistanceMax int
}

// GetConfig dupa
func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Unsuccessful loading .env, ", err)
	}

	token, tokenExists := os.LookupEnv("DISCORD_TOKEN")
	prefix, prefixExists := os.LookupEnv("DISCORD_PREFIX")
	distanceMax2, distanceMaxExists := os.LookupEnv("DISTANCE_MAX")
	distanceMax, err := strconv.Atoi(distanceMax2)
	if err != nil {
		log.Warn("Unsuccessful string converting")
		distanceMax = 2
	}
	if !tokenExists {
		log.Fatal("No discord token detected")
	}
	if !prefixExists {
		log.Warn("No discord prefix detected, default '>' will be used")
		prefix = ">"
	}
	if !distanceMaxExists {
		log.Warn("No max distance detected, default 2 will be used")
		distanceMax = 2
	}

	config := &Config{
		Discord: struct {
			Token  string
			Prefix string
		}{
			Token:  token,
			Prefix: prefix,
		},
		DistanceMax: distanceMax,
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
