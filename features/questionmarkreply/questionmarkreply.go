package questionmarkreply

import (
	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	tele "gopkg.in/telebot.v3"
)

var (
	QuestionMarks = []string{"?", "¿", "？", "❓", "❔"}
)

func OnText(c tele.Context) error {
	text := c.Message().Text

	for _, b := range text {
		if !common.Contains(string(b), QuestionMarks) {
			return nil
		}
	}

	return c.Reply(text)

}

func OnSticker(c tele.Context) error {
	sticker := c.Message().Sticker

	if common.Contains(sticker.Emoji, []string{"❓", "❔"}) {
		return c.Reply(sticker)
	}

	return nil
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
	features.RegisterFeature(tele.OnSticker, OnSticker)
}
