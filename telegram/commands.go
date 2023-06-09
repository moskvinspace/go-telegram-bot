package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type commandRequest struct {
	message *tgbotapi.Message
}

type command interface {
	name() string
	run(req commandRequest) (configs []tgbotapi.Chattable)
}

var commands map[string]command

func initCommands() {
	commands = map[string]command{
		startCommand{}.name(): &startCommand{},
		aiCommand{}.name():    &aiCommand{},
	}
}

func getCommand(s string) string {
	return strings.Split(s, " ")[0]
}

func getCommandText(s string) string {
	return strings.Join(strings.Split(s, " ")[1:], " ")
}
