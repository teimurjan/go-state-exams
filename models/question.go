package models

import (
	"regexp"

	"github.com/teimurjan/go-state-exams/utils"
)

// Question model
type Question struct {
	Title       string
	Variants    map[string]string
	Answer      string
	Explanation string
}

// NewQuestion creates a new question from string
func NewQuestion(questionStr string) *Question {
	questionRegex := regexp.MustCompile(`QUESTION: (?P<question>.*)\n(a\)\s(?P<a>.*)\n)?(b\)\s(?P<b>.*)\n)?(c\)\s(?P<c>.*)\n)?(d\)\s(?P<d>.*)\n)?(Answer:\s(?P<answer>.*)\n)?(Explanation:\s(?P<explanation>.*)\n)?`)

	namedMap := utils.GetNamedMap(questionStr, questionRegex)

	question := Question{Variants: make(map[string]string, 4)}

	if title, ok := namedMap["question"]; ok {
		question.Title = title
	}

	if explanation, ok := namedMap["explanation"]; ok {
		question.Explanation = explanation
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

	title := ""
	if len(q.Title) > 0 {
		title = q.Title
	} else if len(q.Explanation) > 0 && q.Explanation != "None." {
		title = q.Explanation
	} else {
		title = "Unnamed question with variants: " + q.VariantsString()
	}

	return title + "\n*Answer: " + answer + "*"
}

// SearchString returns a unique string to perform search
func (q *Question) SearchString() string {
	return q.Title + " " + q.VariantsString()
}

// VariantsString returns a string repr of the question variants
func (q *Question) VariantsString() string {
	variantsStr := ""

	for k, v := range q.Variants {
		variantsStr += k + ")" + v + " "
	}

	return variantsStr
}
