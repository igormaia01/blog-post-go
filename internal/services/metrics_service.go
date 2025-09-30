package services

import (
	"blog-post/internal/models"
	"sync"
	"time"
)

// MetricsService handles post metrics tracking
type MetricsService struct {
    metrics map[int]*models.PostMetrics
    mu      sync.RWMutex
}

// NewMetricsService creates a new MetricsService
func NewMetricsService() *MetricsService {
    return &MetricsService{
        metrics: make(map[int]*models.PostMetrics),
    }
}

// IncrementViewCount increments the view count for a post
func (ms *MetricsService) IncrementViewCount(postID int) error {
    ms.mu.Lock()
    defer ms.mu.Unlock()

    if _, exists := ms.metrics[postID]; !exists {
        ms.metrics[postID] = &models.PostMetrics{
            PostID: postID,
        }
    }

    ms.metrics[postID].ViewCount++
    ms.metrics[postID].LastViewedAt = time.Now()
    return nil
}

// IncrementShareCount increments the share count for a post
func (ms *MetricsService) IncrementShareCount(postID int, platform string) error {
    ms.mu.Lock()
    defer ms.mu.Unlock()

    if _, exists := ms.metrics[postID]; !exists {
        ms.metrics[postID] = &models.PostMetrics{
            PostID: postID,
        }
    }

    ms.metrics[postID].ShareCount++
    ms.metrics[postID].LastSharedAt = time.Now()

    switch platform {
    case "facebook":
        ms.metrics[postID].FacebookShares++
    case "twitter":
        ms.metrics[postID].TwitterShares++
    case "linkedin":
        ms.metrics[postID].LinkedInShares++
    }

    return nil
}

// GetMetrics returns metrics for a specific post
func (ms *MetricsService) GetMetrics(postID int) (*models.PostMetrics, error) {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    if metrics, exists := ms.metrics[postID]; exists {
        return metrics, nil
    }

    return &models.PostMetrics{PostID: postID}, nil
}

// GetAllMetrics returns all metrics
func (ms *MetricsService) GetAllMetrics() map[int]*models.PostMetrics {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    result := make(map[int]*models.PostMetrics)
    for k, v := range ms.metrics {
        result[k] = v
    }
    return result
}

// GetTotalViews returns total views across all posts
func (ms *MetricsService) GetTotalViews() int {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    total := 0
    for _, metrics := range ms.metrics {
        total += metrics.ViewCount
    }
    return total
}

// GetTotalShares returns total shares across all posts
func (ms *MetricsService) GetTotalShares() int {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    total := 0
    for _, metrics := range ms.metrics {
        total += metrics.ShareCount
    }
    return total
}

// GetTodayViews returns views from today
func (ms *MetricsService) GetTodayViews() int {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    today := time.Now().Truncate(24 * time.Hour)
    count := 0

    for _, metrics := range ms.metrics {
        if metrics.LastViewedAt.After(today) {
            count += metrics.ViewCount
        }
    }
    return count
}

// GetTodayShares returns shares from today
func (ms *MetricsService) GetTodayShares() int {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    today := time.Now().Truncate(24 * time.Hour)
    count := 0

    for _, metrics := range ms.metrics {
        if metrics.LastSharedAt.After(today) {
            count += metrics.ShareCount
        }
    }
    return count
}