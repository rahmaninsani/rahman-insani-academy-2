package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-clean-arch/config"
	"go-echo-clean-arch/handler"
	"go-echo-clean-arch/repository"
	"go-echo-clean-arch/routes"
	"go-echo-clean-arch/usecase"
)

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.Routes(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
