package repository

import (
	"context"
	"github.com/caryxiao/meta-blog/internal/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(&post).Error
}

func (r *PostRepository) Get(ctx context.Context, id int64) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).First(&post, id).Error
	return &post, err
}

func (r *PostRepository) Update(ctx context.Context, id int64, post *model.Post) error {
	return r.db.WithContext(ctx).Model(&model.Post{}).Where("id = ?", id).Updates(post).Error
}

func (r *PostRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Post{}, id).Error
}

func (r *PostRepository) List(ctx context.Context, page, size int, userID *uint) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Post{})

	// Filter by user ID if specified
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	// Get total count
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Paginated query
	offset := (page - 1) * size
	err = query.Offset(offset).Limit(size).Order("created_at DESC").Find(&posts).Error

	return posts, total, err
}

func (r *PostRepository) GetByUserID(ctx context.Context, userID uint, id int64) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&post).Error
	return &post, err
}
