package database

import (
	"cobra-script-center/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// JSONStore is a simple file-based storage using JSON
// This is a fallback when SQLite is not available
type JSONStore struct {
	dataDir string
	mutex   sync.RWMutex
}

// NewJSONStore creates a new JSON-based storage
func NewJSONStore(dataDir string) (*JSONStore, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	return &JSONStore{
		dataDir: dataDir,
	}, nil
}

// Users storage
func (s *JSONStore) SaveUser(user *models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	users, err := s.loadUsers()
	if err != nil {
		users = make(map[string]*models.User)
	}

	users[user.ID] = user
	return s.saveUsers(users)
}

func (s *JSONStore) GetUser(id string) (*models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users, err := s.loadUsers()
	if err != nil {
		return nil, err
	}

	user, exists := users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *JSONStore) GetUserByUsername(username string) (*models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users, err := s.loadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (s *JSONStore) ListUsers() ([]*models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users, err := s.loadUsers()
	if err != nil {
		return nil, err
	}

	var result []*models.User
	for _, user := range users {
		result = append(result, user)
	}

	return result, nil
}

func (s *JSONStore) DeleteUser(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	users, err := s.loadUsers()
	if err != nil {
		return err
	}

	if _, exists := users[id]; !exists {
		return fmt.Errorf("user not found")
	}

	delete(users, id)
	return s.saveUsers(users)
}

func (s *JSONStore) loadUsers() (map[string]*models.User, error) {
	filePath := filepath.Join(s.dataDir, "users.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(map[string]*models.User), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var users map[string]*models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *JSONStore) saveUsers(users map[string]*models.User) error {
	filePath := filepath.Join(s.dataDir, "users.json")

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// Scripts storage
func (s *JSONStore) SaveScript(script *models.Script) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	scripts, err := s.loadScripts()
	if err != nil {
		scripts = make(map[string]*models.Script)
	}

	scripts[script.ID] = script
	return s.saveScripts(scripts)
}

func (s *JSONStore) GetScript(id string) (*models.Script, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	scripts, err := s.loadScripts()
	if err != nil {
		return nil, err
	}

	script, exists := scripts[id]
	if !exists {
		return nil, fmt.Errorf("script not found")
	}

	return script, nil
}

func (s *JSONStore) GetScriptByName(name string) (*models.Script, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	scripts, err := s.loadScripts()
	if err != nil {
		return nil, err
	}

	for _, script := range scripts {
		if script.Name == name && script.IsActive {
			return script, nil
		}
	}

	return nil, fmt.Errorf("script not found")
}

func (s *JSONStore) ListScripts() ([]*models.Script, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	scripts, err := s.loadScripts()
	if err != nil {
		return nil, err
	}

	var result []*models.Script
	for _, script := range scripts {
		result = append(result, script)
	}

	return result, nil
}

func (s *JSONStore) DeleteScript(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	scripts, err := s.loadScripts()
	if err != nil {
		return err
	}

	if _, exists := scripts[id]; !exists {
		return fmt.Errorf("script not found")
	}

	delete(scripts, id)
	return s.saveScripts(scripts)
}

func (s *JSONStore) loadScripts() (map[string]*models.Script, error) {
	filePath := filepath.Join(s.dataDir, "scripts.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(map[string]*models.Script), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var scripts map[string]*models.Script
	if err := json.Unmarshal(data, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

func (s *JSONStore) saveScripts(scripts map[string]*models.Script) error {
	filePath := filepath.Join(s.dataDir, "scripts.json")

	data, err := json.MarshalIndent(scripts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// Executions storage
func (s *JSONStore) SaveExecution(execution *models.Execution) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	executions, err := s.loadExecutions()
	if err != nil {
		executions = make(map[string]*models.Execution)
	}

	executions[execution.ID] = execution
	return s.saveExecutions(executions)
}

func (s *JSONStore) GetExecution(id string) (*models.Execution, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	executions, err := s.loadExecutions()
	if err != nil {
		return nil, err
	}

	execution, exists := executions[id]
	if !exists {
		return nil, fmt.Errorf("execution not found")
	}

	return execution, nil
}

func (s *JSONStore) ListExecutions() ([]*models.Execution, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	executions, err := s.loadExecutions()
	if err != nil {
		return nil, err
	}

	var result []*models.Execution
	for _, execution := range executions {
		result = append(result, execution)
	}

	return result, nil
}

func (s *JSONStore) loadExecutions() (map[string]*models.Execution, error) {
	filePath := filepath.Join(s.dataDir, "executions.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(map[string]*models.Execution), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var executions map[string]*models.Execution
	if err := json.Unmarshal(data, &executions); err != nil {
		return nil, err
	}

	return executions, nil
}

func (s *JSONStore) saveExecutions(executions map[string]*models.Execution) error {
	filePath := filepath.Join(s.dataDir, "executions.json")

	data, err := json.MarshalIndent(executions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
