package repositories

import (
	"awesome-bank/models"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (t *MockTransactionRepository) Create(txn *models.Transaction) error {
	args := t.Called(txn)
	return args.Error(0)
}

func (t *MockTransactionRepository) FindByID(id uint) (*models.Transaction, error) {
	args := t.Called(id)
	if transaction, ok := args.Get(0).(*models.Transaction); ok {
		return transaction, args.Error(1)
	}
	return nil, args.Error(1)
}

func (t *MockTransactionRepository) GetAllByUserID(userID uint) ([]models.Transaction, error) {
	args := t.Called(userID)
	if transactions, ok := args.Get(0).([]models.Transaction); ok {
		return transactions, args.Error(1)
	}
	return nil, args.Error(1)
}

func (t *MockTransactionRepository) GetAllBySenderOrReceiverAfterDate(userID uint, date time.Time) ([]models.Transaction, error) {
	args := t.Called(userID, date)
	if transactions, ok := args.Get(0).([]models.Transaction); ok {
		return transactions, args.Error(1)
	}
	return nil, args.Error(1)
}
