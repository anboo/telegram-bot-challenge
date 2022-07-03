package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TelegramClient interface {
	SendMessage(update tgbotapi.Update, text string)
	ReplyMessage(update tgbotapi.Update, text string)
	GetAPI() *tgbotapi.BotAPI
}

type Telegram struct {
	BptAPI *tgbotapi.BotAPI
}

func (t Telegram) SendMessage(update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(
		update.FromChat().ID,
		text,
	)

	t.send(msg)
}

func (t Telegram) ReplyMessage(update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(
		update.FromChat().ID,
		text,
	)
	msg.ReplyToMessageID = update.Message.MessageID

	t.send(msg)
}

func (t Telegram) send(c tgbotapi.Chattable) {
	_, err := t.GetAPI().Send(c)
	if err != nil {
		log.Print("telegram error " + err.Error())
	}
}

func (t Telegram) GetAPI() *tgbotapi.BotAPI {
	return t.BptAPI
}
