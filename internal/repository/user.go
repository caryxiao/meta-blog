package repository

import (
	"context"
	"github.com/caryxiao/meta-blog/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	return r.db.WithContext(ctx).Create(&u).Error
}

func (r *UserRepository) Get(ctx context.Context, id string) (*model.User, error) {
	var u model.User
	return &u, r.db.WithContext(ctx).First(&u, "id = ?", id).Error
}

func (r *UserRepository) Update(ctx context.Context, u *model.User) error {
	return r.db.WithContext(ctx).Save(&u).Error
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}

// CheckUsernameExists checks if username already exists
func (r *UserRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// CheckEmailExists checks if email already exists
func (r *UserRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// GetByUsername gets user info by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var u model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
	return &u, err
}
