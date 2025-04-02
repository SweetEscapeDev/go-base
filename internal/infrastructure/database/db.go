package database

import (
	"fmt"
	"go-base/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(c *config.Config) {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Errorf("fatal error: invalid database configuration: %w", err))
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
