package conn

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"movie_review_apis/mgs"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5431
	DB_USER     = "postgres"
	DB_PASSWORD = "example"
	DB_NAME     = "movies"
)

//func SetupDB() *sql.DB {
//	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
//	db, err := sql.Open("postgres", dbinfo)
//	mgs.CheckErr(err)
//	return db
//}

func SetupDB() *gorm.DB {
	// https://github.com/jackc/pgx
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	mgs.CheckErr(err)
	return db
}
