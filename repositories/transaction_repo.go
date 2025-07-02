package repositories

import (
	"awesome-bank/configs"
	"awesome-bank/models"
	"gorm.io/gorm"
	"time"
)

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (t *transactionRepository) Create(transaction *models.Transaction) error {
	result := t.DB.Create(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *transactionRepository) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := t.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *transactionRepository) GetAllByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := t.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) GetAllBySenderOrReceiverAfterDate(userIDUint uint, date time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := configs.DB.Where("sender_user_id = ? OR receiver_user_id = ?", userIDUint, userIDUint).
		Where("created_at >= ?", date.Format(time.RFC3339)).
		Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
