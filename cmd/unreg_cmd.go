package cmd

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UnRegCmd struct{}

func (UnRegCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "unreg"
}

func (UnRegCmd) Handle(ctx context.Context, api *tgbotapi.BotAPI, update tgbotapi.Update) {
	txt := fmt.Sprintf("%s, больше ты не участвуешь в игре!", update.Message.From.String())
	msg := tgbotapi.NewMessage(update.FromChat().ID, txt)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := api.Send(msg); err != nil {
		panic(err)
	}
}
