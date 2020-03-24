package language

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	distance := LevenshteinDistance("Hello", "Selol")
	result := 3

	if distance != result {
		t.Errorf("LevenshteinDistance() failed.")
	}
}
