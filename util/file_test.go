package util

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	bytes := ReadFile("res/test/test.txt")

	if string(bytes) != "test" {
		t.Errorf("file.ReadFile() failed.")
	}
}
