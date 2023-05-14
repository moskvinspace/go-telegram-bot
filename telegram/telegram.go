package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

type Bot struct {
	Token             string
	AllowedTelegramId string
	botAPI            *tgbotapi.BotAPI
}

func (b *Bot) Start() {
	var err error

	b.botAPI, err = tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		log.Panic(err)
	}

	b.botAPI.Debug = true
	log.Printf("Authorized on account %s", b.botAPI.Self.UserName)

	config := tgbotapi.UpdateConfig{Timeout: 60}
	updates := b.botAPI.GetUpdatesChan(config)

	initCommands()

	// Loop through each update.
	b.processingUpdates(updates)
}

func (b *Bot) processingUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if !b.isAllowedId(update.SentFrom().ID) {
			continue
		}

		configs := middleware(update)
		for _, config := range configs {
			b.botAPI.Send(config)
		}
	}
}

func middleware(update tgbotapi.Update) (configs []tgbotapi.Chattable) {
	switch {
	case update.Message != nil:
		return handleMessage(update.Message)
	case update.CallbackQuery != nil:
		// todo
	}

	return configs
}

func (b *Bot) isAllowedId(id int64) bool {
	if b.AllowedTelegramId == "" {
		return true
	}

	allowedIds := strings.Split(b.AllowedTelegramId, ",")
	for _, allowedId := range allowedIds {
		if allowedId == strconv.FormatInt(id, 10) {
			return true
		}
	}

	return false
}
