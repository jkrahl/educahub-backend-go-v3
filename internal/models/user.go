package models

import (
	"errors"
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

func (user *User) GetUserFromSub(sub string) error {
	err := GetDB().Where("sub = ?", sub).First(&user).Error
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (user *User) RegisterUser() error {
	userAlreadyExists := User{}
	err := GetDB().Where(&User{Sub: user.Sub}).Or(&User{Username: user.Username}).First(&userAlreadyExists).Error
	if err == nil {
		return errors.New("user already exists")
	}
	err = GetDB().Create(&user).Error
	return err
}

func DeleteUser(user *User) error {
	err := GetDB().Where(&User{Sub: user.Sub}).First(user).Error
	if err != nil {
		return errors.New("user does not exist")
	}
	err = GetDB().Delete(user).Error
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}
