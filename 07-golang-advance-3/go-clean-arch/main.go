package main

import (
	"github.com/labstack/echo/v4"
	"go-clean-arch/config"
	"go-clean-arch/handler"
	"go-clean-arch/repository"
	"go-clean-arch/routes"
	"go-clean-arch/usecase"
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
