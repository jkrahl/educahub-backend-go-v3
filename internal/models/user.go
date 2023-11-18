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
