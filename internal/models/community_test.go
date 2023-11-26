package models

import (
	"reflect"
	"testing"
	"time"
)

func TestCommunity_Create(t *testing.T) {
	type fields struct {
		Name      string
		URL       string
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Create community",
			fields:  fields{Name: "test", URL: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				Name:      tt.fields.Name,
				URL:       tt.fields.URL,
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := c.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Community.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_GetAllPosts(t *testing.T) {
	type fields struct {
		ID        uint
		Name      string
		URL       string
		Posts     []Post
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Post
		wantErr bool
	}{
		{
			name:    "Get all posts",
			fields:  fields{Name: "test", URL: "test"},
			want:    []Post{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				URL:       tt.fields.URL,
				Posts:     tt.fields.Posts,
				CreatedAt: tt.fields.CreatedAt,
			}
			got, err := c.GetAllPosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Community.GetAllPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Community.GetAllPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
