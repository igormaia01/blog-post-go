package handlers

import (
	"net/http"
	"strconv"

	"blog-post/internal/services"

	"github.com/gin-gonic/gin"
)

// BlogHandler handles blog-related HTTP requests
type BlogHandler struct {
	postService    *services.PostService
	metricsService *services.MetricsService
}

// NewBlogHandler creates a new BlogHandler instance
func NewBlogHandler(postService *services.PostService, metricsService *services.MetricsService) *BlogHandler {
	return &BlogHandler{
		postService:    postService,
		metricsService: metricsService,
	}
}

// Home redirects to posts list
func (bh *BlogHandler) Home(c *gin.Context) {
	c.Redirect(http.StatusFound, "/posts")
}

// PostsList renders the list of blog posts
func (bh *BlogHandler) PostsList(c *gin.Context) {
	posts, err := bh.postService.GetPublishedPosts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"Error": "Failed to load posts",
		})
		return
	}

	// Get blog config from context or environment
	blogTitle := c.GetString("BlogTitle")
	blogDescription := c.GetString("BlogDescription")

	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"SiteTitle":       blogTitle,
		"SiteDescription": blogDescription,
		"Posts":           posts,
	})
}

// PostDetail renders a single blog post
func (bh *BlogHandler) PostDetail(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.HTML(http.StatusNotFound, "detail.tmpl", gin.H{
			"Error": "Invalid post ID",
		})
		return
	}

	post, err := bh.postService.GetPostByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "detail.tmpl", gin.H{
			"Error": "Post not found",
		})
		return
	}

	// Increment view count
	bh.metricsService.IncrementViewCount(post.ID)

	// Get related posts
	relatedPosts, _ := bh.postService.GetRelatedPosts(post, 3)

	// Get blog config from context
	blogTitle := c.GetString("BlogTitle")

	c.HTML(http.StatusOK, "detail.tmpl", gin.H{
		"SiteTitle":    blogTitle,
		"Post":         post,
		"RelatedPosts": relatedPosts,
	})
}

// PostBySlug renders a post by its slug
func (bh *BlogHandler) PostBySlug(c *gin.Context) {
	slug := c.Param("slug")
	
	post, err := bh.postService.GetPostBySlug(slug)
	if err != nil {
		c.HTML(http.StatusNotFound, "detail.tmpl", gin.H{
			"Error": "Post not found",
		})
		return
	}

	// Increment view count
	bh.metricsService.IncrementViewCount(post.ID)

	// Get related posts
	relatedPosts, _ := bh.postService.GetRelatedPosts(post, 3)

	// Get blog config from context
	blogTitle := c.GetString("BlogTitle")

	c.HTML(http.StatusOK, "detail.tmpl", gin.H{
		"SiteTitle":    blogTitle,
		"Post":         post,
		"RelatedPosts": relatedPosts,
	})
}

// SearchPosts handles post search requests
func (bh *BlogHandler) SearchPosts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.Redirect(http.StatusFound, "/posts")
		return
	}

	posts, err := bh.postService.SearchPosts(query)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "search.tmpl", gin.H{
			"Error": "Search failed",
			"Query": query,
		})
		return
	}

	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"SiteTitle": "Search Results",
		"Query":     query,
		"Posts":     posts,
	})
}

// PostsByTag renders posts filtered by tag
func (bh *BlogHandler) PostsByTag(c *gin.Context) {
	tag := c.Param("tag")
	
	posts, err := bh.postService.GetPostsByTag(tag)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "tag.tmpl", gin.H{
			"Error": "Failed to load posts",
			"Tag":   tag,
		})
		return
	}

	c.HTML(http.StatusOK, "tag.tmpl", gin.H{
		"SiteTitle": "Posts tagged with: " + tag,
		"Tag":       tag,
		"Posts":     posts,
	})
}

// TagsList renders all available tags
func (bh *BlogHandler) TagsList(c *gin.Context) {
	tags, err := bh.postService.GetAllTags()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "tags.tmpl", gin.H{
			"Error": "Failed to load tags",
		})
		return
	}

	c.HTML(http.StatusOK, "tags.tmpl", gin.H{
		"SiteTitle": "All Tags",
		"Tags":      tags,
	})
}

// RSSFeed generates RSS feed
func (bh *BlogHandler) RSSFeed(c *gin.Context) {
	posts, err := bh.postService.GetPublishedPosts()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate RSS feed")
		return
	}

	// Limit to 20 most recent posts
	if len(posts) > 20 {
		posts = posts[:20]
	}

	// Get blog config from context
	blogTitle := c.GetString("BlogTitle")
	blogURL := c.GetString("BlogURL")

	c.Header("Content-Type", "application/rss+xml")
	c.HTML(http.StatusOK, "rss.tmpl", gin.H{
		"SiteTitle": blogTitle,
		"SiteURL":   blogURL,
		"Posts":     posts,
	})
}

// Sitemap generates XML sitemap
func (bh *BlogHandler) Sitemap(c *gin.Context) {
	posts, err := bh.postService.GetPublishedPosts()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate sitemap")
		return
	}

	// Get blog config from context
	blogURL := c.GetString("BlogURL")

	c.Header("Content-Type", "application/xml")
	c.HTML(http.StatusOK, "sitemap.tmpl", gin.H{
		"SiteURL": blogURL,
		"Posts":   posts,
	})
}

// TrackShare handles share tracking
func (bh *BlogHandler) TrackShare(c *gin.Context) {
	var request struct {
		PostID   int    `json:"post_id" binding:"required"`
		Platform string `json:"platform" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := bh.metricsService.IncrementShareCount(request.PostID, request.Platform)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to track share"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
