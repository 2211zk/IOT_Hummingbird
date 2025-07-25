package models

import (
	"encoding/json"
	"time"
)

// Script represents a script in the system
type Script struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Content     string    `json:"content" db:"content"`
	Language    string    `json:"language" db:"language"`
	Tags        []string  `json:"tags" db:"-"`
	TagsJSON    string    `json:"-" db:"tags"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Version     int       `json:"version" db:"version"`
	IsActive    bool      `json:"is_active" db:"is_active"`
}

// ScriptLanguage defines supported script languages
type ScriptLanguage string

const (
	LanguageBash       ScriptLanguage = "bash"
	LanguagePython     ScriptLanguage = "python"
	LanguageNode       ScriptLanguage = "node"
	LanguageGo         ScriptLanguage = "go"
	LanguagePowershell ScriptLanguage = "powershell"
)

// IsValidLanguage checks if the language is supported
func IsValidLanguage(lang string) bool {
	switch ScriptLanguage(lang) {
	case LanguageBash, LanguagePython, LanguageNode, LanguageGo, LanguagePowershell:
		return true
	default:
		return false
	}
}

// MarshalTags converts tags slice to JSON string for database storage
func (s *Script) MarshalTags() error {
	if s.Tags == nil {
		s.TagsJSON = "[]"
		return nil
	}

	data, err := json.Marshal(s.Tags)
	if err != nil {
		return err
	}
	s.TagsJSON = string(data)
	return nil
}

// UnmarshalTags converts JSON string from database to tags slice
func (s *Script) UnmarshalTags() error {
	if s.TagsJSON == "" {
		s.Tags = []string{}
		return nil
	}

	return json.Unmarshal([]byte(s.TagsJSON), &s.Tags)
}

// ScriptFilter represents filters for script queries
type ScriptFilter struct {
	Name      string   `json:"name"`
	Language  string   `json:"language"`
	Tags      []string `json:"tags"`
	CreatedBy string   `json:"created_by"`
	IsActive  *bool    `json:"is_active"`
	Limit     int      `json:"limit"`
	Offset    int      `json:"offset"`
}
