package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	_ "github.com/AkinoKaede/question-mark-reply-bot/main/distro/all"

	tele "gopkg.in/telebot.v3"
)

var (
	bot *tele.Bot
)

func Handler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	common.Must(err)

	var u tele.Update
	common.Must(json.Unmarshal(body, &u))

	bot.ProcessUpdate(u)
}

func init() {
	b, err := tele.NewBot(tele.Settings{
		Token:       os.Getenv("QMRBOT_TELEGRAM_TOKEN"),
		Synchronous: true,
	})
	common.Must(err)

	features.Handle(b)
}
