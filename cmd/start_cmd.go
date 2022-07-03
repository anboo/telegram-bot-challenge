package cmd

import (
	"awesomeProject/translation"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type StartCmd struct {
	translation *translation.Translation
}

func (s StartCmd) Support(update tgbotapi.Update) bool {
	return true
}

func (s StartCmd) Handle(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Привет! Для регистрации в игре вызови /reg в групповом чате. Для старта челленджа дня вызови /challenge",
	)
	msg.ReplyToMessageID = update.Message.MessageID

	strings.Replace("%name%", "", "", 1)

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
