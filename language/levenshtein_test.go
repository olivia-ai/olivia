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

func TestLevenshteinContains(t *testing.T) {
	if !LevenshteinContains("What is the capital of Frnaec ?", "France", 3) {
		t.Errorf("LevenshteinContains() failed.")
	}
}

func BenchmarkLevenshteinContains(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LevenshteinContains("What is the capital of Frnaec ?", "France", 3)
	}
}
