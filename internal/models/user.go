package models

import (
	"errors"
	"log"
	"time"
)

type Tag struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"unique;not null"`
}
type User struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Sub       string    `json:"sub" gorm:"unique;not null"`
	Tags      []Tag     `json:"tags" gorm:"many2many:user_tags;"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) Find() error {
	err := GetDB().Where(user).First(user).Error
	if err != nil {
		log.Println("FindUser: ", err)
		return errors.New("user not found")
	}
	return nil
}

func (user *User) Register() error {
	err := GetDB().Create(&user).Error
	if err != nil {
		log.Println("Register: ", err)
		return errors.New("user already exists")
	}
	return nil
}

func (user *User) Delete() error {
	err := user.Find()
	if err != nil {
		return err
	}

	err = GetDB().Delete(user).Error
	if err != nil {
		log.Println("DeleteUser: ", err)
		return errors.New("internal server error")
	}
	return nil
}
