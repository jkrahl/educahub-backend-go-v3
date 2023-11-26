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

// consts for errors
var (
	ErrPostNotFound = errors.New("post not found")
)

type Post struct {
	ID          uint      `json:"id" gorm:"primary_key;not null"`
	Type        int       `json:"type" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user"`
	CommunityID uint      `json:"community_id" gorm:"not null"`
	Community   Community `json:"community"`
	URL         string    `json:"url" gorm:"unique;not null"`
	Subject     string    `json:"subject"`
	Unit        string    `json:"unit"`
	Comments    []Comment `json:"comments" gorm:"foreignkey:PostID"`
	CreatedAt   time.Time `json:"created_at"`
}

type PostResponse struct {
	Type         int       `json:"type"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Username     string    `json:"username"`
	Community    string    `json:"community"`
	CommunityURL string    `json:"community_url"`
	URL          string    `json:"url"`
	Subject      string    `json:"subject"`
	Unit         string    `json:"unit"`
	CreatedAt    time.Time `json:"created_at"`
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
	err := p.FindByURL()
	if err != nil {
		return ErrPostNotFound
	}
	err = GetDB().Delete(p).Error
	if err != nil {
		log.Println("Delete post error: ", err.Error())
		return errors.New("post not deleted")
	}
	return err
}

func (p *Post) Update() error {
	err := p.FindByURL()
	if err != nil {
		return ErrPostNotFound
	}
	err = GetDB().Save(p).Error
	if err != nil {
		log.Println("Update post error: ", err.Error())
		return errors.New("post not updated")
	}
	return err
}

func (p *Post) FindByURL() error {
	err := GetDB().Preload("User").Preload("Community").Where("url = ?", p.URL).First(&p).Error
	if err != nil {
		log.Println("Find post error: ", err.Error())
		return ErrPostNotFound
	}
	return nil
}

func (p *Post) GetAllComments() ([]Comment, error) {
	err := GetDB().Preload("Post").Preload("User").Where("post_id = ?", p.ID).Find(&p.Comments).Error
	if err != nil {
		log.Println("Get all comments error: ", err.Error())
		return nil, errors.New("comments not found")
	}
	return p.Comments, nil
}

func PostToPostResponse(post *Post) PostResponse {
	return PostResponse{
		Type:         post.Type,
		Title:        post.Title,
		Content:      post.Content,
		Username:     post.User.Username,
		Community:    post.Community.Name,
		CommunityURL: post.Community.URL,
		URL:          post.URL,
		Subject:      post.Subject,
		Unit:         post.Unit,
		CreatedAt:    post.CreatedAt,
	}
}
