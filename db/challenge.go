package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type ChallengeDAO struct {
	Db *Database
}

func (u ChallengeDAO) FindChallengeToday(ctx context.Context, userId string, chatId string) *User {
	row := u.Db.Conn(ctx).QueryRowContext(
		ctx,
		"SELECT id, telegram_id, chat_id, created_at FROM challenge_result WHERE chat_id = $1 AND telegram_id = $2 LIMIT 1",
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
			log.Printf("error fetching challenge_result from database with chat_id = %s telegram_id = %s error = %s", chatId, userId, err)
		}

		return nil
	}

	return &user
}
