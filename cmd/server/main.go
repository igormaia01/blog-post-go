package main

import (
	"log"
	"time"

	"blog-post/internal/config"
	"blog-post/internal/handlers"
	"blog-post/internal/middleware"
	"blog-post/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services
	cache := services.NewMemoryCache()
	postService := services.NewPostService("posts", cache)
	metricsService := services.NewMetricsService()
	authService := services.NewAuthService(cfg.Admin.Username, cfg.Admin.Password)

	// Start session cleanup routine
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			authService.CleanupExpiredSessions()
		}
	}()

	// Initialize handlers
	blogHandler := handlers.NewBlogHandler(postService, metricsService)
	adminHandler := handlers.NewAdminHandler(postService, metricsService, authService)

	// Setup Gin router
	router := gin.Default()

	// Middleware
	router.Use(middleware.LoggingMiddleware())

	// Load HTML templates
	router.LoadHTMLFiles(
		"web/templates/list.tmpl",
		"web/templates/detail.tmpl",
		"web/templates/index.tmpl",
		"web/templates/search.tmpl",
		"web/templates/tag.tmpl",
		"web/templates/tags.tmpl",
		"web/templates/admin/dashboard.tmpl",
		"web/templates/admin/login.tmpl",
	)

	// Static files
	router.Static("/static", "./web/static")
	router.Static("/images", "./images")

	// Blog routes
	router.GET("/", blogHandler.Home)
	router.GET("/posts", blogHandler.PostsList)
	router.GET("/post/:id", blogHandler.PostDetail)
	router.GET("/post/slug/:slug", blogHandler.PostBySlug)
	router.GET("/search", blogHandler.SearchPosts)
	router.GET("/tag/:tag", blogHandler.PostsByTag)
	router.GET("/tags", blogHandler.TagsList)
	router.GET("/rss.xml", blogHandler.RSSFeed)
	router.GET("/sitemap.xml", blogHandler.Sitemap)

	// API routes
	api := router.Group("/api")
	{
		api.POST("/track-share", blogHandler.TrackShare)
	}

	// Admin routes (public)
	router.GET("/admin/login", adminHandler.Login)
	router.POST("/admin/login", adminHandler.LoginPost)

	// Admin routes (protected)
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(authService))
	{
		admin.GET("/", adminHandler.Dashboard)
	}

	// Start server
	log.Printf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(router.Run(cfg.Server.Host + ":" + cfg.Server.Port))
}
