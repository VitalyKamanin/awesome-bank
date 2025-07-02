package repositories

import (
	"awesome-bank/models"
	"gorm.io/gorm"
)

type cardRepository struct {
	DB *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{DB: db}
}

func (c *cardRepository) IssueCard(card *models.Card) (*models.Card, error) {
	if err := c.DB.Create(card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func (c *cardRepository) GetCards(userIDUint uint) ([]models.Card, error) {
	var cards []models.Card
	if err := c.DB.Where("user_id = ?", userIDUint).Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}
