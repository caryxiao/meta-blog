package di

import (
	"github.com/caryxiao/meta-blog/internal/handler"
	"github.com/caryxiao/meta-blog/internal/repository"
	"github.com/caryxiao/meta-blog/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler    *handler.UserHandler
	PostHandler    *handler.PostHandler
	CommentHandler *handler.CommentHandler
}

func NewContainer(db *gorm.DB) *Container {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	logRepo := repository.NewLogRepository(db)

	// Initialize services
	logService := service.NewLogService(logRepo)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo, postRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userRepo, logService)
	postHandler := handler.NewPostHandler(postService, logService)
	commentHandler := handler.NewCommentHandler(commentService, logService)

	return &Container{
		UserHandler:    userHandler,
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
	}
}
