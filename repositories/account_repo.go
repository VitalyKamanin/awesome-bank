package repositories

import (
	"awesome-bank/models"
	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{DB: db}
}

func (a *accountRepository) Create(account *models.Account) error {
	if err := a.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (a *accountRepository) GetByID(id uint) (*models.Account, error) {
	var account models.Account
	if err := a.DB.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *accountRepository) Update(account *models.Account) error {
	if err := a.DB.Save(account).Error; err != nil {
		return err
	}
	return nil
}

func (a *accountRepository) GetAccounts(userIDUint uint) ([]models.Account, error) {
	var accounts []models.Account
	if err := a.DB.Where("user_id = ?", userIDUint).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
