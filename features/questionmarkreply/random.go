package questionmarkreply

import (
	"math/rand"
	"time"

	tele "gopkg.in/telebot.v3"
)

func RandomForNotPrivate(probability int) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat().Type == tele.ChatPrivate {
				return next(c)
			}

			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(100)

			if randNum < probability {
				return next(c)
			}

			return nil
		}
	}
}
