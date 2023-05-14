package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type Bot struct {
	Token   string
	OwnerId string
	Debug   bool
	BotAPI  *tgbotapi.BotAPI
}

func (b *Bot) Start() {
	var err error

	b.BotAPI, err = tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		log.Panic(err)
	}

	b.BotAPI.Debug = b.Debug
	log.Printf("Authorized on account %s", b.BotAPI.Self.UserName)

	config := tgbotapi.UpdateConfig{Timeout: 60}
	updates := b.BotAPI.GetUpdatesChan(config)

	initCommands()

	// Loop through each update.
	b.processingUpdates(updates)
}

func (b *Bot) processingUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if b.isOwnerExist() && strconv.FormatInt(update.SentFrom().ID, 10) != b.OwnerId {
			continue
		}

		configs := middleware(update)
		for _, config := range configs {
			b.BotAPI.Send(config)
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

func (b *Bot) isOwnerExist() bool {
	return b.OwnerId != ""
}
