package database

import (
	"mini-clickup/internal/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("miniclickup.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{}, &models.TaskLog{})
	return db, nil
}
