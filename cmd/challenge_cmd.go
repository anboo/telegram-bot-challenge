package cmd

import (
	"awesomeProject/db"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type ChallengeCmd struct {
	UserDAO db.UserDAO
	Name    []string
}

func (ChallengeCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "challenge"
}

func (c ChallengeCmd) Handle(api *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(
		update.FromChat().ID,
		"🤡🤡🤡 Итак, начинаем искать "+c.Name[1]+" дня в "+update.Message.Chat.Title,
	)

	_, err := api.Send(msg)
	if err != nil {
		log.Print("telegram error " + err.Error())
	}

	usernames, err := c.UserDAO.FindUsernamesInChat(
		context.TODO(),
		strconv.Itoa(int(update.FromChat().ID)),
	)

	var randMessage tgbotapi.MessageConfig
	if err != nil {
		log.Println("error find usernames in chat " + err.Error())

		randMessage = tgbotapi.NewMessage(
			update.FromChat().ID,
			"🤡🤡🤡 "+c.Name[0]+" дня - разработчик бота, потому что произошла ошибка, попробуйте снова чуть позже...",
		)
	} else {
		if len(usernames) < 2 {
			randMessage = tgbotapi.NewMessage(
				update.FromChat().ID,
				"🤡 Для выбора "+c.Name[1]+" дня нужно чтобы было не меньше 2 игроков",
			)
		} else {
			rand.Seed(time.Now().UnixNano())

			randMessage = tgbotapi.NewMessage(
				update.FromChat().ID,
				"Поздравляю!!! 🤡🤡🤡 Ты "+c.Name[0]+" дня, @"+usernames[rand.Intn(len(usernames))],
			)
		}
	}

	_, err = api.Send(randMessage)
	if err != nil {
		fmt.Println("cannot send challenge message " + err.Error())
	}
}
