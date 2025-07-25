package app

import (
	"fmt"

	"cobra-script-center/internal/config"
	"cobra-script-center/internal/database"
	"cobra-script-center/internal/logger"
)

// App represents the application
type App struct {
	Config *config.Config
	DB     *database.DB
}

// New creates a new application instance
func New() (*App, error) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize logger
	if err := logger.Init(&cfg.Logging); err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	log := logger.GetLogger()
	log.Info("Starting Script Center application")

	// Initialize database
	db, err := database.NewConnection(&cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	if err := db.Migrate(); err != nil {
		log.Warnf("Migration failed (this might be normal on first run): %v", err)
	}

	log.Info("Application initialized successfully")

	return &App{
		Config: cfg,
		DB:     db,
	}, nil
}

// Close closes the application and cleans up resources
func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}
