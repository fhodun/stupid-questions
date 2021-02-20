package main

import (
	"github.com/fhodun/stupid-questions/bot"
	"github.com/fhodun/stupid-questions/config"
	"github.com/fhodun/stupid-questions/qp"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Load()
	qp := qp.QuestionParser{
		Sentences:   cfg.Sentences,
		MinWeight:   5,
		MaxDistance: 2,
	}

	bot, err := bot.New(cfg, qp)
	if err != nil {
		panic(err)
	}
	defer bot.Close()
	if err := bot.Open(); err != nil {
		panic(err)
	}
	logrus.Info("Successfully started connection with discord")
	select {}
}
