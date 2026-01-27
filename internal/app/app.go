package app

import (
	"log"
	"user-management-api/internal/config"
	"user-management-api/internal/routes"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
	modules []Module
}

func NewApplication(config *config.Config) *Application {
	loadEnv()
	router := gin.Default()

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Error initializing validator: %v", err)
	}

	modules := []Module{
		NewUserModule(),
	}


	routes.RegisterRoutes(router, getModulRoutes(modules)...)
	return &Application{
		config: config,
		router: router,
		modules: modules,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModulRoutes(modules []Module) []routes.Route {
 routeList := make([]routes.Route, len(modules))
 for i, module := range modules {
	routeList[i] = module.Routes()
 }
 return routeList
}

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}