package models

import "errors"

type Tag struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"unique_index"`
}
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique_index"`
	Sub      string `json:"sub" gorm:"unique_index"`
	Tags     []Tag  `json:"tags" gorm:"many2many:user_tags;"`
}

// This function is used to get the user from a given sub
func (user *User) GetUserFromSub(sub string) error {
	err := GetDBInstance().Where("sub = ?", sub).First(&user).Error
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (user *User) RegisterUser(sub string, username string) error {
	userAlreadyExists := User{}
	err := GetDBInstance().Where("sub = ?", sub).First(&userAlreadyExists).Error
	if err == nil {
		return errors.New("user already exists")
	}
	user.Sub = sub
	user.Username = username
	err = GetDBInstance().Create(&user).Error
	return err
}

func (user *User) DeleteUser(sub string) error {
	err := GetDBInstance().Where("sub = ?", sub).First(&user).Error
	if err != nil {
		return errors.New("user does not exist")
	}
	err = GetDBInstance().Delete(&user).Error
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}
