package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"awesome-bank/services"
)

type CbrHandler struct {
	creditService *services.CreditService
	logger        *logrus.Logger
}

func NewCbrHandler(creditService *services.CreditService, logger *logrus.Logger) *CbrHandler {
	return &CbrHandler{
		creditService: creditService,
		logger:        logger,
	}
}

func (c *CbrHandler) GetKeyRate(w http.ResponseWriter, r *http.Request) {
	rate, err := c.creditService.GetKeyRateWithMargin()
	if err != nil {
		c.logger.Warnf("Failed to get key rate: " + err.Error())
		http.Error(w, "Failed to get key rate: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"key_rate": fmt.Sprintf("%.2f%%", rate),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
