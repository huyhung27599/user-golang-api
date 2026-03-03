package service

import (
	"strings"
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

func (us *userService) GetAllUser(search string, page int, limit int) ([]models.User, error) {
	users, err := us.repo.FindAllUser()
	if err != nil {
		return nil, utils.WrapError("failed to get all users", utils.ErrCodeInternal, err)
	}

	var filteredUsers []models.User
	if search != "" {
		search = utils.NormalizeString(search)
		for _, user := range users {
			if strings.Contains(strings.ToLower(user.Name), search) || strings.Contains(strings.ToLower(user.Email), search) {
				filteredUsers = append(filteredUsers, user)
			}
		}
	} else {
		filteredUsers = users
	}

	start := (page - 1) * limit
	end := start + limit
	if start >= len(filteredUsers) {
		return []models.User{}, nil
	}

	if end > len(filteredUsers) {
		end = len(filteredUsers)
	}
	filteredUsers = filteredUsers[start:end]

	return filteredUsers, nil
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

func (us *userService) UpdateUser(uuid string, user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)
 if u, ok := us.repo.FindByEmail(user.Email); ok && u.UUID != uuid {
	return models.User{}, utils.WrapError("email already exists", utils.ErrCodeConflict, nil)
 }

 currentUser, found := us.repo.FindUserByUUID(uuid)
 if !found { 
	return models.User{}, utils.WrapError("user not found", utils.ErrCodeNotFound, nil)
 }
 currentUser.Name = user.Name
 currentUser.Age = user.Age
 currentUser.Email = user.Email
 currentUser.Status = user.Status
 currentUser.Level = user.Level

 if user.Password != "" {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError("failed to hash password", utils.ErrCodeInternal, err)
	}
	currentUser.Password = string(hashedPassword)
 }

  	err := us.repo.UpdateUser(uuid, currentUser)
	if err != nil {
		return models.User{}, utils.WrapError("failed to update user", utils.ErrCodeInternal, err)
	}
	return currentUser, nil
}

func (us *userService) DeleteUser(uuid string) error {
	err := us.repo.DeleteUser(uuid)
	if err != nil {
		return utils.WrapError("failed to delete user", utils.ErrCodeInternal, err)
	}
	return nil
}