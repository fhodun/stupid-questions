package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/stupid-questions/config"
	log "github.com/sirupsen/logrus"
	"strings"
)

var cfg = config.GetConfig()

func main() {
	config.InitConfig()
	cfg := config.GetConfig()

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
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, cfg.Discord.Prefix) {
		return
	}

	// TODO: get prefix value another, more efficient way
	args := strings.Split(m.Content[(len(cfg.Discord.Prefix)):], ">")

	if len(args) < 3 {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, "Too few arguments")
		return
	}

	cmd := strings.Trim(args[0], " ")
	arg := args[1] // first argument
	msg := args[2] // second argument

	if cmd == "info" {
		s.ChannelMessageSend(m.ChannelID, "dupa")
	}

	// golang error bypass, need to remove
	if arg == msg {
	}
}
