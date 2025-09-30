package handlers

import (
	"net/http"
	"strconv"

	"blog-post/internal/models"
	"blog-post/internal/services"

	"github.com/gin-gonic/gin"
)

// AdminHandler handles admin panel requests
type AdminHandler struct {
	postService    *services.PostService
	metricsService *services.MetricsService
	authService    *services.AuthService
}

// NewAdminHandler creates a new AdminHandler instance
func NewAdminHandler(postService *services.PostService, metricsService *services.MetricsService, authService *services.AuthService) *AdminHandler {
	return &AdminHandler{
		postService:    postService,
		metricsService: metricsService,
		authService:    authService,
	}
}

// Dashboard renders the admin dashboard
func (ah *AdminHandler) Dashboard(c *gin.Context) {
	posts, err := ah.postService.LoadPostsFromMarkdown()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"Error": "Failed to load posts",
		})
		return
	}

	// Get metrics for all posts
	allMetrics := ah.metricsService.GetAllMetrics()

	// Enrich posts with metrics
	postsWithMetrics := make([]gin.H, 0)
	for _, post := range posts {
		metrics, _ := ah.metricsService.GetMetrics(post.ID)
		postsWithMetrics = append(postsWithMetrics, gin.H{
			"Post":    post,
			"Metrics": metrics,
		})
	}

	// Count posts by status
	var published, drafts, archived int
	for _, post := range posts {
		switch post.Status {
		case models.StatusPublished:
			published++
		case models.StatusDraft:
			drafts++
		case models.StatusArchived:
			archived++
		}
	}

	stats := models.DashboardStats{
		TotalPosts:     len(posts),
		PublishedPosts: published,
		DraftPosts:     drafts,
		ArchivedPosts:  archived,
		TotalViews:     ah.metricsService.GetTotalViews(),
		TotalShares:    ah.metricsService.GetTotalShares(),
		TodayViews:     ah.metricsService.GetTodayViews(),
		TodayShares:    ah.metricsService.GetTodayShares(),
	}

	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"Title":       "Admin Dashboard",
		"Posts":       postsWithMetrics,
		"Stats":       stats,
		"Published":   published,
		"Drafts":      drafts,
		"Archived":    archived,
		"TotalViews":  stats.TotalViews,
		"TotalShares": stats.TotalShares,
		"TodayViews":  stats.TodayViews,
		"TodayShares": stats.TodayShares,
		"AllMetrics":  allMetrics,
	})
}

// PostsList renders the posts management page
func (ah *AdminHandler) PostsList(c *gin.Context) {
	posts, err := ah.postService.LoadPostsFromMarkdown()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"Error": "Failed to load posts",
		})
		return
	}

	c.HTML(http.StatusOK, "posts.tmpl", gin.H{
		"Title": "Manage Posts",
		"Posts": posts,
	})
}

// PostEdit renders the post edit form
func (ah *AdminHandler) PostEdit(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "new" {
		// New post form
		c.HTML(http.StatusOK, "post_edit.tmpl", gin.H{
			"Title": "Create New Post",
			"Post":  nil,
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.tmpl", gin.H{
			"Error": "Invalid post ID",
		})
		return
	}

	post, err := ah.postService.GetPostByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.tmpl", gin.H{
			"Error": "Post not found",
		})
		return
	}

	c.HTML(http.StatusOK, "post_edit.tmpl", gin.H{
		"Title": "Edit Post",
		"Post":  post,
	})
}

// PostSave handles post creation/update
func (ah *AdminHandler) PostSave(c *gin.Context) {
	// This would handle saving posts
	// For now, just redirect back to posts list
	c.Redirect(http.StatusFound, "/posts")
}

// PostDelete handles post deletion
func (ah *AdminHandler) PostDelete(c *gin.Context) {
	idParam := c.Param("id")
	_, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// This would handle post deletion
	// For now, just return success
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// TagsList renders the tags management page
func (ah *AdminHandler) TagsList(c *gin.Context) {
	tags, err := ah.postService.GetAllTags()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"Error": "Failed to load tags",
		})
		return
	}

	c.HTML(http.StatusOK, "tags.tmpl", gin.H{
		"Title": "Manage Tags",
		"Tags":  tags,
	})
}

// Settings renders the settings page
func (ah *AdminHandler) Settings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.tmpl", gin.H{
		"Title": "Settings",
	})
}

// Login renders the login page
func (ah *AdminHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"Title": "Admin Login",
	})
}

// LoginPost handles login form submission
func (ah *AdminHandler) LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	token, err := ah.authService.Login(username, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{
			"Title": "Admin Login",
			"Error": "Invalid credentials. Please try again.",
		})
		return
	}

	c.SetCookie("admin_session", token, 86400, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin")
}

// Logout handles logout
func (ah *AdminHandler) Logout(c *gin.Context) {
	token, _ := c.Cookie("admin_session")
	ah.authService.Logout(token)
	c.SetCookie("admin_session", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin/login")
}
