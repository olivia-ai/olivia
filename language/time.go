package language

import (
	"regexp"
	"strings"
	"time"
)

// SearchTime returns the time found in the given sentence
func SearchTime(sentence string) {

}

// RuleTomorrow checks for "tomorrow" and "after tomorrow" dates in the given sentence, then
// it returns the date parsed.
func RuleTomorrow(sentence string) (result time.Time) {
	tomorrowRegex := regexp.MustCompile(`(after )?tomorrow`)
	date := tomorrowRegex.FindString(sentence)

	// Returns an empty time struct if no date has been found
	if date == "" {
		return time.Time{}
	}

	day := time.Hour * 24
	result = time.Now().Add(day)

	// If the date contains "after", we add 24 hours to tomorrow's date
	if strings.Contains(date, "after") {
		return result.Add(day)
	}

	return
}
