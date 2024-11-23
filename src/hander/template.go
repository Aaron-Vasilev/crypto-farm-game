package handler

import (
	pages "crypto-farm/src/pages/home"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	return pages.Home(pages.HomeProps{
	}).Render(c.Request().Context(), c.Response())
}
