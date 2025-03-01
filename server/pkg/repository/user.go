package repository

import (
	"fmt"
	"globetrotter/pkg/database"
	"globetrotter/pkg/models"
	"time"
)

func AddUser(user *models.User) error {

	user.CreatedAt = time.Now()
	user.UserName = fmt.Sprint(time.Now().Unix())

	return database.Get().Create(user).Error
}
