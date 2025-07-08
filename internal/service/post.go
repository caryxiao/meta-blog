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

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *request.CreatePostRequest, userID uint) (*response.PostResponse, error) {
	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	err := s.postRepo.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return s.modelToResponse(post), nil
}

func (s *PostService) GetPost(ctx context.Context, id int64) (*response.PostResponse, error) {
	post, err := s.postRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return s.modelToResponse(post), nil
}

func (s *PostService) UpdatePost(ctx context.Context, id int64, req *request.UpdatePostRequest, userID uint) (*response.PostResponse, error) {
	// First check if the post exists and belongs to the current user
	_, err := s.postRepo.GetByUserID(ctx, userID, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found or access denied")
		}
		return nil, err
	}

	// Update the post
	updatePost := &model.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	err = s.postRepo.Update(ctx, id, updatePost)
	if err != nil {
		return nil, err
	}

	// Retrieve the updated post
	updatedPost, err := s.postRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.modelToResponse(updatedPost), nil
}

func (s *PostService) DeletePost(ctx context.Context, id int64, userID uint) error {
	// First check if the post exists and belongs to the current user
	_, err := s.postRepo.GetByUserID(ctx, userID, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found or access denied")
		}
		return err
	}

	return s.postRepo.Delete(ctx, id)
}

func (s *PostService) ListPosts(ctx context.Context, req *request.ListPostRequest) (*response.PostListResponse, error) {
	// Set default values
	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.PageSize
	if size <= 0 {
		size = 10
	}

	posts, total, err := s.postRepo.List(ctx, page, size, req.UserID)
	if err != nil {
		return nil, err
	}

	postResponses := make([]response.PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = *s.modelToResponse(&post)
	}

	return &response.PostListResponse{
		Items:    postResponses,
		Total:    total,
		Page:     page,
		PageSize: size,
	}, nil
}

func (s *PostService) modelToResponse(post *model.Post) *response.PostResponse {
	return &response.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
