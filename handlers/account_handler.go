package handlers

import (
	"awesome-bank/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	accountService *services.AccountService
	logger         *logrus.Logger
}

func NewAccountHandler(accountService *services.AccountService, logger *logrus.Logger) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
		logger:         logger,
	}
}

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("userID")
	if userIDInterface == nil {
		a.logger.Warnf("Unauthorized: %v", userIDInterface)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		a.logger.Warnf("Invalid user ID type: %v", userID)
		http.Error(w, "Invalid user ID type", http.StatusInternalServerError)
		return
	}

	var req struct {
		Type string `json:"type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.logger.Warnf("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		a.logger.Warnf("Missing account type")
		http.Error(w, "Missing account type", http.StatusBadRequest)
		return
	}

	err := a.accountService.CreateAccount(userID, req.Type)
	if err != nil {
		a.logger.Warnf("Failed to create account: %v", err)
		http.Error(w, fmt.Sprintf("Failed to create account: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account created"})
}

func (a *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("userID")
	if userIDInterface == nil {
		a.logger.Warnf("Unauthorized: %v", userIDInterface)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		a.logger.Warnf("Invalid user ID type: %v", userID)
		http.Error(w, "Invalid user ID type", http.StatusInternalServerError)
		return
	}

	accounts, err := a.accountService.GetAccounts(userID)

	if err != nil {
		a.logger.Warnf("Login failed: %v", err.Error())
		http.Error(w, "Login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func (a *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDStr := vars["id"]

	accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		a.logger.Warnf("Invalid account ID: %v", accountID)
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Amount int64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.logger.Warnf("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Amount > 0 {
		if err := a.accountService.Deposit(uint(accountID), req.Amount); err != nil {
			a.logger.Warnf("Failed to update balance: %v", err.Error())
			http.Error(w, "Failed to update balance: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if req.Amount < 0 {
		if err := a.accountService.Withdraw(uint(accountID), -req.Amount); err != nil {
			a.logger.Warnf("Failed to update balance: %v", err.Error())
			http.Error(w, "Failed to update balance: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		a.logger.Warnf("Invalid amount")
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":        "Balance updated successfully",
		"balance_change": strconv.FormatInt(req.Amount, 10),
	})
}
