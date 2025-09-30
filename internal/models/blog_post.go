package models

import (
	"html/template"
	"time"
)

// BlogPost represents a blog post with all its metadata and content
type BlogPost struct {
	ID          int           `json:"id" db:"id"`
	Title       string        `json:"title" db:"title"`
	Content     template.HTML `json:"content" db:"content"`
	Excerpt     string        `json:"excerpt" db:"excerpt"`
	Author      string        `json:"author" db:"author"`
	Date        time.Time     `json:"date" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
	Tags        []string      `json:"tags" db:"tags"`
	Slug        string        `json:"slug" db:"slug"`
	Status      PostStatus    `json:"status" db:"status"`
	CategoryID  *int          `json:"category_id" db:"category_id"`
	ReadTime    int           `json:"read_time" db:"read_time"` // in minutes
	ViewCount   int           `json:"view_count" db:"view_count"`
	Featured    bool          `json:"featured" db:"featured"`
	PublishedAt *time.Time    `json:"published_at" db:"published_at"`
}

// PostMetadata represents the frontmatter metadata of a markdown post
type PostMetadata struct {
	Title      string    `yaml:"title"`
	Author     string    `yaml:"author"`
	Date       string    `yaml:"date"`
	Tags       []string  `yaml:"tags"`
	Excerpt    string    `yaml:"excerpt"`
	Slug       string    `yaml:"slug"`
	Status     string    `yaml:"status"`
	Category   string    `yaml:"category"`
	Featured   bool      `yaml:"featured"`
	Published  bool      `yaml:"published"`
}

// PostStatus represents the publication status of a post
type PostStatus string

const (
	StatusDraft     PostStatus = "draft"
	StatusPublished PostStatus = "published"
	StatusArchived  PostStatus = "archived"
)

// Category represents a blog category
type Category struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Slug        string `json:"slug" db:"slug"`
	Description string `json:"description" db:"description"`
	Color       string `json:"color" db:"color"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Tag represents a blog tag
type Tag struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Slug      string    `json:"slug" db:"slug"`
	PostCount int       `json:"post_count" db:"post_count"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// User represents a blog user/author
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Bio       string    `json:"bio" db:"bio"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Role      UserRole  `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// UserRole represents the role of a user
type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleEditor  UserRole = "editor"
	RoleAuthor  UserRole = "author"
	RoleViewer  UserRole = "viewer"
)

// Comment represents a comment on a blog post
type Comment struct {
	ID        int       `json:"id" db:"id"`
	PostID    int       `json:"post_id" db:"post_id"`
	Author    string    `json:"author" db:"author"`
	Email     string    `json:"email" db:"email"`
	Content   string    `json:"content" db:"content"`
	Status    CommentStatus `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CommentStatus represents the moderation status of a comment
type CommentStatus string


// PostMetrics represents metrics for a blog post
type PostMetrics struct {
    PostID          int       `json:"post_id" db:"post_id"`
    ViewCount       int       `json:"view_count" db:"view_count"`
    ShareCount      int       `json:"share_count" db:"share_count"`
    FacebookShares  int       `json:"facebook_shares" db:"facebook_shares"`
    TwitterShares   int       `json:"twitter_shares" db:"twitter_shares"`
    LinkedInShares  int       `json:"linkedin_shares" db:"linkedin_shares"`
    LastViewedAt    time.Time `json:"last_viewed_at" db:"last_viewed_at"`
    LastSharedAt    time.Time `json:"last_shared_at" db:"last_shared_at"`
}

// DashboardStats represents statistics for the admin dashboard
type DashboardStats struct {
    TotalPosts      int
    PublishedPosts  int
    DraftPosts      int
    ArchivedPosts   int
    TotalViews      int
    TotalShares     int
    TodayViews      int
    TodayShares     int
    PopularPosts    []BlogPost
    RecentPosts     []BlogPost
}

const (
	CommentPending   CommentStatus = "pending"
	CommentApproved  CommentStatus = "approved"
	CommentRejected  CommentStatus = "rejected"
	CommentSpam      CommentStatus = "spam"
)
