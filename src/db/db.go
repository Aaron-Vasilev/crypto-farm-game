package db

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDb() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to db", err)
	} else {
		DB = db
	}
	defer db.Close()
}
