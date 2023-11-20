package models

import (
	"testing"

	"educahub/configs"
)

func TestUser_GetUserFromSub(t *testing.T) {
	configs.SetupViper("../../configs")

	user := User{}

	type args struct {
		sub string
	}
	tests := []struct {
		name    string
		user    *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetUserFromSub",
			user: &user,
			args: args{
				sub: configs.GetViperString("testUserSub"),
			},
			wantErr: false,
		},
		{
			name: "NonExistentUser",
			user: &user,
			args: args{
				sub: "NonExistentUser",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.GetUserFromSub(tt.args.sub); (err != nil) != tt.wantErr {
				t.Errorf("User.GetUserFromSub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
