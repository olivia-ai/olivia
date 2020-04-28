package util

import (
	"testing"
)

func TestContains(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}

	if !Contains(slice, "Hola") || Contains(slice, "bonjour") || !Contains(slice, "hey") {
		t.Errorf("slice.Contains() failed.")
	}
}

func TestDifference(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}
	slice2 := []string{"hey", "boNjour"}
	difference := Difference(slice, slice2)

	if len(difference) != 1 || difference[0] != "Hola" {
		t.Errorf("Difference() failed.")
	}
}

func TestIndex(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}

	if Index(slice, "Hola") != 1 || Index(slice, "boNjour") != 2 {
		t.Errorf("slice.Index() failed.")
	}
}
