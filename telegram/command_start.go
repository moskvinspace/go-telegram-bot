package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type startCommand struct{}

func (startCommand) name() string {
	return "/start"
}

func (startCommand) run(req commandRequest) (configs []tgbotapi.Chattable) {
	var text = "Hi!"

	return append(configs, newMessage(messageConfig{chatId: req.message.Chat.ID, text: text}))
}
