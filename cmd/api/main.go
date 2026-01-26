package main

import (
	"user-management-api/internal/config"
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes( userHandler)

	r := gin.Default()

	routes.RegisterRoutes(r, userRoutes)
}