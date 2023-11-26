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
				Username: configs.GetViperString("TEST_USER_USERNAME"),
				Sub:      configs.GetViperString("TEST_USER_SUB"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.Create(); (err != nil) != tt.wantErr {
				t.Errorf("User.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Find(t *testing.T) {
	type fields struct {
		Sub string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "FindUserThatDoesNotExist",
			fields: fields{
				Sub: "sub",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Sub: tt.fields.Sub,
			}
			if err := user.Find(); (err != nil) != tt.wantErr {
				t.Errorf("User.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
