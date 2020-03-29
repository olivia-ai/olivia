package date

import (
	"testing"
)

func TestRuleDayOfWeek(t *testing.T) {
	sentence := "Remind me that I have an exam saturday"
	excepted := 6
	weekday := int(RuleDayOfWeek(sentence).Weekday())

	if excepted != weekday {
		t.Errorf("RuleDayOfWeek() failed, excepted %d got %d.", excepted, weekday)
	}
}
