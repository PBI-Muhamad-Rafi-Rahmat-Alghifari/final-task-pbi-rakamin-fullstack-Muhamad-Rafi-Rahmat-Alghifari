package database

import (
	"FinalTaskAPI/models"
	"fmt"
	"os"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbPort   = os.Getenv("PGPORT")
	dbname   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	fmt.Println("Connection Opened to Database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})

}

func GetDB() *gorm.DB {
	return db
}
