package main

import (
	"assignment_bts/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/account/create", handlers.CreateAccountHandler)
	http.HandleFunc("/account/deposit", handlers.DepositHandler)
	http.HandleFunc("/account/withdraw", handlers.WithdrawHandler)
	http.HandleFunc("/account/balance", handlers.ViewBalanceHandler)
	http.HandleFunc("/account/transactions", handlers.ViewTransactionHistoryHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
