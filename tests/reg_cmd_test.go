package tests

import (
	"awesomeProject/cmd"
	"awesomeProject/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSupport(t *testing.T) {
	regCmd := cmd.RegCmd{
		UserDAO: NewMockUsersRepository(gomock.NewController(t)),
	}

	var tests = []struct {
		input *tgbotapi.Message
		want  bool
	}{
		{
			input: &tgbotapi.Message{
				Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 6}},
				Text:     "/start",
			},
			want: false,
		},
		{
			input: nil,
			want:  false,
		},
		{
			input: &tgbotapi.Message{
				Text: "Hello!",
			},
			want: false,
		},
		{
			input: &tgbotapi.Message{
				Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: len("/reg")}},
				Text:     "/reg",
			},
			want: true,
		},
	}

	for _, v := range tests {
		tgBotUpdate := tgbotapi.Update{
			Message: v.input,
		}

		res := regCmd.Support(tgBotUpdate)
		if res != v.want {
			t.Errorf("regCmd.Support expected %v got %v", v.want, res)
		}
	}
}

func TestHandle_RegisterUser(t *testing.T) {
	userDaoMock := NewMockUsersRepository(gomock.NewController(t))
	tMock := NewMockTelegramClient(gomock.NewController(t))

	userDaoMock.
		EXPECT().
		FindUserInChat(gomock.Any(), gomock.Eq("10"), gomock.Eq("1")).
		Return(nil)

	userDaoMock.
		EXPECT().
		InsertNewUser(gomock.Any(), gomock.Eq("1"), gomock.Eq("10"), gomock.Eq("devanboo")).
		Return(nil)

	regCmd := cmd.RegCmd{
		UserDAO: userDaoMock,
	}

	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From:     &tgbotapi.User{ID: 10, UserName: "devanboo"},
			Chat:     &tgbotapi.Chat{ID: 1},
			Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: len("/reg")}},
			Text:     "/reg",
		},
	}

	tMock.EXPECT().SendMessage(gomock.Eq(update), "🤡 Привет, devanboo. Теперь ты участвуешь в игре!")

	regCmd.Handle(tMock, update)
}

func TestHandle_UserAlreadyExists(t *testing.T) {
	userDaoMock := NewMockUsersRepository(gomock.NewController(t))
	tMock := NewMockTelegramClient(gomock.NewController(t))

	userDaoMock.
		EXPECT().
		FindUserInChat(gomock.Any(), gomock.Eq("10"), gomock.Eq("1")).
		Return(&db.User{
			Id:         "f225543b-921b-4dc2-a604-b42a0db6013c",
			Username:   "devanboo",
			ChatId:     "1",
			TelegramId: "10",
		})

	regCmd := cmd.RegCmd{UserDAO: userDaoMock}

	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From:     &tgbotapi.User{ID: 10},
			Chat:     &tgbotapi.Chat{ID: 1},
			Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: len("/reg")}},
			Text:     "/reg",
		},
	}

	tMock.EXPECT().SendMessage(gomock.Eq(update), "🗿🗿🗿 Ты уже зарегистрирован в игре")

	regCmd.Handle(tMock, update)
}
