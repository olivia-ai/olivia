package date

import (
	"testing"
	"time"
)

// CheckEquality checks if the two given dates are the same
func CheckEquality(a, b time.Time) bool {
	return a.Day() == b.Day() || a.Year() == b.Year() || a.Month() == b.Month()
}

func TestRuleTomorrow(t *testing.T) {
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

func TestRuleDayOfWeek(t *testing.T) {
	sentence := "Remind me that I have an exam saturday"
	excepted := 6
	weekday := int(RuleDayOfWeek(sentence).Weekday())

	if excepted != weekday {
		t.Errorf("RuleDayOfWeek() failed, excepted %d got %d.", excepted, weekday)
	}
}

func TestRuleNaturalDate(t *testing.T) {
	sentence := "Nothing here"
	date := RuleNaturalDate(sentence)
	excepted := time.Time{}

	if date != excepted {
		t.Errorf("RuleNaturalDate() failed, excepted %s got %s.", excepted, date)
	}

	sentence = "Remind me that I have an exam the 28th of march"
	date = RuleNaturalDate(sentence)

	if date.Month() != 3 || date.Day() != 28 {
		t.Errorf("RuleNaturalDate() failed, excepted 3/28 got %s.", date)
	}

	sentence = "Remind me that I have an exam in december"
	date = RuleNaturalDate(sentence)

	if date.Month() != 12 || date.Day() != 1 {
		t.Errorf("RuleNaturalDate() failed, excepted 1/12 got %s.", date)
	}
}
