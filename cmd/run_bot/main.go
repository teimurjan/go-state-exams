package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/teimurjan/go-state-exams/app"
	"github.com/teimurjan/go-state-exams/config"
	"github.com/teimurjan/go-state-exams/repo"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintln(err))
	}

	questionsRepo := repo.NewQuestionsRepo(conf.QuestionsFile)

	app := app.NewTgBotApp(conf, questionsRepo)
	app.Start()
}
