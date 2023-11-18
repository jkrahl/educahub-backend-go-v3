package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDBInstance() *gorm.DB {
	if DB == nil {
		ConnectDatabase()
	}

	return DB
}

func ConnectDatabase() {
	database, err := gorm.Open(
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

	err = database.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}
