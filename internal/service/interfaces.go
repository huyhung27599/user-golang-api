package service

import "user-management-api/internal/models"

type UserService interface {
	GetAllUser(search string, page int, limit int) ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	UpdateUser(uuid string, user models.User) (models.User, error)
	DeleteUser(uuid string) error
}