package app

import (
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/service"
)

type UserModule struct{
	userRoutes routes.Route
}

func NewUserModule() *UserModule {

	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)
	return &UserModule{userRoutes: userRoutes}
}

func (m *UserModule) Routes() routes.Route {
	return m.userRoutes
}