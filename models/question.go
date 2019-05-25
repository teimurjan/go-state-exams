package models

import (
	"regexp"

	"github.com/teimurjan/go-state-exams/utils"
)

// Question model
type Question struct {
	Title    string
	Variants map[string]string
	Answer   string
}

// NewQuestion creates a new question from string
func NewQuestion(questionStr string) *Question {
	questionRegex := regexp.MustCompile(`QUESTION: (?P<question>.*)\na\) (?P<a>.*)\nb\) (?P<b>.*)(\nc\) (?P<c>.*))?(\nd\) (?P<d>.*))?\nAnswer: (?P<answer>.*)\n`)

	namedMap := utils.GetNamedMap(questionStr, questionRegex)

	question := Question{Variants: make(map[string]string, 4)}

	if title, ok := namedMap["question"]; ok {
		question.Title = title
	}

	if a, ok := namedMap["a"]; ok {
		question.Variants["a"] = a
	}

	if b, ok := namedMap["b"]; ok {
		question.Variants["b"] = b
	}

	if c, ok := namedMap["c"]; ok {
		question.Variants["c"] = c
	}

	if d, ok := namedMap["d"]; ok {
		question.Variants["d"] = d
	}

	if answer, ok := namedMap["answer"]; ok {
		question.Answer = answer
	}

	return &question
}

func (q *Question) String() string {
	answer, ok := q.Variants[q.Answer]
	if !ok {
		answer = q.Answer
	}

	return q.Title + "\n*Answer: " + answer
}
