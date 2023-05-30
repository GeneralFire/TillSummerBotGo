package commandhandlers

//go:generate minimock -i Subscriber -o ./ -s "_minimock.go"
//go:generate minimock -i Unsubscriber -o ./ -s "_minimock.go"

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Subscriber interface {
	SubscribeChat(id int64) error
}
type Unsubscriber interface {
	UnsubscribeChat(id int64) error
}

func GetSubscribeHandler(p Subscriber) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		err := p.SubscribeChat(update.Message.Chat.ID)
		if err != nil {
			return tgbotapi.NewMessage(
				update.Message.Chat.ID,
				err.Error(),
			)
		}
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Success",
		)
	}
}

func GetUnsubscribeHandler(p Unsubscriber) func(update tgbotapi.Update) tgbotapi.MessageConfig {
	return func(update tgbotapi.Update) tgbotapi.MessageConfig {
		err := p.UnsubscribeChat(update.Message.Chat.ID)
		if err != nil {
			return tgbotapi.NewMessage(
				update.Message.Chat.ID,
				err.Error(),
			)
		}
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Success :(",
		)
	}
}
