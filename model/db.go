package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload" // Autoload .env variables
	_ "github.com/lib/pq"
)

// DB - The primary model database
var DB *gorm.DB

func init() {
	openDB()
	defer DB.Close()

	DB.AutoMigrate(&Profile{}, &Credentials{}, &Event{})

	fmt.Println("Database connected.")
}

func openDB() {
	var err error
	DB, err = gorm.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		log.Fatalln("Error opening postgres datastore.", err)
	}
}
