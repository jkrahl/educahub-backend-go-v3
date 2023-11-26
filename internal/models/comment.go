package models

import (
	"errors"
	"log"
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	UUID      string    `json:"uuid" gorm:"unique;not null"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user"`
	PostID    uint      `json:"post_id" gorm:"not null"`
	Post      Post      `json:"post"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	UUID      string    `json:"uuid"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	PostUUID  string    `json:"post_uuid"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Comment) Create() error {
	err := GetDB().Create(c).Error
	if err != nil {
		log.Println("Create comment error: ", err.Error())
		return errors.New("comment not created")
	}
	return err
}

func (c *Comment) Delete() error {
	err := c.Find()
	if err != nil {
		return errors.New("comment not found")
	}
	err = GetDB().Delete(c).Error
	if err != nil {
		log.Println("Delete comment error: ", err.Error())
		return errors.New("comment not deleted")
	}
	return err
}

func (c *Comment) Update() error {
	err := c.Find()
	if err != nil {
		return errors.New("comment not found")
	}
	err = GetDB().Save(c).Error
	if err != nil {
		log.Println("Update comment error: ", err.Error())
		return errors.New("comment not updated")
	}
	return err
}

func (c *Comment) Find() error {
	err := GetDB().Preload("User").Preload("Post").Where("uuid = ?", c.UUID).First(c).Error
	if err != nil {
		log.Println("Find comment error: ", err.Error())
		return errors.New("comment not found")
	}
	return nil
}

func CommentToCommentResponse(c *Comment) *CommentResponse {
	return &CommentResponse{
		UUID:      c.UUID,
		Content:   c.Content,
		Username:  c.User.Username,
		PostUUID:  c.Post.URL,
		CreatedAt: c.CreatedAt,
	}
}
