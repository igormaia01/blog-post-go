package middleware

import (
	"blog-post/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates authentication middleware
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := c.Cookie("admin_session")
        if err != nil {
            c.Redirect(http.StatusFound, "/admin/login")
            c.Abort()
            return
        }

        if !authService.ValidateSession(token) {
            c.Redirect(http.StatusFound, "/admin/login")
            c.Abort()
            return
        }

        c.Next()
    }
}