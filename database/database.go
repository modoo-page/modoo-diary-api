package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func SetupDatabase() {
	dbHost := os.Getenv("MYSQL_HOST")
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")

	var err error
	var config gorm.Config
	config.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/modoo_diary?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost)

	DB, err = gorm.Open(mysql.Open(dsn), &config)

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}
