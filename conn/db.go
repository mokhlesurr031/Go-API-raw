package conn

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"movie_review_apis/mgs"
	"movie_review_apis/models"
)

var db *gorm.DB

// Init creates a new connection to the database ...
func Init() {
	viper.AddConfigPath("./config")
	viper.SetConfigFile("./config/db.yaml")
	er := viper.ReadInConfig()
	if er != nil {
		log.Println(er)
	}
	DB_PORT := viper.GetInt("database.port")
	DB_HOST := viper.GetString("database.host")
	DB_USER := viper.GetString("database.user")
	DB_PASSWORD := viper.GetString("database.password")
	DB_NAME := viper.GetString("database.db")

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println("hehehhe", dbinfo)

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
