package handlers

import (
	"awesome-bank/services"
	"awesome-bank/services/utils"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CardHandler struct {
	cardService *services.CardService
	logger      *logrus.Logger
}

func NewCardHandler(cardService *services.CardService, logger *logrus.Logger) *CardHandler {
	return &CardHandler{
		cardService: cardService,
		logger:      logger,
	}
}

func (c *CardHandler) IssueCard(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	accountIDStr := r.URL.Query().Get("account_id")

	if accountIDStr == "" {
		c.logger.Warnf("Missing account ID")
		http.Error(w, "Missing account ID", http.StatusBadRequest)
		return
	}

	accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		c.logger.Warnf("Invalid account ID: %v", accountID)
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	userIDUint, err := utils.ParseUserID(userIDStr)
	if err != nil {
		c.logger.Warnf("Invalid user ID: %v", userIDUint)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	card, err := c.cardService.IssueCard(userIDUint, uint(accountID))
	if err != nil {
		c.logger.Warnf("Failed to issue card: %v", err)
		http.Error(w, fmt.Sprintf("Failed to issue card: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

func (c *CardHandler) GetCards(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("userID")
	if userIDInterface == nil {
		c.logger.Warnf("Unauthorized: %v", userIDInterface)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		c.logger.Warnf("Invalid user ID type: %v", userID)
		http.Error(w, "Invalid user ID type", http.StatusInternalServerError)
		return
	}

	cards, err := c.cardService.GetCards(userID)

	if err != nil {
		c.logger.Warnf("Login failed: " + err.Error())
		http.Error(w, "Login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}
