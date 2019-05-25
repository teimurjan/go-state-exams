package config

import "github.com/kelseyhightower/envconfig"

// Config object
type Config struct {
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN" required:"true"`
	UseWebhook       bool   `envconfig:"USE_WEBHOOK"`
	HerokuBaseURL    string `envconfig:"HEROKU_BASE_URL"`
	Debug            bool   `envconfig:"DEBUG"`
	Port             string `envconfig:"PORT"`
	QuestionsFile    string `envconfig:"QUESTIONS_FILE" default:"questions.txt"`
}

// NewConfig creates new config
func NewConfig() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return &c, err
	}
	return &c, nil
}
