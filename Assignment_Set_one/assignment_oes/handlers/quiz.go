package handlers

import (
	"assignment_oes/models"
	"encoding/json"
	"net/http"
	"time"
)

// Quiz state
var userScore int
var currentQuestion int
var userResponses = make(map[int]int) // Stores user's selected options by question ID

// FetchQuestionsHandler - Fetch all questions
func FetchQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.QuestionBank)
}

// TakeQuizHandler - Answer a question
func TakeQuizHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse question ID and selected option from request
	var request struct {
		QuestionID int `json:"question_id"`
		Option     int `json:"option"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate question ID
	if request.QuestionID <= 0 || request.QuestionID > len(models.QuestionBank) {
		http.Error(w, "Question not found", http.StatusBadRequest)
		return
	}

	// Validate option index
	question := models.QuestionBank[request.QuestionID-1]
	if request.Option < 0 || request.Option >= len(question.Options) {
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}

	// Record user's response
	userResponses[request.QuestionID] = request.Option
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Answer recorded"))
}

// SubmitQuizHandler - Submit the quiz and calculate score
func SubmitQuizHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Calculate score
	userScore = 0
	for _, question := range models.QuestionBank {
		if userResponses[question.ID] == question.Answer {
			userScore++
		}
	}

	// Evaluate performance
	performance := ""
	switch {
	case userScore == len(models.QuestionBank):
		performance = "Excellent"
	case userScore >= len(models.QuestionBank)/2:
		performance = "Good"
	default:
		performance = "Needs Improvement"
	}

	// Reset quiz state
	userResponses = make(map[int]int)

	// Respond with results
	response := map[string]interface{}{
		"score":       userScore,
		"total":       len(models.QuestionBank),
		"performance": performance,
	}
	json.NewEncoder(w).Encode(response)
}

// Timer (optional)
func QuizTimer(duration int, done chan bool) {
	time.Sleep(time.Duration(duration) * time.Second)
	done <- true
}
