package services

import (
	"awesome-bank/services/utils"
	"errors"
	"fmt"

	"awesome-bank/models"
	"awesome-bank/repositories"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(email, username, password string) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Username: username,
		Password: hashed,
	}

	if err := user.Validate(); err != nil {
		utils.Logger.Warnf("Validation failed: %v", err)
		return err
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		utils.Logger.WithFields(map[string]interface{}{
			"email": email,
		}).Warn("User not found")
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		utils.Logger.WithFields(map[string]interface{}{
			"email": user.Email,
		}).Warn("Invalid password")
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(fmt.Sprintf("%d", user.ID))
	if err != nil {
		utils.Logger.WithFields(map[string]interface{}{
			"userID": user.ID,
		}).Error("Failed to generate JWT token")
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	utils.Logger.WithFields(map[string]interface{}{
		"userID": user.ID,
		"email":  user.Email,
	}).Info("User logged in successfully")

	return token, nil
}
