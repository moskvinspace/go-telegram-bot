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
}

func (b *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = b.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	config := tgbotapi.UpdateConfig{Timeout: 60}
	updates := getBotUpdatesChannel(bot, config)

	initCommands()

	// Loop through each update.
	b.checkingBotUpdates(bot, updates)
}

func (b *Bot) checkingBotUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if b.isOwnerExist() {
			if strconv.FormatInt(update.SentFrom().ID, 10) != b.OwnerId {
				continue
			}
		}

		configs := middleware(update)
		for _, config := range configs {
			bot.Send(config)
		}
	}
}

func getBotUpdatesChannel(bot *tgbotapi.BotAPI, config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	return bot.GetUpdatesChan(config)
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
