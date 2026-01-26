package repository

import "user-management-api/internal/models"


type InMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func (ur *InMemoryUserRepository) FindAllUser() ([]models.User, error) {}


func (ur *InMemoryUserRepository) FindUserByUUID(uuid string) (models.User, error) {}

func (ur *InMemoryUserRepository) CreateUser(user models.User) (models.User, error) {}

func (ur *InMemoryUserRepository) UpdateUser(user models.User) (models.User, error) {}

func (ur *InMemoryUserRepository) DeleteUser(uuid string) error {}