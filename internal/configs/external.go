package configs

type (
	ExternalConfig struct {
		Telegram TelegramConfig
	}
	TelegramConfig struct {
		TelegramTo    int64  `yaml:"telegramTo"`
		TelegramToken string `yaml:"telegramToken"`
	}
)
