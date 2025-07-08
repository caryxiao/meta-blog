package handler

import (
	"fmt"
	"github.com/caryxiao/meta-blog/internal/model"
	"github.com/caryxiao/meta-blog/internal/repository"
	"github.com/caryxiao/meta-blog/internal/request"
	"github.com/caryxiao/meta-blog/internal/response"
	"github.com/caryxiao/meta-blog/internal/service"
	"github.com/caryxiao/meta-blog/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo   *repository.UserRepository
	logService *service.LogService
}

func NewUserHandler(userRepo *repository.UserRepository, logService *service.LogService) *UserHandler {
	return &UserHandler{
		userRepo:   userRepo,
		logService: logService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	logger := utils.NewLogger(c)
	var req request.RegisterRequest
	r := &response.Response{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Request Data Validate Failed: %s", err.Error())
		r.Fail(c, 40001, fmt.Sprintf("Validate Failed: %s", err.Error()))
		return
	}

	// Check if username already exists
	usernameExists, err := h.userRepo.CheckUsernameExists(c, req.Username)
	if err != nil {
		logger.Error("Database Error: %s", err.Error())
		r.Fail(c, 50001, fmt.Sprintf("Database Error: %s", err.Error()))
		return
	}

	if usernameExists {
		logger.Warn("username is exists: %s", req.Username)
		r.Fail(c, 40003, "username is exists")
		return
	}

	// Check if email already exists
	emailExists, err := h.userRepo.CheckEmailExists(c, req.Email)
	if err != nil {
		r.Fail(c, 50002, fmt.Sprintf("Database Error: %s", err.Error()))
		return
	}
	if emailExists {
		r.Fail(c, 40004, "email is exists")
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		r.Fail(c, 50003, fmt.Sprintf("Password Hash Error: %s", err.Error()))
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: hashedPassword,
		Username: req.Username,
	}

	err = h.userRepo.Create(c, u)
	if err != nil {
		logger.Error("Create User Failed: %s", err.Error())
		r.Fail(c, 40002, fmt.Sprintf("Create User Failed: %s", err.Error()))
		return
	}
	logger.Info("User Create successful, ID: %d", u.ID)

	// Log user registration
	if logErr := h.logService.LogUserAction(c, u.ID, "register", req); logErr != nil {
		logger.Error("Failed to log user registration: %s", logErr.Error())
	}
	// Generate JWT token
	token, err := utils.GenerateToken(u.ID, u.Username)
	if err != nil {
		r.Fail(c, 50004, fmt.Sprintf("Token Generate Error: %s", err.Error()))
		return
	}

	// Return user info and token
	loginResp := map[string]interface{}{
		"user": &response.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
		"token": token,
	}
	r.Success(c, loginResp)
}

// GetProfile gets user profile (authentication required)
func (h *UserHandler) GetProfile(c *gin.Context) {
	r := &response.Response{}

	// Get user info from context
	userID, exists := c.Get("user_id")
	if !exists {
		r.Fail(c, 40104, "get user info error")
		return
	}

	// Query user info by user ID
	u, err := h.userRepo.Get(c, fmt.Sprintf("%d", userID))
	if err != nil {
		r.Fail(c, 40006, "user is not exists")
		return
	}

	// Return user info
	userResp := &response.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
	r.Success(c, userResp)
}

// Login user login
func (h *UserHandler) Login(c *gin.Context) {
	logger := utils.NewLogger(c)

	var req request.LoginRequest
	r := &response.Response{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Request Data Validate Failed: %s", err.Error())
		r.Fail(c, 40001, fmt.Sprintf("Validate Failed: %s", err.Error()))
		return
	}

	// Find user by username
	u, err := h.userRepo.GetByUsername(c, req.Username)
	if err != nil {
		logger.Warn("User name or password error: %s", req.Username)
		r.Fail(c, 40005, "Username or password error")
		return
	}

	// Verify password
	if !utils.CheckPassword(u.Password, req.Password) {
		logger.Warn("Validate Password Failed: %s", req.Username)
		r.Fail(c, 40005, "username or password error")
		return
	}

	// Generate JWT token
	logger.Info("Generate JWT token: %s", req.Username)
	token, err := utils.GenerateToken(u.ID, u.Username)
	if err != nil {
		logger.Error("JWT token generate failed: %s", err.Error())
		r.Fail(c, 50004, fmt.Sprintf("Token Generate Error: %s", err.Error()))
		return
	}

	// Return user info and token
	logger.Info("User login success: %s, ID: %d", u.Username, u.ID)
	loginResp := map[string]interface{}{
		"user": &response.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
		"token": token,
	}
	r.Success(c, loginResp)
}
