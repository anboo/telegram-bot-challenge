package cmd

import (
	"awesomeProject/db"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type RegCmd struct {
	UserDAO db.UsersRepository
}

func (RegCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "reg"
}

func (c RegCmd) Handle(api *tgbotapi.BotAPI, update tgbotapi.Update) {
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
		context.TODO(),
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

	txt := fmt.Sprintf("🤡 Привет, %s. Теперь ты участвуешь в игре!", update.Message.From.String())
	msg := tgbotapi.NewMessage(update.FromChat().ID, txt)

	if _, err := api.Send(msg); err != nil {
		panic(err)
	}
}
