package repository

import (
	"cobra-script-center/internal/database"
	"cobra-script-center/internal/models"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SimpleUserRepository struct {
	store *database.JSONStore
}

func NewSimpleUserRepository(store *database.JSONStore) *SimpleUserRepository {
	return &SimpleUserRepository{store: store}
}

func (r *SimpleUserRepository) Create(user *models.User) (*models.User, error) {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true

	return user, r.store.SaveUser(user)
}

func (r *SimpleUserRepository) GetByID(id string) (*models.User, error) {
	return r.store.GetUser(id)
}

func (r *SimpleUserRepository) GetByUsername(username string) (*models.User, error) {
	return r.store.GetUserByUsername(username)
}

func (r *SimpleUserRepository) List() ([]*models.User, error) {
	return r.store.ListUsers()
}

func (r *SimpleUserRepository) Update(user *models.User) error {
	user.UpdatedAt = time.Now()
	return r.store.SaveUser(user)
}

func (r *SimpleUserRepository) Delete(id string) error {
	return r.store.DeleteUser(id)
}

func (r *SimpleUserRepository) DeleteByUsername(username string) error {
	user, err := r.store.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	return r.store.DeleteUser(user.ID)
}
