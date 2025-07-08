package handler

import (
	"strconv"

	"github.com/caryxiao/meta-blog/internal/request"
	"github.com/caryxiao/meta-blog/internal/response"
	"github.com/caryxiao/meta-blog/internal/service"
	"github.com/caryxiao/meta-blog/internal/utils"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
	logService  *service.LogService
}

func NewPostHandler(postService *service.PostService, logService *service.LogService) *PostHandler {
	return &PostHandler{
		postService: postService,
		logService:  logService,
	}
}

// CreatePost creates a post
func (h *PostHandler) CreatePost(c *gin.Context) {
	logger := utils.NewLogger(c)
	logger.Info("Starting to create post")
	r := &response.Response{}

	var req request.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Parameter validation failed", "error", err.Error())
		r.Fail(c, 40001, "Parameter validation failed: "+err.Error())
		return
	}

	// Get user ID from JWT
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("User ID not found")
		r.Fail(c, 40101, "User authentication failed")
		return
	}

	logger.Info("Creating post", "title", req.Title, "userID", userID)
	post, err := h.postService.CreatePost(c.Request.Context(), &req, userID.(uint))
	if err != nil {
		logger.Error("Failed to create post", "error", err.Error())
		r.Fail(c, 50001, "Failed to create post: "+err.Error())
		return
	}

	// Log post creation
	if logErr := h.logService.LogPostAction(c, userID.(uint), post.ID, "create_post", req); logErr != nil {
		logger.Error("Failed to log post creation", "error", logErr.Error())
	}

	logger.Info("Post created successfully", "postID", post.ID)
	r.Success(c, post)
}

// GetPost gets post details
func (h *PostHandler) GetPost(c *gin.Context) {
	logger := utils.NewLogger(c)
	logger.Info("Starting to get post details")
	r := &response.Response{}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error("Invalid post ID format", "id", idStr, "error", err.Error())
		r.Fail(c, 40002, "Invalid post ID format")
		return
	}

	logger.Info("Getting post details", "postID", id)
	post, err := h.postService.GetPost(c.Request.Context(), id)
	if err != nil {
		logger.Error("Failed to get post", "postID", id, "error", err.Error())
		r.Fail(c, 40404, err.Error())
		return
	}

	logger.Info("Post retrieved successfully", "postID", id)
	r.Success(c, post)
}

// UpdatePost updates a post
func (h *PostHandler) UpdatePost(c *gin.Context) {
	logger := utils.NewLogger(c)
	logger.Info("Starting to update post")
	r := &response.Response{}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error("Invalid post ID format", "id", idStr, "error", err.Error())
		r.Fail(c, 40002, "Invalid post ID format")
		return
	}

	var req request.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Parameter validation failed", "error", err.Error())
		r.Fail(c, 40001, "Parameter validation failed: "+err.Error())
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("User ID not found")
		r.Fail(c, 40101, "User authentication failed")
		return
	}

	logger.Info("Updating post", "postID", id, "title", req.Title, "userID", userID)
	post, err := h.postService.UpdatePost(c.Request.Context(), id, &req, userID.(uint))
	if err != nil {
		logger.Error("Failed to update post", "postID", id, "error", err.Error())
		r.Fail(c, 50002, err.Error())
		return
	}

	// Log post update
	if logErr := h.logService.LogPostAction(c, userID.(uint), uint(id), "update_post", req); logErr != nil {
		logger.Error("Failed to log post update", "error", logErr.Error())
	}

	logger.Info("Post updated successfully", "postID", id)
	r.Success(c, post)
}

// DeletePost deletes a post
func (h *PostHandler) DeletePost(c *gin.Context) {
	logger := utils.NewLogger(c)
	logger.Info("Starting to delete post")
	r := &response.Response{}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error("Invalid post ID format", "id", idStr, "error", err.Error())
		r.Fail(c, 40002, "Invalid post ID format")
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("User ID not found")
		r.Fail(c, 40101, "User authentication failed")
		return
	}

	logger.Info("Deleting post", "postID", id, "userID", userID)
	err = h.postService.DeletePost(c.Request.Context(), id, userID.(uint))
	if err != nil {
		logger.Error("Failed to delete post", "postID", id, "error", err.Error())
		r.Fail(c, 50003, err.Error())
		return
	}

	// Log post deletion
	if logErr := h.logService.LogPostAction(c, userID.(uint), uint(id), "delete_post", nil); logErr != nil {
		logger.Error("Failed to log post deletion", "error", logErr.Error())
	}

	logger.Info("Post deleted successfully", "postID", id)
	r.Success(c, gin.H{"message": "Post deleted successfully"})
}

// ListPosts gets post list
func (h *PostHandler) ListPosts(c *gin.Context) {
	logger := utils.NewLogger(c)
	logger.Info("Starting to get post list")
	r := &response.Response{}

	var req request.ListPostRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Parameter validation failed", "error", err.Error())
		r.Fail(c, 40001, "Parameter validation failed: "+err.Error())
		return
	}

	logger.Info("Getting post list", "page", req.Page, "size", req.PageSize, "userID", req.UserID)
	result, err := h.postService.ListPosts(c.Request.Context(), &req)
	if err != nil {
		logger.Error("Failed to get post list", "error", err.Error())
		r.Fail(c, 50004, "Failed to get post list: "+err.Error())
		return
	}

	logger.Info("Post list retrieved successfully", "total", result.Total, "count", len(result.Items))
	r.Success(c, result)
}
