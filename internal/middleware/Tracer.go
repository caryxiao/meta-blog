package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), "traceId", traceId)
		c.Request = c.Request.WithContext(ctx)
		// Set trace id to gin context for easy access from other places
		c.Set("traceId", traceId)
		c.Next()
	}
}
