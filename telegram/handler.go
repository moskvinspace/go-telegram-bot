package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(message *tgbotapi.Message) (configs []tgbotapi.Chattable) {
	cmd, ok := commands[getCommand(message.Text)]
	if !ok {
		msg := newMessage(messageConfig{
			chatId:           message.Chat.ID,
			replyToMessageId: message.MessageID,
			text:             "Invalid command!",
		})

		return append(configs, msg)
	}

	return cmd.run(commandRequest{message})
}
