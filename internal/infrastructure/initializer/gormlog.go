package initializer

import (
	"fmt"

	"github.com/SweetEscapeDev/gormodlog"
	"gorm.io/gorm"
)

func InitializeGORMLog(db *gorm.DB) {
	versionPlugin := gormodlog.NewDataChangeLogger()

	// versionPlugin.RegisterModelPair(db, &entity.Order{}, &entity.VersionOrder{})
	// versionPlugin.RegisterModelPair(db, &entity.User{}, &entity.VersionUser{})

	err := db.Use(versionPlugin)

	if err != nil {
		panic(fmt.Errorf("fatal error: invalid initialize gorm log: %w", err))
	}
}
