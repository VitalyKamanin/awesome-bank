package services

import (
	"awesome-bank/integrations"
	"awesome-bank/models"
	"awesome-bank/repositories"
	"math"
	"time"
)

type CreditService struct {
	paymentRepo repositories.PaymentScheduleRepository
	accountRepo repositories.AccountRepository
	creditRepo  repositories.CreditRepository
}

func NewCreditService(paymentRepo repositories.PaymentScheduleRepository, accountRepo repositories.AccountRepository, creditRepo repositories.CreditRepository) *CreditService {
	return &CreditService{
		paymentRepo: paymentRepo,
		accountRepo: accountRepo,
		creditRepo:  creditRepo,
	}
}

func CalculateAnnuityPayment(amount int64, annualRate float64, months int) int64 {
	if months == 0 {
		return 0
	}
	r := annualRate / 12 / 100
	numerator := r * math.Pow(1+r, float64(months))
	denominator := math.Pow(1+r, float64(months)) - 1
	monthlyPayment := float64(amount) * numerator / denominator
	return int64(math.Round(monthlyPayment))
}

func (s *CreditService) GeneratePaymentSchedule(credit *models.Credit) ([]models.PaymentSchedule, error) {
	var schedule []models.PaymentSchedule
	months := credit.DurationDays / 30
	if months == 0 {
		months = 1
	}

	monthlyPayment := CalculateAnnuityPayment(credit.Amount, credit.InterestRate, months)
	startDate := time.Now()

	for i := 0; i < months; i++ {
		date := startDate.AddDate(0, i, 0)
		schedule = append(schedule, models.PaymentSchedule{
			CreditID:   credit.ID,
			AmountDue:  monthlyPayment,
			DueDate:    date.Format(time.RFC3339),
			PaidAmount: 0,
			Status:     "pending",
		})
	}
	return schedule, nil
}

func (s *CreditService) IssueCredit(userID uint, accountID uint, amount int64, interestRate float64, durationDays int) (*models.Credit, error) {
	credit := &models.Credit{
		UserID:         userID,
		AccountID:      accountID,
		Amount:         amount,
		InterestRate:   interestRate,
		DurationDays:   durationDays,
		MonthlyPayment: CalculateAnnuityPayment(amount, interestRate, durationDays/30),
		Status:         "active",
		IssuedAt:       time.Now().Format(time.RFC3339),
		DueDate:        time.Now().AddDate(0, durationDays/30, 0).Format(time.RFC3339),
	}

	s.creditRepo.Create(credit)

	schedule, err := s.GeneratePaymentSchedule(credit)
	if err != nil {
		return nil, err
	}
	for _, p := range schedule {
		s.paymentRepo.Create(&p)
	}

	account, _ := s.accountRepo.GetByID(accountID)
	account.Balance += amount
	s.accountRepo.Update(account)

	return credit, nil
}

func (s *CreditService) GetKeyRateWithMargin() (float64, error) {
	rate, err := integrations.GetCentralBankRate()
	if err != nil {
		return 0, err
	}
	return rate, nil
}
