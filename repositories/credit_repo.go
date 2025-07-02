package repositories

import (
	"awesome-bank/configs"
	"awesome-bank/models"
	"gorm.io/gorm"
)

type creditRepository struct {
	DB *gorm.DB
}

func NewCreditRepository(db *gorm.DB) CreditRepository {
	return &creditRepository{DB: db}
}

func (c *creditRepository) Create(credit *models.Credit) error {
	if err := c.DB.Create(credit).Error; err != nil {
		return err
	}
	return nil
}

func (c *creditRepository) GetCredits(userIDUint uint) ([]models.Credit, error) {
	var credits []models.Credit
	if err := configs.DB.Where("user_id = ?", userIDUint).Find(&credits).Error; err != nil {
		return nil, err
	}
	return credits, nil
}
