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
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := m.Content

	if (strings.Contains(content, "anti")&&strings.Contains(content, "testportal")&&strings.Contains(content, "dziala")&&strings.Contains(content, "?")) {
		s.ChannelMessageSendReply(m.ChannelID, "zamknij wreszcie p***e", m.Reference())
	}
	if (strings.Contains(content, "używam")&&strings.Contains(content, "opera gx")||strings.Contains(content, "opery gx")) {
		s.ChannelMessageSendReply(m.ChannelID, "chyba z drzewa spadłeś", m.Reference())
		//s.GuildBanCreateWithReason(m.GuildID, m.Author.ID, "Używanie gównianych przeglądarek", 0)
	}

	/*switch cmd {
	case "anti testportal dziala?":
		s.ChannelMessageSend(m.ChannelID, "dupa")
	case "testshoter dziala?":
		s.ChannelMessageSend(m.ChannelID, "dupa")
	case "bot testportal dziala?":
		s.ChannelMessageSend(m.ChannelID, "dupa")
	case "opera gx":
		s.GuildBan(m.GuildID, m.Author.ID)
	}*/
}
