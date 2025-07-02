package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"awesome-bank/services"
)

type AnalyticsHandler struct {
	analyticsService *services.AnalyticsService
	logger           *logrus.Logger
}

func NewAnalyticsHandler(analyticsService *services.AnalyticsService, logger *logrus.Logger) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
		logger:           logger,
	}
}

func (a *AnalyticsHandler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	daysStr := r.URL.Query().Get("days")
	days := 30
	if daysStr != "" {
		var err error
		days, err = strconv.Atoi(daysStr)
		if err != nil || days < 1 || days > 365 {
			a.logger.Warnf("Invalid days parameter: %v", days)
			http.Error(w, "Invalid days parameter", http.StatusBadRequest)
			return
		}
	}

	stats, _ := a.analyticsService.GetMonthlyStats(userID)
	creditLoad, _ := a.analyticsService.GetCreditLoad(userID)
	prediction, _ := a.analyticsService.PredictBalance(userID, days)

	response := map[string]interface{}{
		"monthly_stats":      stats,
		"credit_load":        creditLoad,
		"balance_prediction": prediction,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
