package handler

import (
	"net/http"
	"strconv"

	"github.com/caryxiao/meta-blog/internal/request"
	"github.com/caryxiao/meta-blog/internal/response"
	"github.com/caryxiao/meta-blog/internal/service"
	"github.com/caryxiao/meta-blog/internal/utils"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
	logService     *service.LogService
}

func NewCommentHandler(commentService *service.CommentService, logService *service.LogService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
		logService:     logService,
	}
}

// CreateComment creates a comment
func (h *CommentHandler) CreateComment(c *gin.Context) {
	logger := utils.NewLogger(c)
	resp := &response.Response{}

	// Get user ID
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("User not authenticated")
		resp.Fail(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Get post ID from URL path parameter
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		logger.Error("Invalid post ID format: " + err.Error())
		resp.Fail(c, http.StatusBadRequest, "Invalid post ID format")
		return
	}

	// Bind request parameters
	var req request.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Parameter binding failed: " + err.Error())
		resp.Fail(c, http.StatusBadRequest, "Parameter error: "+err.Error())
		return
	}

	// Create comment
	comment, err := h.commentService.CreateComment(&req, uint(postID), userID.(uint))
	if err != nil {
		logger.Error("Failed to create comment: " + err.Error())
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Log comment creation
	if logErr := h.logService.LogCommentAction(c, userID.(uint), comment.ID, "create_comment", req); logErr != nil {
		logger.Error("Failed to log comment creation: " + logErr.Error())
	}

	logger.Info("Comment created successfully")
	resp.Success(c, comment)
}

// GetComments gets post comment list
func (h *CommentHandler) GetComments(c *gin.Context) {
	logger := utils.NewLogger(c)
	resp := &response.Response{}

	// Get post ID from URL path parameter
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		logger.Error("Invalid post ID format: " + err.Error())
		resp.Fail(c, http.StatusBadRequest, "Invalid post ID format")
		return
	}

	// Bind query parameters
	var req request.GetCommentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Parameter binding failed: " + err.Error())
		resp.Fail(c, http.StatusBadRequest, "Parameter error: "+err.Error())
		return
	}

	// Set default values
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// Get comment list
	comments, err := h.commentService.GetCommentsByPostID(&req, uint(postID))
	if err != nil {
		logger.Error("Failed to get comment list: " + err.Error())
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("Comment list retrieved successfully")
	resp.Success(c, comments)
}
