package cmd

import (
	"awesomeProject/db"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type RegCmd struct {
	UserDAO db.UserDAO
}

func (RegCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "reg"
}

func (c RegCmd) Handle(ctx context.Context, api *tgbotapi.BotAPI, update tgbotapi.Update) {
	user := c.UserDAO.FindUserInChat(
		context.TODO(),
		strconv.Itoa(int(update.FromChat().ID)),
		strconv.Itoa(int(update.Message.From.ID)),
	)

	if user != nil {
		m := tgbotapi.NewMessage(update.FromChat().ID, "🗿🗿🗿 Ты уже зарегистрирован в игре")
		_, err := api.Send(m)
		if err != nil {
			log.Println("error send message " + err.Error())
		}
		return
	}

	err := c.UserDAO.InsertNewUser(
		ctx,
		strconv.Itoa(int(update.FromChat().ID)),
		strconv.Itoa(int(update.Message.From.ID)),
		update.Message.From.String(),
	)
	if err != nil {
		errMsg := tgbotapi.NewMessage(update.FromChat().ID, "Произошла техническая ошибка")
		log.Println("error saving user to database " + err.Error())
		_, errSend := api.Send(errMsg)
		if errSend != nil {
			log.Println("error sending about database error message " + errSend.Error())
		}
		return
	}
}
