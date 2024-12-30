package main

import (
	"assignment_oes/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/questions", handlers.FetchQuestionsHandler).Methods("GET")
	router.HandleFunc("/quiz/answer", handlers.TakeQuizHandler).Methods("POST")
	router.HandleFunc("/quiz/submit", handlers.SubmitQuizHandler).Methods("POST")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
