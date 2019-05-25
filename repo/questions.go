package repo

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/teimurjan/go-state-exams/models"
)

const MaxResultsPerQuery = 10

// QuestionRepo is repository interface
type QuestionRepo interface {
	GetAll() []models.Question
	Search(query string) []models.Question
}

type questionRepo struct {
	questions []models.Question
}

// NewQuestionsRepo creates a new question repo instance
func NewQuestionsRepo(filepath string) QuestionRepo {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintln(err))
	}

	fileStr := string(file)

	questionStrs := strings.Split(fileStr, "\n\n")

	questions := make([]models.Question, len(questionStrs))

	for i, questionStr := range questionStrs {
		questions[i] = *models.NewQuestion(questionStr)
	}

	return &questionRepo{questions}
}

// Returns all questions
func (repo *questionRepo) GetAll() []models.Question {
	return repo.questions
}

// Returns questions matching query
func (repo *questionRepo) Search(query string) []models.Question {
	matchingQuestions := make([]models.Question, 0, MaxResultsPerQuery)
	for _, question := range repo.questions {
		if strings.Contains(question.Title, query) {
			matchingQuestions = append(matchingQuestions, question)

			if len(matchingQuestions) == MaxResultsPerQuery {
				return matchingQuestions
			}
		}
	}

	return matchingQuestions
}