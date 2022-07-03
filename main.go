package main

import (
	"awesomeProject/cmd"
	db2 "awesomeProject/db"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

var handlers []cmd.Cmd

func main() {
	db := db2.CreateDatabase(os.Getenv("POSTGRESQL_DSN"))

	ctx := context.Background()

	driver, err := postgres.WithInstance(db.Conn(ctx), &postgres.Config{})
	if err != nil {
		panic(err.Error())
	}

	path, _ := os.Getwd()
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+path+"/migrations",
		"postgresql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		fmt.Printf("migration info: " + err.Error() + "\r\n")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	userDAO := db2.UserDAO{
		Db: db,
	}

	handlers = append(
		handlers,
		cmd.RegCmd{
			UserDAO: userDAO,
		},
		cmd.ChallengeCmd{
			UserDAO: userDAO,
			Name:    []string{"клоун", "клоуна"},
		},
		cmd.UnRegCmd{},
		cmd.StartCmd{},
	)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		for _, h := range handlers {
			if h.Support(update) {
				h.Handle(ctx, bot, update)
				break
			}
		}
	}
}
