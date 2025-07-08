package model

import "time"

type Log struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time `gorm:"type:datetime(3)" json:"created_at"`
	UserID      *uint     `gorm:"column:user_id" json:"user_id"`
	Action      string    `gorm:"type:varchar(100);not null" json:"action"`
	TargetType  *string   `gorm:"type:varchar(100)" json:"target_type"`
	TargetID    *uint     `gorm:"column:target_id" json:"target_id"`
	Description *string   `gorm:"type:text" json:"description"`
	IPAddress   *string   `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent   *string   `gorm:"type:varchar(255)" json:"user_agent"`
}

func (Log) TableName() string {
	return "logs"
}
