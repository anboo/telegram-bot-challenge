package tests

import (
	"awesomeProject/cmd"
	"awesomeProject/translation"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_StartCmd_Support(t *testing.T) {
	regCmd := cmd.StartCmd{}

	var tests = []struct {
		input *tgbotapi.Message
		want  bool
	}{
		{
			input: &tgbotapi.Message{
				Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 6}},
				Text:     "/start",
			},
			want: true,
		},
		{
			input: nil,
			want:  true,
		},
		{
			input: &tgbotapi.Message{
				Text: "Hello!",
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

func Test_StartCmd_StartMessage(t *testing.T) {
	tMock := NewMockTelegramClient(gomock.NewController(t))

	startCmd := cmd.StartCmd{Translation: translation.NewTranslationWithDefault()}

	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			MessageID: 2,
			Entities:  []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 6}},
			Text:      "/start",
		},
	}

	tMock.EXPECT().ReplyMessage(
		gomock.Eq(update),
		gomock.Eq("Привет! Для регистрации в игре вызови /reg в групповом чате. Для старта челленджа дня вызови /challenge"),
	)

	startCmd.Handle(context.TODO(), tMock, update)
}
