package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"
)

// AuthService handles authentication
type AuthService struct {
    username string
    password string
    sessions map[string]*Session
    mu       sync.RWMutex
}

// Session represents a user session
type Session struct {
    Token     string
    Username  string
    CreatedAt time.Time
    ExpiresAt time.Time
}

// NewAuthService creates a new AuthService
func NewAuthService(username, password string) *AuthService {
    return &AuthService{
        username: username,
        password: password,
        sessions: make(map[string]*Session),
    }
}

// Login authenticates a user and creates a session
func (as *AuthService) Login(username, password string) (string, error) {
    if username != as.username || password != as.password {
        return "", errors.New("invalid credentials")
    }

    token, err := generateToken()
    if err != nil {
        return "", err
    }

    session := &Session{
        Token:     token,
        Username:  username,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }

    as.mu.Lock()
    as.sessions[token] = session
    as.mu.Unlock()

    return token, nil
}

// ValidateSession checks if a session token is valid
func (as *AuthService) ValidateSession(token string) bool {
    as.mu.RLock()
    defer as.mu.RUnlock()

    session, exists := as.sessions[token]
    if !exists {
        return false
    }

    if time.Now().After(session.ExpiresAt) {
        return false
    }

    return true
}

// Logout removes a session
func (as *AuthService) Logout(token string) {
    as.mu.Lock()
    defer as.mu.Unlock()
    delete(as.sessions, token)
}

// CleanupExpiredSessions removes expired sessions
func (as *AuthService) CleanupExpiredSessions() {
    as.mu.Lock()
    defer as.mu.Unlock()

    now := time.Now()
    for token, session := range as.sessions {
        if now.After(session.ExpiresAt) {
            delete(as.sessions, token)
        }
    }
}

// generateToken generates a random session token
func generateToken() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}