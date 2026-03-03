package dto

import "user-management-api/internal/models"

type UserDTO struct {
	UUID   string `json:"uuid" binding:"required,uuid"`
	Name   string `json:"full_name" binding:"required,min=3,max=100"`
	Age    int    `json:"age" binding:"required,gt=0,lt=100"`
	Email  string `json:"email_address" binding:"required,email"`
	Status string `json:"status" binding:"required,oneof=1 2"`
	Level  string `json:"level" binding:"required,oneof=1 2 "`
}

type CreateUserInput struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Age int `json:"age" binding:"required,gt=0,lt=100"`
	Email string `json:"email" binding:"required,email,email_address"`
	Password string `json:"password" binding:"required,min=8,max=20,password_strong"`
	Status int `json:"status" binding:"required,oneof=1 2"`
	Level int `json:"level" binding:"required,oneof=1 2 "`
}



type UpdateUserInput struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Age int `json:"age" binding:"required,gt=0,lt=100"`
	Email string `json:"email" binding:"required,email,email_address"`
	Password string `json:"password" binding:"omitempty,min=8,max=20,password_strong"`
	Status int `json:"status" binding:"required,oneof=1 2"`
	Level int `json:"level" binding:"required,oneof=1 2 "`
}

func (input *CreateUserInput) MapCreateUserInputToUserModel() models.User {
	return models.User{
		Name: input.Name,
		Age: input.Age,
		Email: input.Email,
		Password: input.Password,
		Status: input.Status,
		Level: input.Level,
	}
}

func (input *UpdateUserInput) MapUpdateUserInputToUserModel() models.User {
	return models.User{
		Name: input.Name,
		Age: input.Age,
		Email: input.Email,
		Password: input.Password,
		Status: input.Status,
		Level: input.Level,
	}
}
func MapUserToDTO(user models.User) *UserDTO {
	return &UserDTO{
		UUID: user.UUID,
		Name: user.Name,
		Age: user.Age,
		Email: user.Email,
		Status: mapStatusToText(user.Status),
		Level: mapLevelToText(user.Level),
	}
}

func MapUserToDTOList(users []models.User) []*UserDTO {

dtos := make([]*UserDTO, 0, len(users))
for _, user := range users {
	dtos = append(dtos, MapUserToDTO(user))
}
return dtos
}

func mapStatusToText(status int) string {
	switch status {
	case 1:
		return "active"
	case 2:
		return "inactive"
	default:
		return "unknown"
	}
}

func mapLevelToText(level int) string {
	switch level {
	case 1:
		return "admin"
	case 2:
		return "user"
	default:
		return "unknown"
	}
}
