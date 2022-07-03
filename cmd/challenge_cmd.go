package cmd

import (
	"awesomeProject/client"
	"awesomeProject/db"
	"awesomeProject/translation"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type ChallengeCmd struct {
	UserDAO     db.UsersRepository
	Translation *translation.Translation
	Name        []string
}

func (ChallengeCmd) Support(update tgbotapi.Update) bool {
	return update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == "challenge"
}

func (c ChallengeCmd) Handle(ctx context.Context, api client.TelegramClient, update tgbotapi.Update) {
	c.Translation.Trans(translation.RU, translation.ChallengeMessage, &map[string]string{
		"name":     c.Name[1],
		"chatName": update.Message.Chat.Title,
	})

	msg := tgbotapi.NewMessage(
		update.FromChat().ID,
		c.Translation.Trans(translation.RU, translation.ChallengeMessage, &map[string]string{
			"name":     c.Name[1],
			"chatName": update.Message.Chat.Title,
		}),
	)

	_, err := api.GetAPI().Send(msg)
	if err != nil {
		log.Print("telegram error " + err.Error())
	}

	usernames, err := c.UserDAO.FindUsernamesInChat(
		ctx,
		strconv.Itoa(int(update.FromChat().ID)),
	)

	var randMessage tgbotapi.MessageConfig
	if err != nil {
		log.Println("error find usernames in chat " + err.Error())

		randMessage = tgbotapi.NewMessage(
			update.FromChat().ID,
			c.Translation.Trans(translation.RU, translation.ErrorMessage, &map[string]string{
				"name": c.Name[1],
			}),
		)
	} else {
		if len(usernames) < 2 {
			randMessage = tgbotapi.NewMessage(
				update.FromChat().ID,
				c.Translation.Trans(translation.RU, translation.MaxPlayersMessage, &map[string]string{
					"name": c.Name[1],
				}),
			)
		} else {
			rand.Seed(time.Now().UnixNano())

			randMessage = tgbotapi.NewMessage(
				update.FromChat().ID,
				c.Translation.Trans(translation.RU, translation.MaxPlayersMessage, &map[string]string{
					"name":     c.Name[0],
					"username": usernames[rand.Intn(len(usernames))],
				}),
			)
		}
	}

	_, err = api.GetAPI().Send(randMessage)
	if err != nil {
		fmt.Println("cannot send challenge message " + err.Error())
	}
}
