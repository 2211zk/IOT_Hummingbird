package repository

import "cobra-script-center/internal/models"

// UserRepositoryInterface defines the interface for user repository
type UserRepositoryInterface interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	List() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	DeleteByUsername(username string) error
}

// ScriptRepositoryInterface defines the interface for script repository
type ScriptRepositoryInterface interface {
	Create(script *models.Script) (*models.Script, error)
	GetByID(id string) (*models.Script, error)
	GetByName(name string) (*models.Script, error)
	List(filter *models.ScriptFilter) ([]*models.Script, error)
	Update(script *models.Script) error
	Delete(id string) error
	DeleteByName(name string) error
	Search(query string) ([]*models.Script, error)
}

// ExecutionRepositoryInterface defines the interface for execution repository
type ExecutionRepositoryInterface interface {
	Create(execution *models.Execution) (*models.Execution, error)
	GetByID(id string) (*models.Execution, error)
	ListByScriptID(scriptID string, limit int) ([]*models.Execution, error)
	ListByUserID(userID string, limit int) ([]*models.Execution, error)
	Update(execution *models.Execution) error
	Delete(id string) error
	GetRunningExecutions() ([]*models.Execution, error)
}
