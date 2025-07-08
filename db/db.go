package db

import (
	"github.com/caryxiao/meta-blog/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Init(cfg *config.AppConfig) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		DB, err = gorm.Open(mysql.Open(cfg.Database.GormDSN()), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			return
		}
		sqlDB, err := DB.DB()
		if err != nil {
			return
		}
		sqlDB.SetMaxOpenConns(cfg.Database.MaxConn)
		sqlDB.SetMaxIdleConns(cfg.Database.MaxIdle)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.MaxLife) * time.Second)
	})
	return DB, err
}
