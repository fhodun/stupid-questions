package main

import (
	"github.com/agnivade/levenshtein"
	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/stupid-questions/config"
	log "github.com/sirupsen/logrus"
	//"strconv"
	"strings"
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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cancerWords := []string{ "opera", "testportal" }
	content := m.Content
	split := strings.Split(content, " ")
	// TODO: add punctuation characters deletion e.g. '.' and ','
	var words []string
	println("Content words number: ", len(split))
	for i := 0; i < len(split); i++ {
		println("\nTESTED WORD: ", split[i])
		for j := 0; j < len(cancerWords); j++ {
			println("Searched word: ", cancerWords[j])
			distance := levenshtein.ComputeDistance(split[i], cancerWords[j])
			println("Difference: ", distance)
			if distance <= cfg.DistanceMax {
				words = append(words, split[i])
			}
		}
	}
	//println(words)
}
