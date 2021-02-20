package main

import (
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/stupid-questions/config"
	log "github.com/sirupsen/logrus"
)

var cfg = config.GetConfig()

func main() {
	config.InitConfig()

	dg, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		log.Fatal("Discord session creation failed")
	}
	err = dg.Open()
	if err != nil {
		log.Fatal("Unsuccessful opening connection")
	}
	defer dg.Close()

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	log.Info("Bot is now running")
	select {}
}

type CancerWordTag struct {
	Weight uint
	Text   string
}

// CancerWord is used for defining cancer words
type CancerWord struct {
	PrimaryText string
	Tags        []CancerWordTag
	Answer      string
}

const MinWeight uint = 5
const MaxDistance int = 2

func SmartStringContains(str string, substr string) bool {
	for i := 0; i < len(str)-len(substr); i++ {
		if levenshtein.ComputeDistance(substr, str[i:i+len(substr)]) < MaxDistance {
			return true
		}
	}

	return false
}

func GetMessageCancer(str string, cancerWords []CancerWord) *CancerWord {
	words := strings.Split(str, " ")

	for _, cw := range cancerWords {
		if !SmartStringContains(str, cw.PrimaryText) {
			continue
		}

		w := uint(0)
		hasAnyTag := false

		for _, word := range words {
			if word == cw.PrimaryText {
				continue
			}
			for _, tag := range cw.Tags {
				fmt.Printf("distance: %d from %s to %s\n", levenshtein.ComputeDistance(word, tag.Text), word, tag.Text)
				if distance := levenshtein.ComputeDistance(word, tag.Text); distance < MaxDistance {
					hasAnyTag = true
					w += tag.Weight
					if w > MinWeight {
						return &cw
					}
				}
			}
		}
		if !hasAnyTag {
			return nil
		}
		fmt.Printf("w: %d\n", w)
	}
	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cancerWords := []CancerWord{
		{
			PrimaryText: "anti-testportal",
			Answer:      "tak kurwa dziala spierdalaj",
			Tags: []CancerWordTag{
				{
					Weight: 10,
					Text:   "dziala",
				},
			},
		},
	}
	if cw := GetMessageCancer(m.Content, cancerWords); cw != nil {
		s.ChannelMessageSendReply(m.ChannelID, cw.Answer, m.Reference())
	}
}
