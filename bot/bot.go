package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/stupid-questions/config"
	"github.com/fhodun/stupid-questions/qp"
	"github.com/sirupsen/logrus"
)

// Bot dupa
type Bot struct {
	cfg     config.Config
	qp      qp.QuestionParser
	session *discordgo.Session
}

// New dupa
func New(cfg config.Config, qp qp.QuestionParser) (Bot, error) {
	bot := Bot{
		cfg: cfg,
		qp:  qp,
	}

	var err error
	bot.session, err = discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return bot, fmt.Errorf("fail creating discord seesion, err: %s", err.Error())
	}
	return bot, nil
}

// Open dupa
func (bot Bot) Open() error {
	err := bot.session.Open()
	if err != nil {
		return fmt.Errorf("fail opening discord websockets session, err: %s", err.Error())
	}
	bot.session.AddHandler(bot.onMessageCreate)
	bot.session.Identify.Intents = discordgo.IntentsGuildMessages

	return nil
}

// Close dupa
func (bot Bot) Close() {
	bot.session.Close()
}

func (bot Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	logrus.WithFields(logrus.Fields{
		"msg": m.Content,
		"ID":  m.ID,
	}).Infoln("Received message")

	if cw := bot.qp.ParseString(m.Content); cw != nil {
		logrus.WithFields(logrus.Fields{
			"msg":    m.Content,
			"answer": cw.Answer,
			"ID":     m.ID,
		}).Infoln("Sending answer")
		s.ChannelMessageSendReply(m.ChannelID, cw.Answer, m.Reference())
	}
}
