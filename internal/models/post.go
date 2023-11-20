package models

import "time"

type PostType uint

const (
	PostTypeNotes    PostType = 1
	PostTypeQuestion PostType = 2
	PostTypeAnswer   PostType = 3
)

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	Type      PostType  `json:"type" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user"`
	URL       string    `json:"url" gorm:"unique;not null"`
	Subject   string    `json:"subject"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Post) CreatePost() error {
	result := GetDB().Create(p)
	return result.Error
}

func (p *Post) DeletePost() error {
	result := GetDB().Delete(p)
	return result.Error
}

func GetPostByURL(url string) (Post, error) {
	var post Post
	result := GetDB().Where("url = ?", url).First(&post)
	return post, result.Error
}
