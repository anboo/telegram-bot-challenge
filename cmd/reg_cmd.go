package cmd

import (
	"awesomeProject/client"
	"awesomeProject/db"
	"awesomeProject/translation"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type RegCmd struct {
	UserDAO     db.UsersRepository
	Translation *translation.Translation
}

func (RegCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "reg"
}

func (c RegCmd) Handle(ctx context.Context, api client.TelegramClient, update tgbotapi.Update) {
	user := c.UserDAO.FindUserInChat(
		context.TODO(),
		strconv.Itoa(int(update.Message.From.ID)),
		strconv.Itoa(int(update.FromChat().ID)),
	)

	if user != nil {
		api.SendMessage(update, c.Translation.Trans(translation.RU, translation.YouAlreadyRegisteredMessage, nil))
		return
	}

	err := c.UserDAO.InsertNewUser(
		ctx,
		strconv.Itoa(int(update.FromChat().ID)),
		strconv.Itoa(int(update.Message.From.ID)),
		update.Message.From.String(),
	)
	if err != nil {
		errMsg := tgbotapi.NewMessage(
			update.FromChat().ID,
			c.Translation.Trans(translation.RU, translation.InternalError, nil),
		)
		log.Println("error saving user to database " + err.Error())
		_, errSend := api.GetAPI().Send(errMsg)
		if errSend != nil {
			log.Println("error sending about database error message " + errSend.Error())
		}
		return
	}

	msg := c.Translation.Trans(
		translation.RU,
		translation.YouRegistered,
		&map[string]string{"username": update.Message.From.String()},
	)

	api.SendMessage(update, msg)
}
