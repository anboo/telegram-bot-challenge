package cmd

import (
	"awesomeProject/client"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UnRegCmd struct{}

func (UnRegCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "unreg"
}

func (UnRegCmd) Handle(api client.TelegramClient, update tgbotapi.Update) {
	txt := fmt.Sprintf("%s, больше ты не участвуешь в игре!", update.Message.From.String())
	msg := tgbotapi.NewMessage(update.FromChat().ID, txt)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := api.GetAPI().Send(msg); err != nil {
		panic(err)
	}
}
