package repository

import (
	"cobra-script-center/internal/database"
	"cobra-script-center/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SimpleScriptRepository struct {
	store *database.JSONStore
}

func NewSimpleScriptRepository(store *database.JSONStore) *SimpleScriptRepository {
	return &SimpleScriptRepository{store: store}
}

func (r *SimpleScriptRepository) Create(script *models.Script) (*models.Script, error) {
	script.ID = uuid.New().String()
	script.CreatedAt = time.Now()
	script.UpdatedAt = time.Now()
	script.Version = 1
	script.IsActive = true

	return script, r.store.SaveScript(script)
}

func (r *SimpleScriptRepository) GetByID(id string) (*models.Script, error) {
	return r.store.GetScript(id)
}

func (r *SimpleScriptRepository) GetByName(name string) (*models.Script, error) {
	return r.store.GetScriptByName(name)
}

func (r *SimpleScriptRepository) List(filter *models.ScriptFilter) ([]*models.Script, error) {
	scripts, err := r.store.ListScripts()
	if err != nil {
		return nil, err
	}

	// Apply filters
	if filter != nil {
		var filtered []*models.Script
		for _, script := range scripts {
			if r.matchesFilter(script, filter) {
				filtered = append(filtered, script)
			}
		}
		scripts = filtered
	}

	return scripts, nil
}

func (r *SimpleScriptRepository) matchesFilter(script *models.Script, filter *models.ScriptFilter) bool {
	if filter.Name != "" && !strings.Contains(strings.ToLower(script.Name), strings.ToLower(filter.Name)) {
		return false
	}
	if filter.Language != "" && script.Language != filter.Language {
		return false
	}
	if filter.CreatedBy != "" && script.CreatedBy != filter.CreatedBy {
		return false
	}
	if filter.IsActive != nil && script.IsActive != *filter.IsActive {
		return false
	}
	return true
}

func (r *SimpleScriptRepository) Update(script *models.Script) error {
	script.UpdatedAt = time.Now()
	script.Version++
	return r.store.SaveScript(script)
}

func (r *SimpleScriptRepository) Delete(id string) error {
	return r.store.DeleteScript(id)
}

func (r *SimpleScriptRepository) DeleteByName(name string) error {
	script, err := r.store.GetScriptByName(name)
	if err != nil {
		return fmt.Errorf("script not found")
	}
	return r.store.DeleteScript(script.ID)
}

func (r *SimpleScriptRepository) Search(query string) ([]*models.Script, error) {
	scripts, err := r.store.ListScripts()
	if err != nil {
		return nil, err
	}

	var results []*models.Script
	queryLower := strings.ToLower(query)

	for _, script := range scripts {
		if script.IsActive && (strings.Contains(strings.ToLower(script.Name), queryLower) ||
			strings.Contains(strings.ToLower(script.Description), queryLower) ||
			strings.Contains(strings.ToLower(script.Content), queryLower)) {
			results = append(results, script)
		}
	}

	return results, nil
}
