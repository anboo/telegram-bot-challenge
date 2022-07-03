package translation

import "strings"

var templates = map[string]map[string]string{
	RU: {
		helloStartMessage:           "Привет! Для регистрации в игре вызови /reg в групповом чате. Для старта челленджа дня вызови /challenge",
		challengeMessage:            "🤡🤡🤡 Итак, начинаем искать %name% дня в %chatName%",
		errorMessage:                "🤡🤡🤡 %name% дня - разработчик бота, потому что произошла ошибка, попробуйте снова чуть позже...",
		maxPlayersMessage:           "🤡 Для выбора %name% дня нужно чтобы было не меньше 2 игроков",
		resultMessage:               "Поздравляю!!! 🤡🤡🤡 Ты %name% дня, @%username%",
		youAlreadyRegisteredMessage: "🗿🗿🗿 Ты уже зарегистрирован в игре",
		internalError:               "Произошла техническая ошибка",
		youExited:                   "%username%, больше ты не участвуешь в игре!",
	},
}

type Translation struct {
	templates     map[string]map[string]string
	defaultLocale string
}

func NewTranslationWithDefault() *Translation {
	return &Translation{templates: templates}
}

func NewTranslation(templates map[string]map[string]string) *Translation {
	return &Translation{templates: templates}
}

func (t *Translation) Trans(locale string, code string, args map[string]string) string {
	v, ok := t.templates[locale]

	if !ok && t.defaultLocale != "" && t.defaultLocale != locale {
		locale = t.defaultLocale
	}

	val, ok := v[code]
	if !ok {
		return ""
	}

	for transKey, transVal := range args {
		val = strings.Replace(val, "%"+transKey+"%", transVal, 1)
	}

	return val
}
