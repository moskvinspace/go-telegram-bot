package main

import (
	"github.com/moskvinspace/go-telegram-bot/telegram"
	"os"
)

func main() {
	bot := telegram.Bot{
		Token:             os.Getenv("TELEGRAM_APITOKEN"),
		AllowedTelegramId: os.Getenv("ALLOWED_TELEGRAM_ID"),
	}

	bot.Start()
}
