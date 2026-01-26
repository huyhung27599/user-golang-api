package service

import "user-management-api/internal/models"

type UserService interface {
	GetAllUser() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(uuid string) error
}