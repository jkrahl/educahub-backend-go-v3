package models

import (
	"errors"
	"log"
	"time"
)

type Community struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	Name      string    `json:"name" gorm:"unique;not null"`
	URL       string    `json:"url" gorm:"unique;not null"`
	Posts     []Post    `json:"posts" gorm:"foreignkey:CommunityID"`
	CreatedAt time.Time `json:"created_at"`
}

type CommunityResponse struct {
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Community) Create() error {
	err := GetDB().Create(c).Error
	if err != nil {
		log.Println("Create community error: ", err.Error())
		return errors.New("community not created")
	}
	return nil
}

func (c *Community) Find() error {
	err := GetDB().Where("url = ?", c.URL).First(&c).Error
	if err != nil {
		log.Println("Find community error: ", err.Error())
		return errors.New("community not found")
	}
	return nil
}

func (c *Community) GetAllPosts() ([]Post, error) {
	err := GetDB().Preload("User").Preload("Community").Where("community_id = ?", c.ID).Find(&c.Posts).Error
	if err != nil {
		log.Println("Get all posts error: ", err.Error())
		return nil, errors.New("posts not found")
	}
	return c.Posts, nil
}

// This function gets all communities from the database.
// The communities do not have any posts attached to them.
func GetAllCommunities() ([]Community, error) {
	var communities []Community

	err := GetDB().Find(&communities).Error
	if err != nil {
		log.Println("Get all communities error: ", err.Error())
		return nil, errors.New("communities not found")
	}
	return communities, nil
}

func CommunityToCommunityResponse(c *Community) CommunityResponse {
	return CommunityResponse{
		Name:      c.Name,
		URL:       c.URL,
		CreatedAt: c.CreatedAt,
	}
}
