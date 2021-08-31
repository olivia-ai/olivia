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

func TestReadCSV(t *testing.T) {
	data := ReadCSV("data/mock.csv")

	for i, line := range data {
		if i == 0 && (line[0] != "How are you?" || line[1] != "Good and you?") {
			t.Errorf("file.ReadCSV() failed.")
		}
	}
}