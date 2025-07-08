package repository

import (
	"github.com/caryxiao/meta-blog/internal/model"
	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

// Create creates a new log entry
func (r *LogRepository) Create(log *model.Log) error {
	return r.db.Create(log).Error
}

// GetByUserID gets logs by user ID with pagination
func (r *LogRepository) GetByUserID(userID uint, offset, limit int) ([]model.Log, int64, error) {
	var logs []model.Log
	var total int64

	query := r.db.Model(&model.Log{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// GetByAction gets logs by action with pagination
func (r *LogRepository) GetByAction(action string, offset, limit int) ([]model.Log, int64, error) {
	var logs []model.Log
	var total int64

	query := r.db.Model(&model.Log{}).Where("action = ?", action)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
