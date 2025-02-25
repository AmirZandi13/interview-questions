package handlers

import (
	"encoding/json"
	"net/http"
	"questions-generators/internal/services/question/v1"
)

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	jobTitle := r.URL.Query().Get("jobTitle")
	industry := r.URL.Query().Get("industry")
	company := r.URL.Query().Get("company")
	jobDescription := r.URL.Query().Get("jobDescription")

	if jobTitle == "" || industry == "" || company == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	service := v1.NewQuestionService()
	questions := service.GetQuestions(jobTitle, industry, company, jobDescription)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions) // Directly return an array of questions
}
