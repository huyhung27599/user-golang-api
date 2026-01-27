package repository

import (
	"user-management-api/internal/models"
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


func (ur *InMemoryUserRepository) FindUserByUUID(uuid string) (models.User, error) {}

func (ur *InMemoryUserRepository) CreateUser(user models.User) (models.User, error) {
	ur.users = append(ur.users, user)
	return user, nil
}

func (ur *InMemoryUserRepository) UpdateUser(user models.User) (models.User, error) {}

func (ur *InMemoryUserRepository) DeleteUser(uuid string) error {}

func (ur *InMemoryUserRepository) FindByEmail(email string) (models.User,bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}
	return models.User{}, false
}