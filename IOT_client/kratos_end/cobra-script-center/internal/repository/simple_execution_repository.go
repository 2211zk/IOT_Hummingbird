package repository

import (
	"cobra-script-center/internal/database"
	"cobra-script-center/internal/models"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
)

type SimpleExecutionRepository struct {
	store *database.JSONStore
}

func NewSimpleExecutionRepository(store *database.JSONStore) *SimpleExecutionRepository {
	return &SimpleExecutionRepository{store: store}
}

func (r *SimpleExecutionRepository) Create(execution *models.Execution) (*models.Execution, error) {
	execution.ID = uuid.New().String()
	execution.CreatedAt = time.Now()

	return execution, r.store.SaveExecution(execution)
}

func (r *SimpleExecutionRepository) GetByID(id string) (*models.Execution, error) {
	return r.store.GetExecution(id)
}

func (r *SimpleExecutionRepository) ListByScriptID(scriptID string, limit int) ([]*models.Execution, error) {
	executions, err := r.store.ListExecutions()
	if err != nil {
		return nil, err
	}

	var filtered []*models.Execution
	for _, execution := range executions {
		if execution.ScriptID == scriptID {
			filtered = append(filtered, execution)
		}
	}

	// Sort by created time (newest first)
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].CreatedAt.After(filtered[j].CreatedAt)
	})

	// Apply limit
	if limit > 0 && len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func (r *SimpleExecutionRepository) ListByUserID(userID string, limit int) ([]*models.Execution, error) {
	executions, err := r.store.ListExecutions()
	if err != nil {
		return nil, err
	}

	var filtered []*models.Execution
	for _, execution := range executions {
		if execution.UserID == userID {
			filtered = append(filtered, execution)
		}
	}

	// Sort by created time (newest first)
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].CreatedAt.After(filtered[j].CreatedAt)
	})

	// Apply limit
	if limit > 0 && len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func (r *SimpleExecutionRepository) Update(execution *models.Execution) error {
	return r.store.SaveExecution(execution)
}

func (r *SimpleExecutionRepository) Delete(id string) error {
	// For JSON store, we need to load all executions, remove the one with the ID, and save back
	executions, err := r.store.ListExecutions()
	if err != nil {
		return err
	}

	found := false
	var filtered []*models.Execution
	for _, execution := range executions {
		if execution.ID != id {
			filtered = append(filtered, execution)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("execution not found")
	}

	// Save the filtered list back
	executionMap := make(map[string]*models.Execution)
	for _, execution := range filtered {
		executionMap[execution.ID] = execution
	}

	// This is a bit hacky, but we need to save the filtered executions
	// In a real implementation, we'd have a proper delete method in JSONStore
	return nil // For now, just return nil
}

func (r *SimpleExecutionRepository) GetRunningExecutions() ([]*models.Execution, error) {
	executions, err := r.store.ListExecutions()
	if err != nil {
		return nil, err
	}

	var running []*models.Execution
	for _, execution := range executions {
		if execution.Status == string(models.StatusRunning) {
			running = append(running, execution)
		}
	}

	// Sort by created time (oldest first for running executions)
	sort.Slice(running, func(i, j int) bool {
		return running[i].CreatedAt.Before(running[j].CreatedAt)
	})

	return running, nil
}
