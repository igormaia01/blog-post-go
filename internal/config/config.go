package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all configuration for our application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	Blog     BlogConfig
	Admin    AdminConfig
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
	Title       string
	Description string
	Author      string
	URL         string
	PostsPerPage int
	Theme       string
}

// AdminConfig holds admin panel configuration
type AdminConfig struct {
	Username string
	Password string
	Secret   string
}

// Load loads configuration from environment variables
func Load() *Config {
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
			Title:        getEnv("BLOG_TITLE", "Igor's Blog"),
			Description:  getEnv("BLOG_DESCRIPTION", "A personal blog about programming, technology, and life"),
			Author:       getEnv("BLOG_AUTHOR", "Igor"),
			URL:          getEnv("BLOG_URL", "http://localhost:3100"),
			PostsPerPage: getIntEnv("BLOG_POSTS_PER_PAGE", 10),
			Theme:        getEnv("BLOG_THEME", "default"),
		},
		Admin: AdminConfig{
			Username: getEnv("ADMIN_USERNAME", "admin"),
			Password: getEnv("ADMIN_PASSWORD", "admin123"),
			Secret:   getEnv("ADMIN_SECRET", "your-secret-key"),
		},
	}
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

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
