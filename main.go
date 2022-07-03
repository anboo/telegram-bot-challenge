package main

import (
	"awesomeProject/client"
	"awesomeProject/cmd"
	db2 "awesomeProject/db"
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var handlers []cmd.Cmd

func main() {
	db := db2.CreateDatabase(os.Getenv("POSTGRESQL_DSN"))

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	driver, err := postgres.WithInstance(db.Conn(ctx), &postgres.Config{})
	if err != nil {
		panic(err.Error())
	}

	path, _ := os.Getwd()
	m, err := migrate.NewWithDatabaseInstance("file:///"+path+"/migrations", "postgresql", driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	telegram := client.Telegram{BptAPI: bot}
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

	names := strings.Split(os.Getenv("CHALLENGE_NAME"), ",")
	if len(names) != 2 {
		fmt.Printf("Expected 2 words for plurization. Got value %v", names)
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

	go func() {
		<-ctx.Done()
		log.Println("stop receiving updates")
		bot.StopReceivingUpdates()
	}()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		for _, h := range handlers {
			if h.Support(update) {
				h.Handle(ctx, &telegram, update)
				break
			}
		}
	}
}
