package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload" // Autoload .env variables
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		log.Fatalln("Error opening postgres datastore.", err)
	}

	DB.AutoMigrate(&User{}, &Credentials{})

	fmt.Println("Database connected.")
}
