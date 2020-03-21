package util

import "testing"

func TestContains(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}

	if !Contains(slice, "Hola") || Contains(slice, "bonjour") || !Contains(slice, "hey") {
		t.Errorf("slice.Contains() failed.")
	}
}
