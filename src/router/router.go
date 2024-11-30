package router

import (
	"crypto-farm/src/handler"
	// 	"net/http"

	"github.com/labstack/echo/v4"
)

func ConnectRoutes(app *echo.Echo) {
	// Pages
	app.GET("/", handler.Home)

	// APIs
	app.POST("/api/pot", handler.CreatePot)

	app.PUT("/api/plant", handler.HarvestCoin)
	app.POST("/api/plant", handler.PlantCoin)

	//	app.GET("/*", func(c echo.Context) error {
	//		return c.Redirect(http.StatusFound, "/")
	//	})
}
