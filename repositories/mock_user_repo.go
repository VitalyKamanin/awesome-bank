package repositories

import (
	"awesome-bank/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (u *MockUserRepository) Create(user *models.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := u.Called(email)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (u *MockUserRepository) FindByID(id uint) (*models.User, error) {
	args := u.Called(id)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}
