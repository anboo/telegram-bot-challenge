package cmd

import (
	"awesomeProject/client"
	"awesomeProject/translation"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UnRegCmd struct {
	Translation *translation.Translation
}

func (UnRegCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "unreg"
}

func (c UnRegCmd) Handle(ctx context.Context, api client.TelegramClient, update tgbotapi.Update) {
	txt := c.Translation.Trans(translation.RU, translation.YouExited, &map[string]string{
		"username": update.Message.From.String(),
	})

	msg := tgbotapi.NewMessage(update.FromChat().ID, txt)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := api.GetAPI().Send(msg); err != nil {
		panic(err)
	}
}
