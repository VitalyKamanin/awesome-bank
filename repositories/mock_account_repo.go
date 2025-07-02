package repositories

import (
	"awesome-bank/models"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (a *MockAccountRepository) Create(account *models.Account) error {
	args := a.Called(account)
	return args.Error(0)
}

func (a *MockAccountRepository) GetByID(id uint) (*models.Account, error) {
	args := a.Called(id)
	if account, ok := args.Get(0).(*models.Account); ok {
		return account, args.Error(1)
	}
	return nil, args.Error(1)
}

func (a *MockAccountRepository) Update(account *models.Account) error {
	args := a.Called(account)
	return args.Error(0)
}

func (a *MockAccountRepository) GetAccounts(userID uint) ([]models.Account, error) {
	args := a.Called(userID)
	if accounts, ok := args.Get(0).([]models.Account); ok {
		return accounts, args.Error(1)
	}
	return nil, args.Error(1)
}
