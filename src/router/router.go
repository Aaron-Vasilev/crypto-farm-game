package router

import (
	handler "crypto-farm/src/hander"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConnectRoutes(app *echo.Echo) {
	// Pages
	app.GET("/", handler.Home)
	app.GET("/*", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/")
	})

	// APIs
}
