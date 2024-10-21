package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
)

func NewDB() *sql.DB {
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_DBNAME")
	)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Error("---数据库连接失败---", err.Error())
		return nil
	}

	return db
}

// func NewDB() *gorm.DB {
// 	var (
// 		user     = os.Getenv("DB_USER")
// 		password = os.Getenv("DB_PASSWORD")
// 		host     = os.Getenv("DB_HOST")
// 		port     = os.Getenv("DB_PORT")
// 		dbName   = os.Getenv("DB_DBNAME")
// 	)
// 	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
// 	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
// 	if err != nil {
// 		log.Error("---数据库连接失败---", err.Error())
// 		return nil
// 	}

// 	return db
// }
