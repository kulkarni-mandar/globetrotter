package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"globetrotter/pkg/database"
	"globetrotter/pkg/models"
	"time"
)

func AddUser(user *models.User) error {

	user.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	user.UserName = fmt.Sprint(time.Now().Unix())

	return database.Get().Create(user).Error
}

func GetUserID(userName string) (int, error) {
	var userId int

	err := database.Get().Model(&models.User{}).Where(&models.User{UserName: userName}).Select("id").Scan(&userId).Error
	if err != nil {
		return 0, err
	}

	if userId == 0 {
		return 0, errors.New("username does not exists")
	}

	return userId, nil
}

func GetUserName(userID int) (string, error) {
	var userName string

	err := database.Get().Model(&models.User{}).Where(&models.User{ID: userID}).Select("user_name").Scan(&userName).Error
	if err != nil {
		return "", err
	}

	if userName == "" {
		return "", errors.New("userID does not exists")
	}

	return userName, nil
}
