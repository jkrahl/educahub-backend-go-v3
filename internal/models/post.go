package models

import (
	"errors"
	"log"
	"time"
)

const (
	PostTypeNotes    = 1
	PostTypeQuestion = 2
)

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	Type      int       `json:"type" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user"`
	URL       string    `json:"url" gorm:"unique;not null"`
	Subject   string    `json:"subject"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
}

type PostResponse struct {
	Type      int       `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	URL       string    `json:"url"`
	Subject   string    `json:"subject"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Post) Create() error {
	err := GetDB().Create(p).Error
	if err != nil {
		log.Println("Create post error: ", err.Error())
		return errors.New("post not created")
	}
	return err
}

func (p *Post) Delete() error {
	err := p.Find()
	if err != nil {
		return err
	}
	err = GetDB().Delete(p).Error
	if err != nil {
		log.Println("Delete post error: ", err.Error())
		return errors.New("post not deleted")
	}
	return err
}

func (p *Post) Update() error {
	err := GetDB().Save(p).Error
	if err != nil {
		log.Println("Update post error: ", err.Error())
		return errors.New("post not updated")
	}
	return err
}

func (p *Post) Find() error {
	err := GetDB().Where(p).First(p).Error
	if err != nil {
		log.Println("Find post error: ", err.Error())
		return errors.New("post not found")
	}
	return nil
}

func GetAllPosts(posts *[]Post) error {
	err := GetDB().Find(posts).Error
	if err != nil {
		log.Println("Get all posts error: ", err.Error())
		return errors.New("posts not found")
	}
	return nil
}

func PostToPostResponse(post *Post) *PostResponse {
	user := User{
		ID: post.UserID,
	}
	err := user.Find()
	if err != nil {
		log.Println("Error finding user: ", err.Error())
		return nil
	}

	return &PostResponse{
		Type:      post.Type,
		Title:     post.Title,
		Content:   post.Content,
		Username:  user.Username,
		URL:       post.URL,
		Subject:   post.Subject,
		Unit:      post.Unit,
		CreatedAt: post.CreatedAt,
	}
}
