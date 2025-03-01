package service

import (
	"globetrotter/pkg/models"
	"globetrotter/pkg/repository"
)

func AddNewUser(user *models.User) error {
	return repository.AddUser(user)
}
