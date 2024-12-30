package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

// Account struct to hold account details
type Account struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Balance           float64  `json:"balance"`
	TransactionHistory []string `json:"transaction_history"`
}

// Slice to store accounts
var accounts []Account
var mu sync.Mutex // Mutex for thread-safe access

// CreateAccountHandler handles account creation
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newAccount Account
	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if ID is unique
	for _, acc := range accounts {
		if acc.ID == newAccount.ID {
			http.Error(w, "Account ID must be unique", http.StatusBadRequest)
			return
		}
	}

	newAccount.TransactionHistory = []string{strconv.FormatFloat(newAccount.Balance, 'f', 2, 64) + " Added!"}

	// Add new account
	accounts = append(accounts, newAccount)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAccount)
}

// DepositHandler handles deposit transactions
func DepositHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	type DepositRequest struct {
		ID     int     `json:"id"`
		Amount float64 `json:"amount"`
	}

	var req DepositRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate deposit amount
	if req.Amount <= 0 {
		http.Error(w, "Deposit amount must be greater than zero", http.StatusBadRequest)
		return
	}

	// Find account and deposit
	for i, acc := range accounts {
		if acc.ID == req.ID {
			accounts[i].Balance += req.Amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, "Deposited: "+strconv.FormatFloat(req.Amount, 'f', 2, 64))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accounts[i])
			return
		}
	}
	http.Error(w, "Account not found", http.StatusNotFound)
}

// WithdrawHandler handles withdrawal transactions
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	type WithdrawRequest struct {
		ID     int     `json:"id"`
		Amount float64 `json:"amount"`
	}

	var req WithdrawRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate withdraw amount
	if req.Amount <= 0 {
		http.Error(w, "Withdraw amount must be greater than zero", http.StatusBadRequest)
		return
	}

	// Find account and withdraw
	for i, acc := range accounts {
		if acc.ID == req.ID {
			if acc.Balance < req.Amount {
				http.Error(w, "Insufficient balance", http.StatusBadRequest)
				return
			}
			accounts[i].Balance -= req.Amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, "Withdrawn: "+strconv.FormatFloat(req.Amount, 'f', 2, 64))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accounts[i])
			return
		}
	}
	http.Error(w, "Account not found", http.StatusNotFound)
}

// ViewBalanceHandler handles balance inquiries
func ViewBalanceHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	for _, acc := range accounts {
		if acc.ID == id {
			json.NewEncoder(w).Encode(map[string]float64{"balance": acc.Balance})
			return
		}
	}
	http.Error(w, "Account not found", http.StatusNotFound)
}

// ViewTransactionHistoryHandler handles transaction history inquiries
func ViewTransactionHistoryHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	for _, acc := range accounts {
		if acc.ID == id {
			json.NewEncoder(w).Encode(acc.TransactionHistory)
			return
		}
	}
	http.Error(w, "Account not found", http.StatusNotFound)
}
