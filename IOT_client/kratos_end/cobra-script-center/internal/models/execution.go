package models

import (
	"encoding/json"
	"time"
)

// Execution represents a script execution record
type Execution struct {
	ID         string            `json:"id" db:"id"`
	ScriptID   string            `json:"script_id" db:"script_id"`
	UserID     string            `json:"user_id" db:"user_id"`
	Status     string            `json:"status" db:"status"`
	StartTime  *time.Time        `json:"start_time" db:"start_time"`
	EndTime    *time.Time        `json:"end_time" db:"end_time"`
	Output     string            `json:"output" db:"output"`
	Error      string            `json:"error" db:"error"`
	Params     map[string]string `json:"params" db:"-"`
	ParamsJSON string            `json:"-" db:"params"`
	CreatedAt  time.Time         `json:"created_at" db:"created_at"`
}

// ExecutionStatus defines execution statuses
type ExecutionStatus string

const (
	StatusPending   ExecutionStatus = "pending"
	StatusRunning   ExecutionStatus = "running"
	StatusSuccess   ExecutionStatus = "success"
	StatusFailed    ExecutionStatus = "failed"
	StatusCancelled ExecutionStatus = "cancelled"
	StatusTimeout   ExecutionStatus = "timeout"
)

// IsValidStatus checks if the status is valid
func IsValidStatus(status string) bool {
	switch ExecutionStatus(status) {
	case StatusPending, StatusRunning, StatusSuccess, StatusFailed, StatusCancelled, StatusTimeout:
		return true
	default:
		return false
	}
}

// MarshalParams converts params map to JSON string for database storage
func (e *Execution) MarshalParams() error {
	if e.Params == nil {
		e.ParamsJSON = "{}"
		return nil
	}

	data, err := json.Marshal(e.Params)
	if err != nil {
		return err
	}
	e.ParamsJSON = string(data)
	return nil
}

// UnmarshalParams converts JSON string from database to params map
func (e *Execution) UnmarshalParams() error {
	if e.ParamsJSON == "" {
		e.Params = make(map[string]string)
		return nil
	}

	return json.Unmarshal([]byte(e.ParamsJSON), &e.Params)
}

// Duration returns the execution duration
func (e *Execution) Duration() time.Duration {
	if e.StartTime == nil || e.EndTime == nil {
		return 0
	}
	return e.EndTime.Sub(*e.StartTime)
}

// IsCompleted checks if the execution is completed
func (e *Execution) IsCompleted() bool {
	status := ExecutionStatus(e.Status)
	return status == StatusSuccess || status == StatusFailed || status == StatusCancelled || status == StatusTimeout
}

// ExecutionResult represents the result of a script execution
type ExecutionResult struct {
	ID        string            `json:"id"`
	Status    string            `json:"status"`
	Output    string            `json:"output"`
	Error     string            `json:"error"`
	Duration  time.Duration     `json:"duration"`
	Params    map[string]string `json:"params"`
	StartTime *time.Time        `json:"start_time"`
	EndTime   *time.Time        `json:"end_time"`
}
