package tests

import (
	"awesomeProject/cmd"
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
				Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 4}},
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

func TestHandle(t *testing.T) {

}
