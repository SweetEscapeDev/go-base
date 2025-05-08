package initializer

import (
	"fmt"
	"go-base/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDatabase(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Info),
		FullSaveAssociations: false,
	})

	if err != nil {
		panic(fmt.Errorf("fatal error: invalid database configuration: %w", err))
	}

	return db
}
