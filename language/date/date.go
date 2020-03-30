package date

import (
	"regexp"
	"strings"
	"time"
)

// SearchTime returns the found date in the given sentence and the sentence without the date, if no date has
// been found, it returns an empty date and the given sentence.
func SearchTime(sentence string) (string, time.Time) {
	_time := RuleTime(sentence)
	// Set the time to 12am if no time has been found
	if _time == (time.Time{}) {
		_time = time.Date(0, 0, 0, 12, 0, 0, 0, time.UTC)
	}

	for _, rule := range rules {
		date := rule(sentence)

		// If the current rule found a date
		if date != (time.Time{}) {
			date = time.Date(date.Year(), date.Month(), date.Day(), _time.Hour(), _time.Minute(), 0, 0, time.UTC)

			sentence = DeleteTimes(sentence)
			return DeleteDates(sentence), date
		}
	}

	return sentence, time.Time{}
}

// DeleteDates removes the dates of the given sentence and returns it
func DeleteDates(sentence string) string {
	// Create a regex to match the patterns of dates to remove them.
	datePatterns := regexp.MustCompile(
		`(of )?(the )?((after )?tomorrow|((next )?(monday|tuesday|wednesday|thursday|friday|saturday|sunday))|(\d{2}|\d)(th|rd|st|nd)? (of )?(january|february|march|april|may|june|july|august|september|october|november|december)|((\d{2}|\d)/(\d{2}|\d)))`,
	)

	// Replace the dates by empty string
	sentence = datePatterns.ReplaceAllString(sentence, "")
	// Trim the spaces and return
	return strings.TrimSpace(sentence)
}

// DeleteTimes removes the times of the given sentence and returns it
func DeleteTimes(sentence string) string {
	// Create a regex to match the patterns of times to remove them.
	timePatterns := regexp.MustCompile(`(at )?(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am)`)

	// Replace the times by empty string
	sentence = timePatterns.ReplaceAllString(sentence, "")
	// Trim the spaces and return
	return strings.TrimSpace(sentence)
}
