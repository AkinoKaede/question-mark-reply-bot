package questionmarkreply

import (
	"strings"
	"unicode/utf8"

	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	tele "gopkg.in/telebot.v3"
)

var (
	QuestionMarks      = append([]rune{'?', '¿', '？'}, QuestionMarkEmojis...)
	QuestionMarkEmojis = []rune{'❓', '❔'}
)

func OnText(c tele.Context) error {
	var markCount int
	text := c.Message().Text
	textLen := utf8.RuneCountInString(text)

	markCountMap := make(map[rune]int)
	for _, b := range text {
		markCountMap[b]++
	}

	for _, mark := range QuestionMarks {
		markCount += markCountMap[mark]
	}

	if markCount == textLen {
		if c.Message().ReplyTo != nil && textLen > 1 {
			for k, v := range markCountMap {
				if v == textLen {
					replyToText := c.Message().ReplyTo.Text
					if strings.Count(replyToText, string(k)) == utf8.RuneCountInString(replyToText) {
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

	return nil
}

func OnSticker(c tele.Context) error {
	sticker := c.Message().Sticker
	emoji, _ := utf8.DecodeRuneInString(sticker.Emoji)
	if common.Contains(emoji, QuestionMarkEmojis) {
		return c.Reply(sticker)
	}

	return nil
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
	features.RegisterFeature(tele.OnSticker, OnSticker)
}
