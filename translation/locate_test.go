package translation

import "testing"

func TestTranslation_Trans(t *testing.T) {
	trans := NewTranslation(map[string]map[string]string{
		"ru": {
			"hello": "Здравствуй, %username%!!!",
		},
	})

	res := trans.Trans("ru", "hello", map[string]string{
		"username": "Bob",
	})

	if res != "Здравствуй, Bob!!!" {
		t.Errorf("Expected %v got %v", "Здравствуй, Bob!!!", res)
	}
}
