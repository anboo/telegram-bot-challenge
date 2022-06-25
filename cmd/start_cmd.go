package cmd

import (
	"awesomeProject/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartCmd struct {
}

func (s StartCmd) Support(update tgbotapi.Update) bool {
	return true
}

func (s StartCmd) Handle(bot client.TelegramClient, update tgbotapi.Update) {
	bot.ReplyMessage(
		update,
		"Привет! Для регистрации в игре вызови /reg в групповом чате. Для старта челленджа дня вызови /challenge",
	)
}
