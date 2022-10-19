package conn

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"movie_review_apis/mgs"
	"movie_review_apis/models"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5431
	DB_USER     = "postgres"
	DB_PASSWORD = "example"
	DB_NAME     = "movies"
)

var db *gorm.DB

// Init creates a new connection to the database ...
func Init() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	db, err = gorm.Open(postgres.Open(dbinfo))

	if err != nil {
		log.Println("Failed to connect to database")
		mgs.CheckErr(err)
	}
	log.Println("Database connected")
	db.AutoMigrate(models.Movie{})
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

// CloseDB ...
func CloseDB() {
	db, _ := db.DB()
	db.Close()
}
