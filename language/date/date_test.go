package date

import "testing"

func TestDeleteDates(t *testing.T) {
	sentences := map[string]string{
		"Remind me to call mom tomorrow":                 "Remind me to call mom",
		"Remind me to cook eggs after tomorrow":          "Remind me to cook eggs",
		"Remind me that I have an exam saturday":         "Remind me that I have an exam",
		"Remind me to wash the dishes the 28th of march": "Remind me to wash the dishes",
		"Remind me the conference call of the 04/12":     "Remind me the conference call",
	}

	for sentence, excepted := range sentences {
		deleteDatesSentence := DeleteDates(sentence)

		if excepted != deleteDatesSentence {
			t.Errorf("DeleteDates() failed, excepted %s got %s.", excepted, deleteDatesSentence)
		}
	}
}
