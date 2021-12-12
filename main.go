package main

import (
	"majoo-test-case/config"

	"majoo-test-case/entity/user/delivery"
	userRepo "majoo-test-case/entity/user/repository"
	"majoo-test-case/entity/user/usecase"

	omzetDelivery "majoo-test-case/entity/omzet/delivery"
	omzetRepo "majoo-test-case/entity/omzet/repository"
	omzetUsecase "majoo-test-case/entity/omzet/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	config.GetEnvVariable()
	dbConn := config.InitDatabase()
	e := echo.New()

	repo := userRepo.NewMySQLUserRepository(dbConn)
	usecase := usecase.NewUserUseCase(repo)
	delivery.NewHttpDelivery(e, usecase)

	omzetRepository := omzetRepo.NewMySQLOmzetRepository(dbConn)
	omzetUsecaseInstance := omzetUsecase.NewOmzetUseCase(omzetRepository)
	omzetDelivery.NewHttpDelivery(e, omzetUsecaseInstance, repo)

	e.Start(":8000")
}
