package service

import (
	"cobra-script-center/internal/models"
	"cobra-script-center/internal/repository"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type ExecutionService struct {
	executionRepo repository.ExecutionRepositoryInterface
	scriptRepo    repository.ScriptRepositoryInterface
}

func NewExecutionService(executionRepo repository.ExecutionRepositoryInterface, scriptRepo repository.ScriptRepositoryInterface) *ExecutionService {
	return &ExecutionService{
		executionRepo: executionRepo,
		scriptRepo:    scriptRepo,
	}
}

func (s *ExecutionService) ExecuteScript(scriptNameOrID, userID string, params map[string]string) (*models.Execution, error) {
	// Get script by name or ID
	script, err := s.scriptRepo.GetByName(scriptNameOrID)
	if err != nil {
		// Try by ID if name lookup failed
		script, err = s.scriptRepo.GetByID(scriptNameOrID)
		if err != nil {
			return nil, fmt.Errorf("script not found: %s", scriptNameOrID)
		}
	}

	if !script.IsActive {
		return nil, fmt.Errorf("script is not active")
	}

	// Create execution record
	execution := &models.Execution{
		ScriptID: script.ID,
		UserID:   userID,
		Status:   string(models.StatusPending),
		Params:   params,
	}

	createdExecution, err := s.executionRepo.Create(execution)
	if err != nil {
		return nil, fmt.Errorf("failed to create execution record: %w", err)
	}

	// Execute script asynchronously
	go s.executeScriptAsync(createdExecution, script)

	return createdExecution, nil
}

func (s *ExecutionService) executeScriptAsync(execution *models.Execution, script *models.Script) {
	// Update status to running
	execution.Status = string(models.StatusRunning)
	startTime := time.Now()
	execution.StartTime = &startTime

	if err := s.executionRepo.Update(execution); err != nil {
		fmt.Printf("Failed to update execution status: %v\n", err)
		return
	}

	// Execute the script
	output, execErr := s.runScript(script, execution.Params)

	// Update execution with results
	endTime := time.Now()
	execution.EndTime = &endTime
	execution.Output = output

	if execErr != nil {
		execution.Status = string(models.StatusFailed)
		execution.Error = execErr.Error()
	} else {
		execution.Status = string(models.StatusSuccess)
	}

	if err := s.executionRepo.Update(execution); err != nil {
		fmt.Printf("Failed to update execution results: %v\n", err)
	}
}

func (s *ExecutionService) runScript(script *models.Script, params map[string]string) (string, error) {
	// Create temporary file for script
	tmpFile, err := s.createTempScript(script)
	if err != nil {
		return "", fmt.Errorf("failed to create temp script: %w", err)
	}
	defer os.Remove(tmpFile)

	// Prepare command based on language
	cmd, err := s.prepareCommand(script.Language, tmpFile, params)
	if err != nil {
		return "", fmt.Errorf("failed to prepare command: %w", err)
	}

	// Execute command
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func (s *ExecutionService) createTempScript(script *models.Script) (string, error) {
	// Create temp file with appropriate extension
	var ext string
	switch script.Language {
	case "bash":
		ext = ".sh"
	case "python":
		ext = ".py"
	case "node":
		ext = ".js"
	case "go":
		ext = ".go"
	case "powershell":
		ext = ".ps1"
	default:
		ext = ".txt"
	}

	tmpFile, err := os.CreateTemp("", "script-*"+ext)
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	// Write script content
	if _, err := tmpFile.WriteString(script.Content); err != nil {
		return "", err
	}

	// Make executable for shell scripts
	if script.Language == "bash" {
		if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
			return "", err
		}
	}

	return tmpFile.Name(), nil
}

func (s *ExecutionService) prepareCommand(language, scriptPath string, params map[string]string) (*exec.Cmd, error) {
	var cmd *exec.Cmd

	switch language {
	case "bash":
		cmd = exec.Command("bash", scriptPath)
	case "python":
		cmd = exec.Command("python3", scriptPath)
	case "node":
		cmd = exec.Command("node", scriptPath)
	case "go":
		cmd = exec.Command("go", "run", scriptPath)
	case "powershell":
		cmd = exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", scriptPath)
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}

	// Set environment variables from params
	env := os.Environ()
	for key, value := range params {
		env = append(env, fmt.Sprintf("%s=%s", strings.ToUpper(key), value))
	}
	cmd.Env = env

	return cmd, nil
}

func (s *ExecutionService) GetExecution(id string) (*models.Execution, error) {
	return s.executionRepo.GetByID(id)
}

func (s *ExecutionService) GetExecutionsByScript(scriptID string, limit int) ([]*models.Execution, error) {
	return s.executionRepo.ListByScriptID(scriptID, limit)
}

func (s *ExecutionService) GetExecutionsByUser(userID string, limit int) ([]*models.Execution, error) {
	return s.executionRepo.ListByUserID(userID, limit)
}

func (s *ExecutionService) GetRunningExecutions() ([]*models.Execution, error) {
	return s.executionRepo.GetRunningExecutions()
}

func (s *ExecutionService) CancelExecution(id string) error {
	execution, err := s.executionRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("execution not found: %w", err)
	}

	if execution.Status != string(models.StatusRunning) {
		return fmt.Errorf("execution is not running")
	}

	// Update status to cancelled
	execution.Status = string(models.StatusCancelled)
	endTime := time.Now()
	execution.EndTime = &endTime

	return s.executionRepo.Update(execution)
}

func (s *ExecutionService) DeleteExecution(id string) error {
	return s.executionRepo.Delete(id)
}
