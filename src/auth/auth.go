package auth

import (
	"crypto-farm/src/consts"
	"crypto-farm/src/utils"
	"log"

	"github.com/labstack/echo/v4"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if utils.IsProd() {
			//TODO get user from request and validate him
		} else {
			c.Set(consts.UserID, int64(1))
		}
		return next(c)
	}
}

func GetUserIDFromCtx(c echo.Context) int64 {
	userId := c.Get(consts.UserID)

	if id, ok := userId.(int64); ok {
		return id
	} else {
		log.Fatalf("No user in the context")
		return 0
	}
}
