package cmd

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Cmd interface {
	Support(update tgbotapi.Update) bool
	Handle(api *tgbotapi.BotAPI, update tgbotapi.Update)
}
