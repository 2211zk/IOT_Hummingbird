package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID           string    `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	IsActive     bool      `json:"is_active" db:"is_active"`
}

// UserRole defines user roles
type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleUser   UserRole = "user"
	RoleViewer UserRole = "viewer"
)

// IsValidRole checks if the role is valid
func IsValidRole(role string) bool {
	switch UserRole(role) {
	case RoleAdmin, RoleUser, RoleViewer:
		return true
	default:
		return false
	}
}

// HasPermission checks if user has permission for an action
func (u *User) HasPermission(action string) bool {
	switch u.Role {
	case string(RoleAdmin):
		return true // Admin has all permissions
	case string(RoleUser):
		// Users can perform most actions except user management
		return action != "user_create" && action != "user_delete" && action != "user_update"
	case string(RoleViewer):
		// Viewers can only read
		return action == "script_read" || action == "execution_read"
	default:
		return false
	}
}
