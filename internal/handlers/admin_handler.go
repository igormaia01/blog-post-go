package handlers

import (
	"net/http"
	"strconv"

	"blog-post/internal/services"

	"github.com/gin-gonic/gin"
)

// AdminHandler handles admin panel requests
type AdminHandler struct {
	postService *services.PostService
}

// NewAdminHandler creates a new AdminHandler instance
func NewAdminHandler(postService *services.PostService) *AdminHandler {
	return &AdminHandler{
		postService: postService,
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

	// Count posts by status
	var published, drafts, archived int
	for _, post := range posts {
		switch post.Status {
		case "published":
			published++
		case "draft":
			drafts++
		case "archived":
			archived++
		}
	}

	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"Title":     "Admin Dashboard",
		"Posts":     posts,
		"Published": published,
		"Drafts":    drafts,
		"Archived":  archived,
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

	// Simple authentication (in production, use proper auth)
	if username == "admin" && password == "admin123" {
		c.SetCookie("admin_session", "authenticated", 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/admin")
	} else {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"Title": "Admin Login",
			"Error": "Invalid credentials",
		})
	}
}

// Logout handles logout
func (ah *AdminHandler) Logout(c *gin.Context) {
	c.SetCookie("admin_session", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}
