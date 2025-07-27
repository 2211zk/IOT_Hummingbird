package service

import (
	"cobra-script-center/internal/config"
	"cobra-script-center/internal/models"
	"cobra-script-center/internal/repository"
	"crypto/sha256"
	"fmt"
)

type UserService struct {
	repo     repository.UserRepositoryInterface
	security *config.SecurityConfig
}

func NewUserService(repo repository.UserRepositoryInterface, security *config.SecurityConfig) *UserService {
	return &UserService{
		repo:     repo,
		security: security,
	}
}

func (s *UserService) CreateUser(user *models.User, password string) (*models.User, error) {
	// Check if username already exists
	existingUser, err := s.repo.GetByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Validate role
	if !models.IsValidRole(user.Role) {
		return nil, fmt.Errorf("invalid role: %s", user.Role)
	}

	// Hash password
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user.PasswordHash = hashedPassword

	return s.repo.Create(user)
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) ListUsers() ([]*models.User, error) {
	return s.repo.List()
}

func (s *UserService) UpdateUser(user *models.User) error {
	// Validate role if it's being changed
	if !models.IsValidRole(user.Role) {
		return fmt.Errorf("invalid role: %s", user.Role)
	}

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(username string) error {
	return s.repo.DeleteByUsername(username)
}

func (s *UserService) ChangePassword(userID, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Verify old password
	if !s.verifyPassword(oldPassword, user.PasswordHash) {
		return fmt.Errorf("invalid old password")
	}

	// Hash new password
	hashedPassword, err := s.hashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	user.PasswordHash = hashedPassword
	return s.repo.Update(user)
}

func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user account is disabled")
	}

	if !s.verifyPassword(password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	// Simple SHA256 hash with salt
	// In production, use bcrypt or similar
	saltedPassword := password + s.security.PasswordSalt
	hash := sha256.Sum256([]byte(saltedPassword))
	return fmt.Sprintf("%x", hash), nil
}

func (s *UserService) verifyPassword(password, hash string) bool {
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return false
	}
	return hashedPassword == hash
}
