package v1

import (
	"questions-generators/internal/providers/ai"
	"questions-generators/internal/providers/websearch"
)

type QuestionService struct{}

func NewQuestionService() *QuestionService {
	return &QuestionService{}
}

func (q *QuestionService) GetQuestions(jobTitle, industry, company string, jobDescription string) []string {
	trends := websearch.FetchLatestTrends(jobTitle, industry, company)
	questions := ai.GenerateInterviewQuestions(jobTitle, industry, company, trends, jobDescription)
	return questions
}
