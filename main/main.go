package main

import (
	"os"
	"time"

	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	_ "github.com/AkinoKaede/question-mark-reply-bot/main/distro/all"

	tele "gopkg.in/telebot.v3"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  os.Getenv("QMRBOT_TELEGRAM_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	features.Handle(b)

	b.Start()
}
