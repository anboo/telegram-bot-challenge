package cmd

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Cmd interface {
	Support(update tgbotapi.Update) bool
	Handle(ctx context.Context, api *tgbotapi.BotAPI, update tgbotapi.Update)
}
