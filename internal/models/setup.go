package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"educahub/configs"
)

var db *gorm.DB

func init() {
	ConnectDatabase()
}

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
			configs.GetViperString("DB_USER"),
			configs.GetViperString("DB_PASSWORD"),
			configs.GetViperString("DB_HOST"),
			configs.GetViperString("DB_PORT"),
			configs.GetViperString("DB_NAME"),
		)),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migrate models
	err = db.AutoMigrate(&User{}, &Post{}, &Tag{}, &Comment{}, &Community{}, &Subject{})
	if err != nil {
		panic("Failed to migrate database!")
	}
}
