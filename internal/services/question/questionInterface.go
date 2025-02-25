package question

type QuestionInterface interface {
	GetQuestions(jobTitle, industry, company string) []string
}
