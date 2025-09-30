package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	Blog     BlogConfig
	Admin    AdminConfig
	Logging  LoggingConfig
	Features FeatureFlags
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxConns int
	MinConns int
}

// CacheConfig holds cache-related configuration
type CacheConfig struct {
	Type     string // "memory" or "redis"
	Host     string
	Port     string
	Password string
	DB       int
	TTL      time.Duration
}

// BlogConfig holds blog-specific configuration
type BlogConfig struct {
	Title        string
	Description  string
	Author       string
	URL          string
	PostsPerPage int
	Theme        string
}

// AdminConfig holds admin panel configuration
type AdminConfig struct {
	Username        string
	Password        string
	Secret          string
	SessionDuration time.Duration
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
	File   string
}

// FeatureFlags holds feature toggle flags
type FeatureFlags struct {
	EnableComments    bool
	EnableAnalytics   bool
	EnableNewsletter  bool
	RateLimitEnabled  bool
	RateLimitRequests int
	CORSEnabled       bool
	CORSOrigins       string
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
			Host:         getEnv("SERVER_HOST", "localhost"),
			Port:         getEnv("SERVER_PORT", "3100"),
			ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", 30*time.Second),
			WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", 30*time.Second),
			IdleTimeout:  getDurationEnv("SERVER_IDLE_TIMEOUT", 120*time.Second),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "blog_user"),
			Password: getEnv("DB_PASSWORD", "blog_password"),
			DBName:   getEnv("DB_NAME", "blog_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			MaxConns: getIntEnv("DB_MAX_CONNS", 25),
			MinConns: getIntEnv("DB_MIN_CONNS", 5),
		},
		Cache: CacheConfig{
			Type:     getEnv("CACHE_TYPE", "memory"),
			Host:     getEnv("CACHE_HOST", "localhost"),
			Port:     getEnv("CACHE_PORT", "6379"),
			Password: getEnv("CACHE_PASSWORD", ""),
			DB:       getIntEnv("CACHE_DB", 0),
			TTL:      getDurationEnv("CACHE_TTL", 1*time.Hour),
		},
		Blog: BlogConfig{
			Title:        getEnv("BLOG_TITLE", "My Blog"),
			Description:  getEnv("BLOG_DESCRIPTION", "A personal blog about programming, technology, and life"),
			Author:       getEnv("BLOG_AUTHOR", "Blog Author"),
			URL:          getEnv("BLOG_URL", "http://localhost:3100"),
			PostsPerPage: getIntEnv("BLOG_POSTS_PER_PAGE", 10),
			Theme:        getEnv("BLOG_THEME", "default"),
		},
		Admin: AdminConfig{
			Username:        getEnv("ADMIN_USERNAME", "admin"),
			Password:        getEnv("ADMIN_PASSWORD", "admin123"),
			Secret:          getEnv("ADMIN_SECRET", "change-this-secret-key-in-production"),
			SessionDuration: getDurationEnv("ADMIN_SESSION_DURATION", 24*time.Hour),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "text"),
			File:   getEnv("LOG_FILE", "server.log"),
		},
		Features: FeatureFlags{
			EnableComments:    getBoolEnv("ENABLE_COMMENTS", false),
			EnableAnalytics:   getBoolEnv("ENABLE_ANALYTICS", true),
			EnableNewsletter:  getBoolEnv("ENABLE_NEWSLETTER", false),
			RateLimitEnabled:  getBoolEnv("RATE_LIMIT_ENABLED", true),
			RateLimitRequests: getIntEnv("RATE_LIMIT_REQUESTS_PER_MINUTE", 60),
			CORSEnabled:       getBoolEnv("CORS_ENABLED", false),
			CORSOrigins:       getEnv("CORS_ALLOWED_ORIGINS", "*"),
		},
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// Check critical settings
	if c.Admin.Username == "admin" && c.Admin.Password == "admin123" {
		log.Println("WARNING: Using default admin credentials! Please change them in production.")
	}

	if c.Admin.Secret == "change-this-secret-key-in-production" {
		log.Println("WARNING: Using default admin secret! Please change it in production.")
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

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
