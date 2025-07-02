package repositories

import (
	"awesome-bank/models"
	"github.com/stretchr/testify/mock"
)

type MockCardRepository struct {
	mock.Mock
}

func (c *MockCardRepository) IssueCard(card *models.Card) (*models.Card, error) {
	args := c.Called(card)
	if card, ok := args.Get(0).(*models.Card); ok {
		return card, args.Error(1)
	}
	return nil, args.Error(1)
}

func (c *MockCardRepository) GetCards(userID uint) ([]models.Card, error) {
	args := c.Called(userID)
	if cards, ok := args.Get(0).([]models.Card); ok {
		return cards, args.Error(1)
	}
	return nil, args.Error(1)
}
