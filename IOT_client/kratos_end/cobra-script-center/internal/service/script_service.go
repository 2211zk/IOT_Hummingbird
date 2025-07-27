package service

import (
	"cobra-script-center/internal/models"
	"cobra-script-center/internal/repository"
	"fmt"
)

type ScriptService struct {
	repo repository.ScriptRepositoryInterface
}

func NewScriptService(repo repository.ScriptRepositoryInterface) *ScriptService {
	return &ScriptService{
		repo: repo,
	}
}

func (s *ScriptService) CreateScript(script *models.Script, userID string) (*models.Script, error) {
	// Validate language
	if !models.IsValidLanguage(script.Language) {
		return nil, fmt.Errorf("unsupported language: %s", script.Language)
	}

	// Check if script name already exists
	existingScript, err := s.repo.GetByName(script.Name)
	if err == nil && existingScript != nil {
		return nil, fmt.Errorf("script with name '%s' already exists", script.Name)
	}

	script.CreatedBy = userID
	return s.repo.Create(script)
}

func (s *ScriptService) GetScript(id string) (*models.Script, error) {
	return s.repo.GetByID(id)
}

func (s *ScriptService) GetScriptByName(name string) (*models.Script, error) {
	return s.repo.GetByName(name)
}

func (s *ScriptService) ListScripts(filter *models.ScriptFilter) ([]*models.Script, error) {
	return s.repo.List(filter)
}

func (s *ScriptService) UpdateScript(script *models.Script) error {
	// Validate language
	if !models.IsValidLanguage(script.Language) {
		return fmt.Errorf("unsupported language: %s", script.Language)
	}

	return s.repo.Update(script)
}

func (s *ScriptService) DeleteScript(nameOrID string) error {
	// Try to get by name first, then by ID
	script, err := s.repo.GetByName(nameOrID)
	if err != nil {
		// If not found by name, try by ID
		script, err = s.repo.GetByID(nameOrID)
		if err != nil {
			return fmt.Errorf("script not found: %s", nameOrID)
		}
	}

	return s.repo.Delete(script.ID)
}

func (s *ScriptService) SearchScripts(query string) ([]*models.Script, error) {
	return s.repo.Search(query)
}

func (s *ScriptService) ValidateScript(script *models.Script) error {
	if script.Name == "" {
		return fmt.Errorf("script name is required")
	}

	if script.Content == "" {
		return fmt.Errorf("script content is required")
	}

	if !models.IsValidLanguage(script.Language) {
		return fmt.Errorf("unsupported language: %s", script.Language)
	}

	return nil
}

func (s *ScriptService) GetScriptsByUser(userID string) ([]*models.Script, error) {
	filter := &models.ScriptFilter{
		CreatedBy: userID,
	}
	return s.repo.List(filter)
}

func (s *ScriptService) GetScriptsByLanguage(language string) ([]*models.Script, error) {
	if !models.IsValidLanguage(language) {
		return nil, fmt.Errorf("unsupported language: %s", language)
	}

	filter := &models.ScriptFilter{
		Language: language,
	}
	return s.repo.List(filter)
}

func (s *ScriptService) GetScriptsByTags(tags []string) ([]*models.Script, error) {
	filter := &models.ScriptFilter{
		Tags: tags,
	}
	return s.repo.List(filter)
}
