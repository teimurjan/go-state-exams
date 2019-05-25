package app

import (
	"fmt"
	"strconv"

	"github.com/teimurjan/go-state-exams/repo"
)

// CLIApp is a CLI application object
type CLIApp struct {
	questionsRepo repo.QuestionRepo
}

// NewCLIApp creates new CLI app
func NewCLIApp(questionsRepo repo.QuestionRepo) *CLIApp {
	return &CLIApp{questionsRepo}
}

// Start starts CLI app
func (app *CLIApp) Start() {
	var responseText string
	var input string

	for input != "/quit" {
		fmt.Println("Enter a text to search:")

		fmt.Scanln(&input)

		foundQuestions := app.questionsRepo.Search(input)
		for i, question := range foundQuestions {
			responseText += strconv.Itoa(i+1) + ". " + question.String() + "\n\n"
		}

		fmt.Println(responseText)
	}
}
