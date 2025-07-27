package app

import (
	"fmt"

	"cobra-script-center/internal/config"
	"cobra-script-center/internal/database"
	"cobra-script-center/internal/logger"
	"cobra-script-center/internal/repository"
	"cobra-script-center/internal/service"
)

// App represents the application
type App struct {
	Config           *config.Config
	DB               *database.SimpleDB
	UserService      *service.UserService
	ScriptService    *service.ScriptService
	ExecutionService *service.ExecutionService
}

// NewApp creates a new application instance
func NewApp() (*App, error) {
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

	// Initialize simple database
	db, err := database.NewSimpleConnection(&cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations (create data directory)
	if err := db.Migrate(); err != nil {
		log.Warnf("Migration failed: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewSimpleUserRepository(db.GetStore())
	scriptRepo := repository.NewSimpleScriptRepository(db.GetStore())
	executionRepo := repository.NewSimpleExecutionRepository(db.GetStore())

	// Initialize services
	userService := service.NewUserService(userRepo, &cfg.Security)
	scriptService := service.NewScriptService(scriptRepo)
	executionService := service.NewExecutionService(executionRepo, scriptRepo)

	log.Info("Application initialized successfully")

	return &App{
		Config:           cfg,
		DB:               db,
		UserService:      userService,
		ScriptService:    scriptService,
		ExecutionService: executionService,
	}, nil
}

// Close closes the application and cleans up resources
func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}
