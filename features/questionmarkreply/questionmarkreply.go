package questionmarkreply

import (
	"strings"

	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	tele "gopkg.in/telebot.v3"
)

var (
	QuestionMarks      = append([]string{"?", "¿", "？"}, QuestionMarkEmojis...)
	QuestionMarkEmojis = []string{"❓", "❔"}
)

func OnText(c tele.Context) error {
	text := c.Message().Text

	for _, b := range text {
		if !common.Contains(string(b), QuestionMarks) {
			return nil
		}
	}

	if c.Message().ReplyTo != nil && len(text) > 1 {
		markCount := make(map[rune]int)
		for _, b := range text {
			markCount[b]++
		}

		for k, v := range markCount {
			if v == len(text) {
				replyToText := c.Message().ReplyTo.Text
				if strings.Count(replyToText, string(k)) == len(replyToText) {
					replyMsg := &strings.Builder{}
					replyMsg.WriteString(text)
					replyMsg.WriteRune(k)

					return c.Reply(replyMsg.String())
				}
			}
		}
	}

	return c.Reply(text)

}

func OnSticker(c tele.Context) error {
	sticker := c.Message().Sticker

	if common.Contains(sticker.Emoji, QuestionMarkEmojis) {
		return c.Reply(sticker)
	}

	return nil
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
	features.RegisterFeature(tele.OnSticker, OnSticker)
}
