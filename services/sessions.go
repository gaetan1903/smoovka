package services

import (
	"os"
	"smoovka/models"
	"smoovka/utils"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetCurrentUserID() uint {
	token := utils.GetPreference("token")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return 0
	}
	return uint(claims["userID"].(float64))
}

func GetCurrentStatus() string {
	var userID uint = GetCurrentUserID()

	if userID == 0 {
		return "offline"
	}

	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return "offline"
	}
	var session models.Session

	db.Joins("User").Where("user_id = ? AND start_time >= ?", userID, time.Now().Format("2006-01-02")).First(&session)

	if session.ID == 0 {
		return "offline"
	}

	if session.EndTime == nil {
		return "working"
	}

	return "offline"

	// get all sessions of user to
	// if login == "" {
	// 	return "offline"
	// }

	// db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	// if err != nil {
	// 	return "offline"
	// }
	// var session models.Session
	// // get all sessions of user today
	// db.Where("login = ? AND start_time > ?", login, time.Now().Format("2006-01-02")).First(&session)
	// if session.ID == 0 {
	// 	return "offline"
	// }
	// if session.EndTime == nil {
	// 	return "working"
	// }
	// return "offline"
}
