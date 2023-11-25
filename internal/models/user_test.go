package models

import (
	"educahub/configs"
	"testing"
)

func TestUser_Register(t *testing.T) {
	configs.SetupViper()

	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		{
			name: "RegisterUserThatAlreadyExists",
			user: &User{
				Username: configs.GetViperString("testUserUsername"),
				Sub:      configs.GetViperString("testUserSub"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.Register(); (err != nil) != tt.wantErr {
				t.Errorf("User.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
