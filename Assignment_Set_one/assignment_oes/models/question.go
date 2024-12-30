package models

// Question struct for a question in the quiz
type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"` // Correct answer index
}

// Slice to store the question bank
var QuestionBank = []Question{
	{
		ID:       1,
		Question: "What is the capital of France?",
		Options:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		Answer:   2, // Index of "Paris"
	},
	{
		ID:       2,
		Question: "What is 5 + 7?",
		Options:  []string{"10", "12", "14", "15"},
		Answer:   1, // Index of "12"
	},
	{
		ID:       3,
		Question: "Who wrote 'To Kill a Mockingbird'?",
		Options:  []string{"Harper Lee", "J.K. Rowling", "Jane Austen", "Mark Twain"},
		Answer:   0, // Index of "Harper Lee"
	},
}
