package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		log.Fatalln("Error opening postgres datastore.", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Could not connect to server.", err)
	}
	fmt.Println("Database connected.")
}
