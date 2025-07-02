package services

import (
	"awesome-bank/repositories"
	"time"
)

type Scheduler struct {
	paymentRepo repositories.PaymentScheduleRepository
	accountRepo repositories.AccountRepository
}

func (s *Scheduler) Start() {
	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		defer ticker.Stop()

		for {
			s.processPendingPayments()
			<-ticker.C
		}
	}()
}

func (s *Scheduler) processPendingPayments() {
	payments, _ := s.paymentRepo.GetPendingPayments()
	for _, p := range payments {
		account, _ := s.accountRepo.GetByID(p.ID)
		if account.Balance >= p.AmountDue {
			account.Balance -= p.AmountDue
			s.accountRepo.Update(account)
			p.PaidAmount += p.AmountDue
			p.DueDate = time.Now().AddDate(0, 1, 0).Format(time.RFC3339)
			p.Status = "paid"
			s.paymentRepo.Update(&p)
		} else {
			p.AmountDue += int64(float64(p.AmountDue) * 0.1)
			s.paymentRepo.Update(&p)
		}
	}
}
