package service

import (
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) GetAllUser() ([]models.User, error) {
	users, err := us.repo.FindAllUser()
	if err != nil {
		return nil, utils.WrapError("failed to get all users", utils.ErrCodeInternal, err)
	}
	return users, nil
}

func (us *userService) CreateUser(user models.User) (models.User, error) {
 user.Email = utils.NormalizeString(user.Email)
 if _, ok := us.repo.FindByEmail(user.Email); ok {
	return models.User{}, utils.WrapError("email already exists", utils.ErrCodeConflict, nil)
 }
 user.UUID = uuid.New().String()

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
	return models.User{}, utils.WrapError("failed to hash password", utils.ErrCodeInternal, err)
  }
  user.Password = string(hashedPassword)

  user, err = us.repo.CreateUser(user)
  if err != nil {
	return models.User{}, utils.WrapError("failed to create user", utils.ErrCodeInternal, err)
  }
  return user, nil
}

func (us *userService) GetUserByUUID(uuid string) (models.User, error) {
	user, found := us.repo.FindUserByUUID(uuid)
	if !found {
		return models.User{}, utils.WrapError("user not found", utils.ErrCodeNotFound, nil)
	}
	return user, nil
}

func (us *userService) UpdateUser(user models.User) (models.User, error) {}

func (us *userService) DeleteUser(uuid string) error {}