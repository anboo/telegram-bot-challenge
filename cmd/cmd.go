package cmd

import (
	"awesomeProject/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Cmd interface {
	Support(update tgbotapi.Update) bool
	Handle(api client.TelegramClient, update tgbotapi.Update)
}
