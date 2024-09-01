package services

import (
	"errors"
	"fmt"
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

	db.Joins("User").Where("user_id = ? AND start_time >= ?", userID, time.Now().Format("2006-01-02")).Last(&session)

	if session.ID == 0 {
		return "offline"
	}

	if session.EndTime == nil {
		return "working"
	}

	return "offline"
}

func UserStartSession() bool {
	var userID uint = GetCurrentUserID()

	if userID == 0 {
		return false
	}

	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return false
	}

	var session models.Session
	db.Joins("User").Where("user_id = ? AND start_time >= ? AND end_time IS NULL", userID, time.Now().Format("2006-01-02")).First(&session)

	if session.ID != 0 {
		fmt.Println("session already started")
		return false
	}

	session = models.Session{UserID: userID, StartTime: time.Now()}
	db.Create(&session)

	return true
}

func UserStopSession() bool {
	var userID uint = GetCurrentUserID()

	if userID == 0 {
		return false
	}

	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return false
	}

	var session models.Session
	db.Joins("User").Where("user_id = ? AND start_time >= ? AND end_time IS NULL", userID, time.Now().Format("2006-01-02")).First(&session)

	if session.ID == 0 {
		fmt.Println("session not found")
		return false
	}

	session.EndTime = &[]time.Time{time.Now()}[0]
	db.Save(&session)

	return true
}

func UserTotalTime(byDate string) (float64, error) {
	var userID uint = GetCurrentUserID()

	if userID == 0 {
		return 0.0, errors.New("user not found")
	}

	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return 0.0, err
	}

	var sessions []models.Session
	db.Joins("User").Where("user_id = ? AND start_time >= ?", userID, byDate).Find(&sessions)

	var totalTime time.Duration
	for _, session := range sessions {
		if session.EndTime != nil {
			totalTime += session.EndTime.Sub(session.StartTime)
		} else {
			totalTime += time.Since(session.StartTime)
		}
	}

	return totalTime.Hours(), nil
}

func UserTotalTimeToDay() (float64, error) {
	return UserTotalTime(time.Now().Format("2006-01-02"))
}
