package models

type MonthlyStats struct {
	Month         string               `json:"month"`
	TotalIncome   int64                `json:"total_income"`
	TotalExpenses int64                `json:"total_expenses"`
	Balance       int64                `json:"balance"`
	Transactions  []TransactionSummary `json:"transactions"`
}

type TransactionSummary struct {
	ID          uint   `json:"id"`
	Amount      int64  `json:"amount"`
	Direction   string `json:"direction"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type CreditLoad struct {
	TotalCredits      int64   `json:"total_credits"`
	OutstandingDebt   int64   `json:"outstanding_debt"`
	UpcomingPayments  int64   `json:"upcoming_payments"`
	CreditsInProgress int     `json:"credits_in_progress"`
	AverageInterest   float64 `json:"average_interest"`
	CreditUtilization float64 `json:"credit_utilization"`
}

type BalancePrediction struct {
	Days           int            `json:"days"`
	StartDate      string         `json:"start_date"`
	EndDate        string         `json:"end_date"`
	BalanceHistory []BalancePoint `json:"balance_history"`
}

type BalancePoint struct {
	AccountID uint           `json:"-"`
	History   []BalanceEntry `json:"history"`
}

type BalanceEntry struct {
	Date    string `json:"date"`
	Balance int64  `json:"balance"`
}
