package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if user is authenticated for admin routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the request is for admin routes
		if c.Request.URL.Path[:6] == "/admin" && c.Request.URL.Path != "/admin/login" {
			session, err := c.Cookie("admin_session")
			if err != nil || session != "authenticated" {
				c.Redirect(http.StatusFound, "/admin/login")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
