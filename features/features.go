package features

import (
	tb "gopkg.in/telebot.v3"
)

type Feature struct {
	Handler    tb.HandlerFunc
	Middleware []tb.MiddlewareFunc
}

var features = make(map[interface{}]*Feature)

func RegisterFeature(endpoint interface{}, handler tb.HandlerFunc, middleware ...tb.MiddlewareFunc) {
	features[endpoint] = &Feature{
		Handler:    handler,
		Middleware: middleware,
	}
}

func Handle(b *tb.Bot) {
	for endpoint, feature := range features {
		b.Handle(endpoint, feature.Handler, feature.Middleware...)
	}
}
