package database

import (
	"cobra-script-center/internal/config"
	"fmt"
	"os"
)

// SimpleDB is a simplified database interface using JSON files
type SimpleDB struct {
	store *JSONStore
}

// NewSimpleConnection creates a new simple database connection
func NewSimpleConnection(cfg *config.DatabaseConfig) (*SimpleDB, error) {
	// Use data directory instead of SQLite file
	dataDir := "./data"
	if cfg.DSN != "" && cfg.DSN != ":memory:" {
		// Extract directory from DSN if it's a file path
		dataDir = "./data"
	}

	store, err := NewJSONStore(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create JSON store: %w", err)
	}

	return &SimpleDB{store: store}, nil
}

// GetStore returns the underlying JSON store
func (db *SimpleDB) GetStore() *JSONStore {
	return db.store
}

// Close closes the database connection (no-op for JSON store)
func (db *SimpleDB) Close() error {
	return nil
}

// Migrate creates necessary directories (no-op for JSON store)
func (db *SimpleDB) Migrate() error {
	return nil
}

// RunSimpleMigrations runs simple migrations (creates data directory)
func RunSimpleMigrations() error {
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}
	fmt.Println("Data directory created successfully.")
	return nil
}

// RunMigrations is an alias for RunSimpleMigrations for compatibility
func RunMigrations() error {
	return RunSimpleMigrations()
}
