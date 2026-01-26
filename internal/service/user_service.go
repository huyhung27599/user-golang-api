package service

import (
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) GetAllUser() ([]models.User, error) {
	return us.repo.FindAllUser()
}

func (us *userService) CreateUser(user models.User) (models.User, error) {}

func (us *userService) GetUserByUUID(uuid string) (models.User, error) {}

func (us *userService) UpdateUser(user models.User) (models.User, error) {}

func (us *userService) DeleteUser(uuid string) error {}