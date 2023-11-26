package models

import (
	"testing"
)

func TestConnectDatabase(t *testing.T) {
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
