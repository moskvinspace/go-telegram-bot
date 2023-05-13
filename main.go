package main

import (
	"github.com/moskvinspace/go-telegram-bot/telegram"
	"os"
)

func main() {
	bot := telegram.Bot{
		Token:   os.Getenv("TELEGRAM_APITOKEN"),
		OwnerId: os.Getenv("OWNER_ID"),
		Debug:   true,
	}

	bot.Start()
}
