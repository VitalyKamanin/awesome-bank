package services

import (
	"awesome-bank/repositories"
	"awesome-bank/services/utils"
	"errors"
	"fmt"
	"time"

	"awesome-bank/models"
)

type AccountService struct {
	accountRepo     repositories.AccountRepository
	transactionRepo repositories.TransactionRepository
}

func NewAccountService(accountRepo repositories.AccountRepository, transactionRepo repositories.TransactionRepository) *AccountService {
	return &AccountService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *AccountService) CreateAccount(userID string, accountType string) error {
	userIDUint, err := utils.ParseUserID(userID)
	if err != nil {
		return err
	}

	account := &models.Account{
		UserID:  userIDUint,
		Balance: 0,
		Type:    accountType,
	}

	return s.accountRepo.Create(account)
}

func (s *AccountService) GetAccounts(userID string) ([]models.Account, error) {
	userIDUint, err := utils.ParseUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	return s.accountRepo.GetAccounts(userIDUint)
}

func (s *AccountService) Deposit(accountID uint, amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return err
	}

	account.Balance += amount
	s.accountRepo.Update(account)

	txn := createTransaction(account, nil, amount, "deposit", "Deposit successful")
	if err := s.transactionRepo.Create(txn); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Withdraw(accountID uint, amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient funds")
	}

	account.Balance -= amount
	s.accountRepo.Update(account)

	txn := createTransaction(nil, account, amount, "withdraw", "Withdraw successful")
	if err := s.transactionRepo.Create(txn); err != nil {
		return err
	}

	return nil
}

func createTransaction(fromAccount, toAccount *models.Account, amount int64, transType, description string) *models.Transaction {
	transaction := &models.Transaction{
		Amount:      amount,
		Type:        transType,
		Description: description,
		Status:      "success",
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	if fromAccount != nil {
		transaction.FromAccountID = fromAccount.ID
		transaction.SenderUserID = fromAccount.UserID
	}

	if toAccount != nil {
		transaction.ToAccountID = toAccount.ID
		transaction.ReceiverUserID = toAccount.UserID
	}

	return transaction
}
