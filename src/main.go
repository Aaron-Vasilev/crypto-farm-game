package main

import (
	"crypto-farm/db"
	"crypto-farm/src/router"
	"crypto-farm/src/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	loadEnv()
	PORT := os.Getenv("PORT")
	app := echo.New()

	db.ConnectDb()
	router.ConnectRoutes(app)

	if utils.IsProd() {
		app.Use(middleware)
	} else {
		app.Static("public/", "public/")
		fmt.Printf("Server started at localhost%s\n", PORT)
	}

	err := app.Start(PORT)
	log.Fatal("† line 33 err", err)
}

func loadEnv() {
	env := os.Getenv("ENV")

	if env == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("No .env")
		}
	}
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		return next(c)
	}
}
