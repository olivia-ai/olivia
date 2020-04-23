package language

import (
	"testing"
)

func TestSearchReason(t *testing.T) {
	sentences := map[string]string{
		"Remind me to call mom":         "call mom",
		"Remind me to cook eggs":        "cook eggs",
		"Remind me that I have an exam": "I have an exam",
		"Remind me to wash the dishes":  "wash the dishes",
		"Remind me the conference call": "conference call",
	}

	for sentence, excepted := range sentences {
		reason := SearchReason("en", sentence)
		if reason != excepted {
			t.Errorf("SearchReason() failed, excepted %s got %s.", excepted, reason)
		}
	}
}

func BenchmarkSearchReason(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SearchReason("en", "Remind me to wash the dishes the 28th of march")
	}
}
