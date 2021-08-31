package embeddings

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	distance := LevenshteinDistance("Hello", "Selol")
	result := 3

	if distance != result {
		t.Errorf("LevenshteinDistance() failed.")
	}
}

var tokenBaseSample = []string{
	"What",
	"is",
	"the",
	"capital",
	"of",
	"Farnce",
	"?",
}

func TestLevenshteinContains(t *testing.T) {
	if !LevenshteinContains(tokenBaseSample, "France", 3) {
		t.Errorf("LevenshteinContains() failed.")
	}
}

func BenchmarkLevenshteinContains(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LevenshteinContains(tokenBaseSample, "France", 3)
	}
}