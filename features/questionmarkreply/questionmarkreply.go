package questionmarkreply

import (
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	tele "gopkg.in/telebot.v3"
)

var (
	QuestionMarks = []string{"?", "¿", "？"}
)

func OnText(c tele.Context) error {
	for _, mark := range QuestionMarks {
		if c.Message().Text == mark {
			c.Reply(mark)
			return nil
		}
	}

	return nil
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
}
