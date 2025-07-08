package service

import (
	"context"
	"errors"
	"github.com/caryxiao/meta-blog/internal/model"
	"github.com/caryxiao/meta-blog/internal/repository"
	"github.com/caryxiao/meta-blog/internal/request"
	"github.com/caryxiao/meta-blog/internal/response"
	"gorm.io/gorm"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
	postRepo    *repository.PostRepository
}

func NewCommentService(commentRepo *repository.CommentRepository, postRepo *repository.PostRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (s *CommentService) CreateComment(req *request.CreateCommentRequest, postID uint, userID uint) (*response.CommentResponse, error) {
	// Check if the post exists
	_, err := s.postRepo.Get(context.Background(), int64(postID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	// Create comment
	comment := &model.Comment{
		Content: req.Content,
		PostID:  postID,
		UserID:  userID,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	// Return response
	return &response.CommentResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}, nil
}

func (s *CommentService) GetCommentsByPostID(req *request.GetCommentsRequest, postID uint) (*response.CommentListResponse, error) {
	// Check if the post exists
	_, err := s.postRepo.Get(context.Background(), int64(postID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	// Calculate offset
	offset := (req.Page - 1) * req.PageSize

	// Get comment list
	comments, total, err := s.commentRepo.GetByPostID(postID, offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	// Convert to response format
	var items []response.CommentResponse
	for _, comment := range comments {
		items = append(items, response.CommentResponse{
			ID:        comment.ID,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return &response.CommentListResponse{
		Items:    items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
