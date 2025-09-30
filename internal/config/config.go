package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our application
type Config struct {
	Server  ServerConfig
	Blog    BlogConfig
	Admin   AdminConfig
	Logging LoggingConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Host string
	Port string
}

// BlogConfig holds blog-specific configuration
type BlogConfig struct {
	Title       string
	Description string
	Author      string
	URL         string
}

// AdminConfig holds admin panel configuration
type AdminConfig struct {
	Username string
	Password string
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level string
}

// Load loads configuration from environment variables
// It tries to load from multiple .env file locations
func Load() *Config {
	// Try to load .env from multiple locations
	envPaths := []string{
		".env",                    // Root directory
		"configs/.env",            // Configs directory
		"../configs/.env",         // One level up (if running from cmd/)
		"../../configs/.env",      // Two levels up (if running from cmd/server/)
	}

	loaded := false
	for _, path := range envPaths {
		if absPath, err := filepath.Abs(path); err == nil {
			if err := godotenv.Load(absPath); err == nil {
				log.Printf("Loaded configuration from: %s", absPath)
				loaded = true
				break
			}
		}
	}

	if !loaded {
		log.Println("No .env file found, using environment variables and defaults")
	}

	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "3100"),
		},
		Blog: BlogConfig{
			Title:       getEnv("BLOG_TITLE", "My Blog"),
			Description: getEnv("BLOG_DESCRIPTION", "A personal blog about programming, technology, and life"),
			Author:      getEnv("BLOG_AUTHOR", "Blog Author"),
			URL:         getEnv("BLOG_URL", "http://localhost:3100"),
		},
		Admin: AdminConfig{
			Username: getEnv("ADMIN_USERNAME", "admin"),
			Password: getEnv("ADMIN_PASSWORD", "admin123"),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// Check critical settings
	if c.Admin.Username == "admin" && c.Admin.Password == "admin123" {
		log.Println("WARNING: Using default admin credentials! Please change them in production.")
	}

	return nil
}

// Helper functions for environment variable parsing

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
