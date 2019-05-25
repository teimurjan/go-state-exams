package factory

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/teimurjan/go-state-exams/config"
)

// MakeTelegramBot creates a new instance of tgbotapi.BotAPI
func MakeTelegramBot(c *config.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(c.TelegramBotToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = c.Debug

	return bot, nil
}
