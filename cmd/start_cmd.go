package cmd

import (
	"awesomeProject/client"
	"awesomeProject/translation"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartCmd struct {
	Translation *translation.Translation
}

func (s StartCmd) Support(update tgbotapi.Update) bool {
	return true
}

func (s StartCmd) Handle(ctx context.Context, bot client.TelegramClient, update tgbotapi.Update) {
	bot.ReplyMessage(
		update,
		s.Translation.Trans(translation.RU, translation.HelloStartMessage, nil),
	)
}
