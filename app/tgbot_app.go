package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/teimurjan/go-state-exams/config"
	botFactory "github.com/teimurjan/go-state-exams/factory"
	"github.com/teimurjan/go-state-exams/repo"
)

// TgBotApp is an interface for telegram bot application
type TgBotApp interface {
	Start()
}

type tgBotApp struct {
	botAPI        *tgbotapi.BotAPI
	conf          *config.Config
	questionsRepo repo.QuestionRepo
}

// NewTgBotApp creates new tgBotApp instance
func NewTgBotApp(conf *config.Config, questionsRepo repo.QuestionRepo) TgBotApp {
	botAPI, err := botFactory.MakeTelegramBot(conf)
	if err != nil {
		panic(fmt.Sprintln(err))
	}

	return &tgBotApp{botAPI, conf, questionsRepo}
}

// NewTgBotApp creates new tg bot application
func (tgBotApp *tgBotApp) Start() {
	updates, err := tgBotApp.getUpdates()
	if err != nil {
		panic(fmt.Sprintln(err))
	}

	for update := range updates {
		if update.Message != nil && len(update.Message.Text) > 0 {
			go tgBotApp.handleText(&update)
		}
	}
}

func (tgBotApp *tgBotApp) getUpdates() (tgbotapi.UpdatesChannel, error) {
	if !tgBotApp.conf.UseWebhook {
		return tgBotApp.setupPolling()
	}
	return tgBotApp.setupWebhook()
}

func (tgBotApp *tgBotApp) setupPolling() (tgbotapi.UpdatesChannel, error) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 5
	fmt.Println("Start polling.")
	return tgBotApp.botAPI.GetUpdatesChan(updateConfig)
}

func (tgBotApp *tgBotApp) setupWebhook() (tgbotapi.UpdatesChannel, error) {
	webhookURL := tgBotApp.conf.HerokuBaseURL + "/" + tgBotApp.conf.TelegramBotToken
	_, err := tgBotApp.botAPI.SetWebhook(
		tgbotapi.NewWebhook(webhookURL),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	updates := tgBotApp.botAPI.ListenForWebhook("/" + tgBotApp.conf.TelegramBotToken)
	go http.ListenAndServe(":"+tgBotApp.conf.Port, nil)

	fmt.Println("Listening port " + tgBotApp.conf.Port + ". Webhook url is " + webhookURL + ".")
	return updates, nil
}

func (tgBotApp *tgBotApp) handleText(update *tgbotapi.Update) {
	msgText := ""
	foundQuestions := tgBotApp.questionsRepo.Search(update.Message.Text)
	for i, question := range foundQuestions {
		answer, ok := question.Variants[question.Answer]
		if !ok {
			answer = question.Answer
		}

		msgText += strconv.Itoa(i+1) + ". " + question.Title + "\n*Answer: " + answer + "*\n\n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	msg.ParseMode = tgbotapi.ModeMarkdown
	tgBotApp.botAPI.Send(msg)
}
