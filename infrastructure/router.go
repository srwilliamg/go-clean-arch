package infrastructure

import (
	controller "go-clean-arch/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, c controller.AppController) {

	// Routes
	e.GET("/", func(context echo.Context) error { return c.User.GetByName(context) })
}
