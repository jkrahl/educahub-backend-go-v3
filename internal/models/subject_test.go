package models

import (
	"testing"
	"time"
)

func TestSubject_Create(t *testing.T) {
	community := Community{
		URL: "test",
	}
	community.Find()

	type fields struct {
		ID          uint
		Name        string
		URL         string
		CommunityID uint
		Community   Community
		Posts       []Post
		CreatedAt   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Create subject",
			fields: fields{
				Name:        "Test subject",
				URL:         "test-subject",
				CommunityID: community.ID,
				Community:   community,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				URL:         tt.fields.URL,
				CommunityID: tt.fields.CommunityID,
				Community:   tt.fields.Community,
				Posts:       tt.fields.Posts,
				CreatedAt:   tt.fields.CreatedAt,
			}
			if err := s.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Subject.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
