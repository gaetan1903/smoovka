package services

import (
	"fmt"
	"log"
	"os"
	"smoovka/models"
	"smoovka/utils"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func UserLogin(login string, password string) bool {
	var user models.User

	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return false
	}
	db.Where("login = ? AND password = ?", login, password).First(&user)

	if user.ID == 0 {
		err = gorm.ErrRecordNotFound
	} else {
		userToken := utils.CreateToken(user.ID)
		utils.SetPreference("token", userToken)
	}

	return err == nil
}

func UserIsLoggedIn() bool {
	token := utils.GetPreference("token")
	log.Println("token", token)
	if token == "" {
		return false
	}

	_, err := utils.VerifyToken(token)

	return err == nil
}

func UserGetAccount() map[string]interface{} {
	var user models.User
	token := utils.GetPreference("token")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return nil
	}

	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return nil
	}
	db.Where("id = ?", claims["userID"]).First(&user)

	if user.ID == 0 {
		return nil
	}

	return map[string]interface{}{
		"id":          user.ID,
		"displayName": user.Name,
		"email":       user.Email,
		"photoURL": fmt.Sprintf(
			"/assets/images/avatars/avatar_%d.jpg", user.ID%25+1,
		),
	}

}
