package models

import (
	"testing"

	"educahub/configs"
)

func TestConnectDatabase(t *testing.T) {
	configs.SetupViper()

	tests := []struct {
		name string
	}{
		{
			name: "ConnectDatabase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectDatabase()
		})
	}
}
