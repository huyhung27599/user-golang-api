package repository

import "user-management-api/internal/models"

type UserRepository interface {
	FindAllUser() ([]models.User, error)
	FindUserByUUID(uuid string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(uuid string) error
}