package embeddings

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
	condition, _ := LevenshteinContains(tokenBaseSample, "France", 3)
	if !condition {
		t.Errorf("LevenshteinContains() failed.")
	}
}

func BenchmarkLevenshteinContains(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LevenshteinContains(tokenBaseSample, "France", 3)
	}
}

func TestGetLevenshteinEmbedding(t *testing.T) {
	embedding := GetLevenshteinEmbedding(tokenBaseSample, "France")
	if embedding[7] != 1 || embedding[0] != 0 {
		t.Errorf("GetLevenshteinEmbedding() failed.")
	}
}