package language

import (
	"testing"
	"time"
)

// CheckDate checks if the two given dates are the same, if not it prints the error
func CheckDate(t *testing.T, a, b time.Time) {
	if a.Day() != b.Day() || a.Year() != b.Year() || a.Month() != b.Month() {
		t.Errorf("RuleTomorrow() failed, excepted %s got %s.", a, b)
	}
}

func TestRuleTomorrow(t *testing.T) {
	sentence := "Remind me to call mom tomorrow"
	excepted := time.Now().Add(time.Hour * 24)
	date := RuleTomorrow(sentence)

	CheckDate(t, excepted, date)

	sentence = "Remind me to call mom after tomorrow"
	excepted = time.Now().Add(time.Hour * 48)
	date = RuleTomorrow(sentence)

	CheckDate(t, excepted, date)
}
