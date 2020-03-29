package date

import "time"

// SearchTime returns the found date in the given sentence, if no date has been found, it
// returns an empty date.
func SearchTime(sentence string) time.Time {
	for _, rule := range rules {
		date := rule(sentence)

		// If the current rule found a date
		if date != (time.Time{}) {
			return date
		}
	}

	return time.Time{}
}
