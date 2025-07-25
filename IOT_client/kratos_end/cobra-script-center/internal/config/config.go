package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	Security SecurityConfig `mapstructure:"security"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// SecurityConfig holds security configuration
type SecurityConfig struct {
	JWTSecret     string `mapstructure:"jwt_secret"`
	PasswordSalt  string `mapstructure:"password_salt"`
	MaxExecutions int    `mapstructure:"max_executions"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	File   string `mapstructure:"file"`
}

// Load loads configuration from viper
func Load() (*Config, error) {
	// Set default values
	setDefaults()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaults() {
	// Database defaults
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.dsn", "./script-center.db")

	// Server defaults
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)

	// Security defaults
	viper.SetDefault("security.jwt_secret", "your-secret-key-change-in-production")
	viper.SetDefault("security.password_salt", "your-salt-change-in-production")
	viper.SetDefault("security.max_executions", 10)

	// Logging defaults
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.format", "json")
	viper.SetDefault("logging.file", "")
}
