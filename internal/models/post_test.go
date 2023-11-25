package models

import (
	"educahub/configs"
	"testing"
)

func TestPost_Delete(t *testing.T) {
	tests := []struct {
		name    string
		p       *Post
		wantErr bool
	}{
		{
			"DeleteNonExistentPost",
			&Post{URL: "non-existent-post"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Delete(); (err != nil) != tt.wantErr {
				t.Errorf("Post.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPost_GetAllComments(t *testing.T) {
	type args struct {
		comments *[]Comment
	}
	tests := []struct {
		name    string
		post    Post
		args    args
		wantErr bool
	}{
		{
			"NonExistentPost",
			Post{
				URL: "non-existent-post",
			},
			args{comments: &[]Comment{}},
			true,
		},
		{
			"PostWithNoComments",
			Post{
				URL: configs.GetViperString("TEST_POST_URL"),
			},
			args{comments: &[]Comment{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				ID:        tt.post.ID,
				Type:      tt.post.Type,
				Title:     tt.post.Title,
				Content:   tt.post.Content,
				UserID:    tt.post.UserID,
				User:      tt.post.User,
				URL:       tt.post.URL,
				Subject:   tt.post.Subject,
				Unit:      tt.post.Unit,
				CreatedAt: tt.post.CreatedAt,
			}
			if err := p.GetAllComments(tt.args.comments); (err != nil) != tt.wantErr {
				t.Errorf("Post.GetAllComments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
