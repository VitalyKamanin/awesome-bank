package repositories

import (
	"awesome-bank/models"
	"time"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
}

type AccountRepository interface {
	Create(account *models.Account) error
	GetByID(id uint) (*models.Account, error)
	Update(account *models.Account) error
	GetAccounts(userID uint) ([]models.Account, error)
}

type CardRepository interface {
	IssueCard(card *models.Card) (*models.Card, error)
	GetCards(userID uint) ([]models.Card, error)
}

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindByID(id uint) (*models.Transaction, error)
	GetAllByUserID(userID uint) ([]models.Transaction, error)
	GetAllBySenderOrReceiverAfterDate(userIDUint uint, date time.Time) ([]models.Transaction, error)
}

type PaymentScheduleRepository interface {
	Create(schedule *models.PaymentSchedule) error
	GetByCreditID(creditID uint) ([]models.PaymentSchedule, error)
	GetPendingByUserID(userID uint) ([]models.PaymentSchedule, error)
	Update(schedule *models.PaymentSchedule) error
	GetPendingPayments() ([]models.PaymentSchedule, error)
	GetPendingPaymentsByUserAndDate(userIDUint uint, date time.Time) ([]models.PaymentSchedule, error)
}

type CreditRepository interface {
	Create(account *models.Credit) error
	GetCredits(userID uint) ([]models.Credit, error)
}
