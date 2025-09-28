package services

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"blog-post/internal/models"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

// PostService handles all post-related operations
type PostService struct {
	postsDir string
	cache    CacheService
}

// NewPostService creates a new PostService instance
func NewPostService(postsDir string, cache CacheService) *PostService {
	return &PostService{
		postsDir: postsDir,
		cache:    cache,
	}
}

// LoadPostsFromMarkdown loads all posts from markdown files
func (ps *PostService) LoadPostsFromMarkdown() ([]models.BlogPost, error) {
	// Check cache first
	if cached, found := ps.cache.Get("all_posts"); found {
		if posts, ok := cached.([]models.BlogPost); ok {
			return posts, nil
		}
	}

	var posts []models.BlogPost
	
	files, err := filepath.Glob(filepath.Join(ps.postsDir, "*.md"))
	fmt.Println(files)
	if err != nil {
		return nil, err
	}

	for i, file := range files {
		post, err := ps.loadPostFromFile(file, i+1)
		if err != nil {
			continue // Skip invalid posts
		}
		posts = append(posts, *post)
	}

	// Sort posts by date (newest first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	// Cache the result
	ps.cache.Set("all_posts", posts, 1*time.Hour)

	return posts, nil
}

// loadPostFromFile loads a single post from a markdown file
func (ps *PostService) loadPostFromFile(filePath string, id int) (*models.BlogPost, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}


	parts := strings.Split(string(content), "---")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown format: %s", filePath)
	}

	var metadata models.PostMetadata
	if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
		return nil, err
	}

	// Parse date
	date, err := time.Parse("2006-01-02", metadata.Date)
	if err != nil {
		date = time.Now()
	}

	// Convert markdown to HTML
	markdownContent := strings.Join(parts[2:], "---")
	htmlContent := blackfriday.Run([]byte(markdownContent))

	// Calculate reading time (average 200 words per minute)
	wordCount := len(strings.Fields(markdownContent))
	readTime := (wordCount + 199) / 200 // Round up

	// Determine status
	status := models.StatusPublished
	if metadata.Status != "" {
		status = models.PostStatus(metadata.Status)
	} else if !metadata.Published {
		status = models.StatusDraft
	}

	post := &models.BlogPost{
		ID:        id,
		Title:     metadata.Title,
		Content:   template.HTML(htmlContent),
		Excerpt:   metadata.Excerpt,
		Author:    metadata.Author,
		Date:      date,
		UpdatedAt: time.Now(),
		Tags:      metadata.Tags,
		Slug:      metadata.Slug,
		Status:    status,
		ReadTime:  readTime,
		Featured:  metadata.Featured,
	}

	// Set published date if published
	if status == models.StatusPublished {
		post.PublishedAt = &date
	}

	return post, nil
}

// GetPostByID retrieves a post by its ID
func (ps *PostService) GetPostByID(id int) (*models.BlogPost, error) {
	posts, err := ps.LoadPostsFromMarkdown()
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		if post.ID == id {
			return &post, nil
		}
	}

	return nil, fmt.Errorf("post with ID %d not found", id)
}

// GetPostBySlug retrieves a post by its slug
func (ps *PostService) GetPostBySlug(slug string) (*models.BlogPost, error) {
	posts, err := ps.LoadPostsFromMarkdown()
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		if post.Slug == slug {
			return &post, nil
		}
	}

	return nil, fmt.Errorf("post with slug %s not found", slug)
}

// GetPublishedPosts retrieves only published posts
func (ps *PostService) GetPublishedPosts() ([]models.BlogPost, error) {
	posts, err := ps.LoadPostsFromMarkdown()
	if err != nil {
		return nil, err
	}
	fmt.Println(posts)
	var published []models.BlogPost
	for _, post := range posts {
		if post.Status == models.StatusPublished {
			published = append(published, post)
		}
	}
	fmt.Println(published)
	return published, nil
}

// GetPostsByTag retrieves posts filtered by tag
func (ps *PostService) GetPostsByTag(tag string) ([]models.BlogPost, error) {
	posts, err := ps.GetPublishedPosts()
	if err != nil {
		return nil, err
	}

	var filtered []models.BlogPost
	for _, post := range posts {
		for _, postTag := range post.Tags {
			if strings.EqualFold(postTag, tag) {
				filtered = append(filtered, post)
				break
			}
		}
	}

	return filtered, nil
}

// GetRelatedPosts retrieves related posts based on tags
func (ps *PostService) GetRelatedPosts(post *models.BlogPost, limit int) ([]models.BlogPost, error) {
	posts, err := ps.GetPublishedPosts()
	if err != nil {
		return nil, err
	}

	var related []models.BlogPost
	for _, p := range posts {
		if p.ID == post.ID {
			continue
		}

		// Check for common tags
		for _, tag := range post.Tags {
			for _, pTag := range p.Tags {
				if strings.EqualFold(tag, pTag) {
					related = append(related, p)
					break
				}
			}
		}

		if len(related) >= limit {
			break
		}
	}

	return related, nil
}

// SearchPosts searches posts by title, content, and tags
func (ps *PostService) SearchPosts(query string) ([]models.BlogPost, error) {
	posts, err := ps.GetPublishedPosts()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []models.BlogPost

	for _, post := range posts {
		// Search in title
		if strings.Contains(strings.ToLower(post.Title), query) {
			results = append(results, post)
			continue
		}

		// Search in excerpt
		if strings.Contains(strings.ToLower(post.Excerpt), query) {
			results = append(results, post)
			continue
		}

		// Search in tags
		for _, tag := range post.Tags {
			if strings.Contains(strings.ToLower(tag), query) {
				results = append(results, post)
				break
			}
		}

		// Search in content (convert HTML to text for searching)
		contentText := ps.htmlToText(string(post.Content))
		if strings.Contains(strings.ToLower(contentText), query) {
			results = append(results, post)
		}
	}

	return results, nil
}

// htmlToText converts HTML content to plain text for searching
func (ps *PostService) htmlToText(html string) string {
	// Simple HTML tag removal - in production, use a proper HTML parser
	html = strings.ReplaceAll(html, "<br>", " ")
	html = strings.ReplaceAll(html, "<br/>", " ")
	html = strings.ReplaceAll(html, "<br />", " ")
	
	// Remove HTML tags using a simple regex-like approach
	var result bytes.Buffer
	inTag := false
	for _, char := range html {
		if char == '<' {
			inTag = true
		} else if char == '>' {
			inTag = false
		} else if !inTag {
			result.WriteRune(char)
		}
	}
	
	return result.String()
}

// GetAllTags retrieves all unique tags from published posts
func (ps *PostService) GetAllTags() ([]string, error) {
	posts, err := ps.GetPublishedPosts()
	if err != nil {
		return nil, err
	}

	tagMap := make(map[string]bool)
	for _, post := range posts {
		for _, tag := range post.Tags {
			tagMap[tag] = true
		}
	}

	var tags []string
	for tag := range tagMap {
		tags = append(tags, tag)
	}

	sort.Strings(tags)
	return tags, nil
}

// IncrementViewCount increments the view count for a post
func (ps *PostService) IncrementViewCount(postID int) error {
	// In a real implementation, this would update the database
	// For now, we'll just log it
	fmt.Printf("Post %d view count incremented\n", postID)
	return nil
}
