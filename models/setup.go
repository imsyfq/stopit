package models

import (
	"stopit/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	username := config.Env("DB_USERNAME")
	password := config.Env("DB_PASSWORD")
	host := config.Env("DB_HOST")
	port := config.Env("DB_PORT")
	name := config.Env("DB_NAME")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Action{}, &User{})
	// database.Model(&Action{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	DB = database
}
