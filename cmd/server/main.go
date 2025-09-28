package main

import (
	"log"

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

	// Initialize handlers
	blogHandler := handlers.NewBlogHandler(postService)
	adminHandler := handlers.NewAdminHandler(postService)

	// Setup Gin router
	router := gin.Default()

	// Middleware
	router.Use(middleware.LoggingMiddleware())
	// router.Use(middleware.AuthMiddleware())

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

	// Admin routes
	admin := router.Group("/admin")
	{
		admin.GET("/login", adminHandler.Login)
		admin.POST("/login", adminHandler.LoginPost)
		admin.GET("/logout", adminHandler.Logout)
		admin.GET("/", adminHandler.Dashboard)
		admin.GET("/posts", adminHandler.PostsList)
		admin.GET("/posts/new", adminHandler.PostEdit)
		admin.GET("/posts/:id/edit", adminHandler.PostEdit)
		admin.POST("/posts", adminHandler.PostSave)
		admin.POST("/posts/:id", adminHandler.PostSave)
		admin.DELETE("/posts/:id", adminHandler.PostDelete)
		admin.GET("/tags", adminHandler.TagsList)
		admin.GET("/settings", adminHandler.Settings)
	}

	// Start server
	log.Printf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(router.Run(cfg.Server.Host + ":" + cfg.Server.Port))
}
