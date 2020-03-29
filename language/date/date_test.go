package date

import (
	"testing"
	"time"
)

// CheckEquality checks if the two given dates are the same
func CheckEquality(a, b time.Time) bool {
	return a.Day() == b.Day() || a.Year() == b.Year() || a.Month() == b.Month()
}

func TestSearchTime(t *testing.T) {
	day := time.Hour * 24

	sentences := map[string]time.Time{
		"Remind me to call mom tomorrow":       time.Now().Add(day),
		"Remind me to call mom after tomorrow": time.Now().Add(day * 2),
	}

	for sentence, date := range sentences {
		foundDate := SearchTime(sentence)
		if !CheckEquality(date, foundDate) {
			t.Errorf("SearchTime() failed, excepted %s got %s.", date, foundDate)
		}
	}
}
