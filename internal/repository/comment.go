package repository

import (
	"github.com/caryxiao/meta-blog/internal/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetByPostID(postID uint, offset, limit int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	// Get total count
	if err := r.db.Model(&model.Comment{}).Where("post_id = ?", postID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated data, ordered by creation time descending
	if err := r.db.Where("post_id = ?", postID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

func (r *CommentRepository) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Comment{}, id).Error
}
