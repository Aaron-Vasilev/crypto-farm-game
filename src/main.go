package main

import (
	"crypto-farm/src/auth"
	"crypto-farm/src/db"
	"crypto-farm/src/router"
	"crypto-farm/src/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	loadEnv()
	PORT := os.Getenv("PORT")
	app := echo.New()

	db.ConnectDb()
	defer db.DB.Close()
	router.ConnectRoutes(app)
	app.Use(auth.Middleware)

	if !utils.IsProd() {
		app.Use(middleware.Logger())
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
