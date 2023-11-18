package models

import (
	"fmt"

	viper "github.com/spf13/viper"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(psql.Open(
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_USER"),
			viper.GetString("DB_NAME"),
			viper.GetString("DB_PASSWORD"),
		)),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}
