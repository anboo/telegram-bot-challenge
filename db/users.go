package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id         string
	Username   string
	TelegramId string
	ChatId     string
	CreatedAt  uint
}

type UserDAO struct {
	Db *Database
}

func (u *UserDAO) FindUserInChat(ctx context.Context, userId string, chatId string) *User {
	row := u.Db.GetDatabase().QueryRowContext(
		ctx,
		"SELECT id, username, telegram_id, chat_id, created_at FROM users WHERE chat_id = $1 AND telegram_id = $2 LIMIT 1",
		chatId,
		userId,
	)

	if row == nil {
		return nil
	}

	user := User{}
	err := row.Scan(&user.Id, &user.Username, &user.TelegramId, &user.ChatId, &user.CreatedAt)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("error fetching from database with chat_id = %s telegram_id = %s error = %s", chatId, userId, err)
		}

		return nil
	}

	return &user
}

func (u *UserDAO) InsertNewUser(ctx context.Context, userId string, chatId string, username string) error {
	_, err := u.Db.GetDatabase().ExecContext(
		ctx,
		"INSERT INTO users (id, username, telegram_id, chat_id, created_at) "+
			"VALUES (uuid_generate_v4(), $1, $2, $3, $4) ON CONFLICT DO NOTHING",
		username,
		chatId,
		userId,
		time.Now().UnixNano(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserDAO) FindUsernamesInChat(ctx context.Context, chatId string) ([]string, error) {
	rows, err := u.Db.GetDatabase().QueryContext(ctx, "SELECT username FROM users WHERE chat_id = $1", chatId)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("error close rows " + err.Error())
		}
	}(rows)

	var usernames []string

	for rows.Next() {
		var username = ""
		err := rows.Scan(&username)
		if err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}

	if rows.Err() != nil {
		return usernames, err
	}

	return usernames, nil
}
