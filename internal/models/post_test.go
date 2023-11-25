package models

import (
	"testing"

	"educahub/configs"
)

func TestPost_Delete(t *testing.T) {
	configs.SetupViper()

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
