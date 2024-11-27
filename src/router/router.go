package router

import (
	"crypto-farm/src/handler"
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
	app.POST("/api/pot", handler.CreatePot)
	app.POST("/api/plant", handler.PlantCoin)
}
