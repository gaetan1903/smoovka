package services

import (
	"os"
	"smoovka/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDb() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Pause{})

	return db, nil
}
