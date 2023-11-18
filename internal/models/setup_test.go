package models

import (
	"testing"

	"github.com/jkrahl/educahub-api/configs"
)

func TestConnectDatabase(t *testing.T) {
	configs.SetupViper("../../configs")

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
