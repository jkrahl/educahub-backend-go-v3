package models

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
	err := DB.Where("sub = ?", sub).First(&user).Error
	return err
}
