package service

import (
	"encoding/json"
	"time"

	"github.com/caryxiao/meta-blog/internal/model"
	"github.com/caryxiao/meta-blog/internal/repository"
	"github.com/gin-gonic/gin"
)

type LogService struct {
	logRepo *repository.LogRepository
}

func NewLogService(logRepo *repository.LogRepository) *LogService {
	return &LogService{
		logRepo: logRepo,
	}
}

// LogAction logs an action to the database
func (s *LogService) LogAction(c *gin.Context, userID *uint, action string, targetType *string, targetID *uint, description interface{}) error {
	// Convert description to JSON string if it's not nil
	var descStr *string
	if description != nil {
		if descBytes, err := json.Marshal(description); err == nil {
			descString := string(descBytes)
			descStr = &descString
		}
	}

	// Get IP address
	ipAddress := c.ClientIP()

	// Get User-Agent
	userAgent := c.GetHeader("User-Agent")

	log := &model.Log{
		CreatedAt:   time.Now(),
		UserID:      userID,
		Action:      action,
		TargetType:  targetType,
		TargetID:    targetID,
		Description: descStr,
		IPAddress:   &ipAddress,
		UserAgent:   &userAgent,
	}

	return s.logRepo.Create(log)
}

// LogUserAction logs a user-related action
func (s *LogService) LogUserAction(c *gin.Context, userID uint, action string, description interface{}) error {
	targetType := "user"
	return s.LogAction(c, &userID, action, &targetType, &userID, description)
}

// LogPostAction logs a post-related action
func (s *LogService) LogPostAction(c *gin.Context, userID uint, postID uint, action string, description interface{}) error {
	targetType := "post"
	return s.LogAction(c, &userID, action, &targetType, &postID, description)
}

// LogCommentAction logs a comment-related action
func (s *LogService) LogCommentAction(c *gin.Context, userID uint, commentID uint, action string, description interface{}) error {
	targetType := "comment"
	return s.LogAction(c, &userID, action, &targetType, &commentID, description)
}
