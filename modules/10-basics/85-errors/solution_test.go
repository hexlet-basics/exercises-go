package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		wantExt    string
		shouldFail bool
	}{
		{"Simple extension", "file.txt", "txt", false},
		{"Multiple dots", "archive.tar.gz", "gz", false},
		{"No extension", "README", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ext, err := GetFileExtension(tt.filename)
			if tt.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantExt, ext)
			}
		})
	}
}
