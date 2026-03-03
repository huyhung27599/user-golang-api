package repository

import (
	"slices"
	"user-management-api/internal/models"
	"user-management-api/internal/utils"
)


type InMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func (ur *InMemoryUserRepository) FindAllUser() ([]models.User, error) {
	return ur.users, nil
}


func (ur *InMemoryUserRepository) FindUserByUUID(uuid string) (models.User,bool) {
	for _, user := range ur.users {
		if user.UUID == uuid {
			return user, true
		}
	}
	return models.User{}, false
}	

func (ur *InMemoryUserRepository) CreateUser(user models.User) (models.User, error) {
	ur.users = append(ur.users, user)
	return user, nil
}

func (ur *InMemoryUserRepository) UpdateUser(uuid string, user models.User)  error {
	for i, u := range ur.users {
		if u.UUID == uuid {
			ur.users[i] = user
			return nil
		}
	}
	return utils.WrapError("user not found", utils.ErrCodeNotFound, nil)
}

func (ur *InMemoryUserRepository) DeleteUser(uuid string) error {
	for i, u := range ur.users {
		if u.UUID == uuid {
			ur.users = slices.Delete(ur.users, i, 1)
			return nil
		}
	}
	return utils.WrapError("user not found", utils.ErrCodeNotFound, nil)
}

func (ur *InMemoryUserRepository) FindByEmail(email string) (models.User,bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}
	return models.User{}, false
}