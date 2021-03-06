package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	controller "go-clean-arch/app-controller"
	database "go-clean-arch/infrastructure/db"
	router "go-clean-arch/infrastructure/router"
	useradapter "go-clean-arch/interface/user"
	userCase "go-clean-arch/use-cases/user"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	appController := controller.NewAppController()
	db := database.New()

	userInteractor := userCase.NewUserInteractor(useradapter.NewUserInterfaceRepository(db), useradapter.NewUserPresenter())
	userController := useradapter.NewUseInterfaceController(userInteractor)

	appController.AddUserController(&userController)

	router.NewRouter(e, *appController)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
