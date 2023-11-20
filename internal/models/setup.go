package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		ConnectDatabase()
	}
	return db
}

func ConnectDatabase() {
	var err error
	db, err = gorm.Open(
		mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true&interpolateParams=true",
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_NAME"),
		)),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migrate models
	err = db.AutoMigrate(&User{}, &Post{}, &Tag{})
	if err != nil {
		panic("Failed to migrate database!")
	}
}
