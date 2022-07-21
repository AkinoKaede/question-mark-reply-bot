package features

import (
	tb "gopkg.in/telebot.v3"
)

// type FeatureFunc func(context.Context) tb.HandlerFunc

var features = make(map[interface{}]tb.HandlerFunc)

func RegisterFeature(endpoint interface{}, feature tb.HandlerFunc) {
	features[endpoint] = feature
}

func Handle(b *tb.Bot) {

	for endpoint, feature := range features {
		b.Handle(endpoint, feature)
	}
}
