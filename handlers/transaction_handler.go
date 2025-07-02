package handlers

import (
	"awesome-bank/services"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
	logger             *logrus.Logger
}

func NewTransactionHandler(transactionService *services.TransactionService, logger *logrus.Logger) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		logger:             logger,
	}
}

func (t *TransactionHandler) TransferFunds(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var req struct {
		FromAccountID uint   `json:"from_account_id"`
		ToAccountID   uint   `json:"to_account_id"`
		Amount        int64  `json:"amount"`
		Description   string `json:"description,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		t.logger.Warnf("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.FromAccountID == 0 || req.ToAccountID == 0 || req.Amount <= 0 {
		t.logger.Warnf("Invalid transfer parameters")
		http.Error(w, "Invalid transfer parameters", http.StatusBadRequest)
		return
	}

	err := t.transactionService.Transfer(userID, req.FromAccountID, req.ToAccountID, req.Amount, req.Description)
	if err != nil {
		t.logger.Warnf("Transfer failed: %v", err)
		http.Error(w, fmt.Sprintf("Transfer failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}
