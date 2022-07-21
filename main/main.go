package main

import (
	"os"
	"time"

	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	_ "github.com/AkinoKaede/question-mark-reply-bot/main/distro/all"

	tb "gopkg.in/telebot.v3"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("QMRBOT_TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	features.Handle(b)

	b.Start()
}
