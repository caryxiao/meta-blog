package router

import (
	"github.com/caryxiao/meta-blog/internal/di"
	"github.com/caryxiao/meta-blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, c *di.Container) {
	api := r.Group("/api")

	// Public routes (no authentication required)
	userGroup := api.Group("/user")
	{
		userGroup.POST("/register", c.UserHandler.Register)
		userGroup.POST("/login", c.UserHandler.Login)
	}

	// Public post routes
	postGroup := api.Group("/posts")
	{
		postGroup.GET("", c.PostHandler.ListPosts)                   // Get post list
		postGroup.GET("/:id", c.PostHandler.GetPost)                 // Get post details
		postGroup.GET("/:id/comments", c.CommentHandler.GetComments) // Get post comment list
	}

	// Routes requiring authentication
	authGroup := api.Group("/user")
	authGroup.Use(middleware.JWTAuth())
	{
		authGroup.GET("/profile", c.UserHandler.GetProfile)
	}

	// Post routes requiring authentication
	authPostGroup := api.Group("/posts")
	authPostGroup.Use(middleware.JWTAuth())
	{
		authPostGroup.POST("", c.PostHandler.CreatePost)       // Create post
		authPostGroup.PUT("/:id", c.PostHandler.UpdatePost)    // Update post
		authPostGroup.DELETE("/:id", c.PostHandler.DeletePost) // Delete post
	}

	// Add create comment route to authenticated post routes
	authPostGroup.POST("/:id/comments", c.CommentHandler.CreateComment) // Create comment
}
