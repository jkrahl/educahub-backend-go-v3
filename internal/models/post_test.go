package models

import (
	"reflect"
	"testing"

	"educahub/configs"
)

func TestPost_DeletePost(t *testing.T) {
	configs.SetupViper("../../configs")

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
			if err := tt.p.DeletePost(); (err != nil) != tt.wantErr {
				t.Errorf("Post.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPostByURL(t *testing.T) {
	configs.SetupViper("../../configs")

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    Post
		wantErr bool
	}{
		{
			"GetNonExistentPost",
			args{"non-existent-post"},
			Post{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPostByURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPostByURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPostByURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
