package middleware

import (
	"strings"

	"github.com/caryxiao/meta-blog/internal/response"
	"github.com/caryxiao/meta-blog/internal/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth JWT authentication middleware
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &response.Response{}

		// Get token from Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			r.Fail(c, 40101, "Please provide the authentication token")
			c.Abort()
			return
		}

		// Check Bearer prefix
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.Fail(c, 40102, "Token format error")
			c.Abort()
			return
		}

		// Parse token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			r.Fail(c, 40103, "The token is invalid or expired")
			c.Abort()
			return
		}

		// Store user info in context
		c.Set("user_id", claims.UserID)
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
