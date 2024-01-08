package database

import (
	"fmt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func Init(conf *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if conf.Env == "test" {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(conf.DB.Idle)

	sqlDB.SetMaxOpenConns(conf.DB.Open)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
