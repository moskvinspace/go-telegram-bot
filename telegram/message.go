package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type messageConfig struct {
	chatId           int64
	replyMarkup      interface{}
	replyToMessageId int
	text             string
}

func newMessage(cfg messageConfig) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           cfg.chatId,
			ReplyMarkup:      cfg.replyMarkup,
			ReplyToMessageID: cfg.replyToMessageId,
		},
		Text: cfg.text,
	}
}
