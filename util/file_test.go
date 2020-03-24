package util

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	bytes := ReadFile(".env")

	if string(bytes) != "PORT=8080" {
		t.Errorf("file.ReadFile() failed.")
	}
}
