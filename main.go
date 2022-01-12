package main

import (
	"os"
	accountC "sample-api/account/controller"
	accountR "sample-api/account/repository"
	accountU "sample-api/account/usecase"
	authC "sample-api/auth/controller"
	authU "sample-api/auth/usecase"
	"sample-api/database"
	"sample-api/database/migrations"
	"sample-api/utils"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// DB接続
	db := database.ConnectDB()
	// migration
	migrations.Execute()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// mount token handler
	token := utils.NewTokenHandler()

	accountRepo := accountR.NewAccountRepository(db)

	accountUcase := accountU.NewAccountUsecase(accountRepo)
	authUcase := authU.NewAuthUsecase(accountRepo, token)

	accountC.NewAccountController(e, accountUcase, token)
	authC.NewAuthController(e, authUcase)

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
