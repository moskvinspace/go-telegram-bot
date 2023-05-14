package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"os"
)

type aiCommand struct{}

func (aiCommand) name() string {
	return "Ai"
}

func (aiCommand) run(req commandRequest) (configs []tgbotapi.Chattable) {
	text := getCommandText(req.message.Text)

	if len(text) == 0 {
		return append(configs, newMessage(messageConfig{
			chatId: req.message.Chat.ID,
			text:   "Ai [message to OpenAI]"}),
		)
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		return append(configs, newMessage(messageConfig{
			chatId:           req.message.Chat.ID,
			replyToMessageId: req.message.MessageID,
			text:             fmt.Sprintf("ChatCompletion error: %v\n", err)}),
		)
	}

	return append(configs, newMessage(messageConfig{
		chatId:           req.message.Chat.ID,
		replyToMessageId: req.message.MessageID,
		text:             resp.Choices[0].Message.Content,
	}))
}
