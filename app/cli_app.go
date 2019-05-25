package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	var input string

	reader := bufio.NewReader(os.Stdin)

	for input != "/quit" {
		responseText := ""

		fmt.Println("Enter a text to search:")

		input, _ = reader.ReadString('\n')

		foundQuestions := app.questionsRepo.Search(strings.TrimSpace(input))
		for i, question := range foundQuestions {
			responseText += strconv.Itoa(i+1) + ". " + question.String() + "\n\n"
		}

		fmt.Println(responseText)

	}
}
