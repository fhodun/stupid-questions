package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/stupid-questions/config"
	"github.com/fhodun/stupid-questions/qp"
	"github.com/fhodun/stupid-questions/utils"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	cfg     config.Config
	qp      qp.QuestionParser
	session *discordgo.Session
}

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

func (bot Bot) Open() error {
	err := bot.session.Open()
	if err != nil {
		return fmt.Errorf("fail opening discord websockets session, err: %s", err.Error())
	}
	bot.session.AddHandler(bot.onMessageCreate)
	bot.session.Identify.Intents = discordgo.IntentsGuildMessages

	return nil
}

func (bot Bot) Close() {
	bot.session.Close()
}

func (bot Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Printf("Received msg:%s\n ", m.Content)

	pureString, err := utils.RemovePolishCharacters(m.Content)
	if err != nil {
		logrus.Errorf("Fail removing polish shit %s\n", err.Error())
		return
	}

	if cw := bot.qp.ParseString(pureString); cw != nil {
		s.ChannelMessageSendReply(m.ChannelID, cw.Answer, m.Reference())
	}
}
