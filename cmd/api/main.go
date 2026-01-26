package main

import (
	"user-management-api/internal/app"
	"user-management-api/internal/config"
)

func main() {
	cfg := config.NewConfig()

	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}


}