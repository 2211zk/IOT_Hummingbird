package models

import (
	"time"
)

// ScheduledTask represents a scheduled script execution
type ScheduledTask struct {
	ID        string     `json:"id" db:"id"`
	ScriptID  string     `json:"script_id" db:"script_id"`
	CronExpr  string     `json:"cron_expr" db:"cron_expr"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	IsActive  bool       `json:"is_active" db:"is_active"`
	LastRun   *time.Time `json:"last_run" db:"last_run"`
	NextRun   *time.Time `json:"next_run" db:"next_run"`
}

// TaskStatus represents the status of a scheduled task
type TaskStatus struct {
	ID             string     `json:"id"`
	ScriptName     string     `json:"script_name"`
	CronExpr       string     `json:"cron_expr"`
	IsActive       bool       `json:"is_active"`
	LastRun        *time.Time `json:"last_run"`
	NextRun        *time.Time `json:"next_run"`
	LastStatus     string     `json:"last_status"`
	ExecutionCount int        `json:"execution_count"`
}

// IsValidCronExpression validates a cron expression format
// This is a basic validation - in practice you'd use a proper cron parser
func IsValidCronExpression(expr string) bool {
	// Basic validation - should have 5 or 6 parts
	// Real implementation would use github.com/robfig/cron/v3 parser
	if expr == "" {
		return false
	}

	// For now, just check it's not empty
	// TODO: Implement proper cron validation using robfig/cron
	return len(expr) > 0
}

// ShouldRun checks if the task should run now
func (t *ScheduledTask) ShouldRun() bool {
	if !t.IsActive {
		return false
	}

	if t.NextRun == nil {
		return true // First run
	}

	return time.Now().After(*t.NextRun)
}

// UpdateNextRun updates the next run time based on cron expression
func (t *ScheduledTask) UpdateNextRun() error {
	// TODO: Implement using robfig/cron parser
	// For now, just set to 1 hour from now as placeholder
	nextRun := time.Now().Add(time.Hour)
	t.NextRun = &nextRun
	return nil
}
