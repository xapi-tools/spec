package spec

import (
	"testing"
)

func TestNewSpecFromFile(t *testing.T) {

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name: "basic",
			path: "api.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSpecFromFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpecFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
