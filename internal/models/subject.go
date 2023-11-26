package models

import (
	"errors"
	"log"
	"time"
)

type Subject struct {
	ID          uint      `json:"id" gorm:"primary_key;not null"`
	Name        string    `json:"name" gorm:"unique;not null"`
	URL         string    `json:"url" gorm:"unique;not null"`
	CommunityID uint      `json:"community_id" gorm:"not null"`
	Community   Community `json:"community"`
	Posts       []Post    `json:"posts" gorm:"foreignkey:SubjectID"`
	CreatedAt   time.Time `json:"created_at"`
}

type SubjectResponse struct {
	Name         string    `json:"name"`
	URL          string    `json:"url"`
	Community    string    `json:"community"`
	CommunityURL string    `json:"community_url"`
	CreatedAt    time.Time `json:"created_at"`
}

func (s *Subject) Create() error {
	err := GetDB().Create(s).Error
	if err != nil {
		log.Println("Create subject error: ", err.Error())
		return errors.New("subject not created")
	}
	return nil
}

func (s *Subject) GetAllPosts() ([]Post, error) {
	err := GetDB().Preload("User").Preload("Community").Where("subject_id = ?", s.ID).Find(&s.Posts).Error
	if err != nil {
		log.Println("Get all posts error: ", err.Error())
		return nil, errors.New("posts not found")
	}
	return s.Posts, nil
}

func (s *Subject) FindByURL() error {
	err := GetDB().Preload("Community").Where("url = ?", s.URL).First(&s).Error
	if err != nil {
		log.Println("Find subject error: ", err.Error())
		return errors.New("subject not found")
	}
	return nil
}

func (s *Subject) FindByID() error {
	err := GetDB().Where("id = ?", s.ID).First(&s).Error
	if err != nil {
		log.Println("Find subject error: ", err.Error())
		return errors.New("subject not found")
	}
	return nil
}

func SubjectToSubjectResponse(subject *Subject) SubjectResponse {
	return SubjectResponse{
		Name:         subject.Name,
		URL:          subject.URL,
		Community:    subject.Community.Name,
		CommunityURL: subject.Community.URL,
		CreatedAt:    subject.CreatedAt,
	}
}
