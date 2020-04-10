package date

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const day = time.Hour * 24

var daysOfWeek = map[string]time.Weekday{
	"monday":    time.Monday,
	"tuesday":   time.Tuesday,
	"wednesday": time.Wednesday,
	"thursday":  time.Thursday,
	"friday":    time.Friday,
	"saturday":  time.Saturday,
	"sunday":    time.Sunday,
}

func init() {
	// Register the rules
	RegisterRule(RuleToday)
	RegisterRule(RuleTomorrow)
	RegisterRule(RuleDayOfWeek)
	RegisterRule(RuleNaturalDate)
	RegisterRule(RuleDate)
}

// RuleToday checks for today, tonight, this afternoon dates in the given sentence, then
// it returns the date parsed.
func RuleToday(sentence string) (result time.Time) {
	todayRegex := regexp.MustCompile("today|tonight")
	today := todayRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if today == "" {
		return time.Time{}
	}

	return time.Now()
}

// RuleTomorrow checks for "tomorrow" and "after tomorrow" dates in the given sentence, then
// it returns the date parsed.
func RuleTomorrow(sentence string) (result time.Time) {
	tomorrowRegex := regexp.MustCompile(`(after )?tomorrow`)
	date := tomorrowRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if date == "" {
		return time.Time{}
	}

	result = time.Now().Add(day)

	// If the date contains "after", we add 24 hours to tomorrow's date
	if strings.Contains(date, "after") {
		return result.Add(day)
	}

	return
}

// RuleDayOfWeek checks for the days of the week and the keyword "next" in the given sentence,
// then it returns the date parsed.
func RuleDayOfWeek(sentence string) time.Time {
	dayOfWeekRegex := regexp.MustCompile(`((next )?(monday|tuesday|wednesday|thursday|friday|saturday|sunday))`)
	date := dayOfWeekRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if date == "" {
		return time.Time{}
	}

	var foundDayOfWeek int
	// Find the integer value of the found day of the week
	for _, dayOfWeek := range daysOfWeek {
		// Down case the day of the week to match the found date
		stringDayOfWeek := strings.ToLower(dayOfWeek.String())

		if strings.Contains(date, stringDayOfWeek) {
			foundDayOfWeek = int(dayOfWeek)
		}
	}

	currentDay := int(time.Now().Weekday())
	// Calculate the date of the found day
	calculatedDate := foundDayOfWeek - currentDay

	// If the day is already passed in the current week, then we add another week to the count
	if calculatedDate <= 0 {
		calculatedDate += 7
	}

	// If there is "next" in the sentence, then we add another week
	if strings.Contains(date, "next") {
		calculatedDate += 7
	}

	// Then add the calculated number of day to the actual date
	return time.Now().Add(day * time.Duration(calculatedDate))
}

// RuleDayOfWeek checks for the dates written in natural language in the given sentence,
// then it returns the date parsed.
func RuleNaturalDate(sentence string) time.Time {
	naturalMonthRegex := regexp.MustCompile(
		`january|february|march|april|may|june|july|august|september|october|november|december`,
	)
	naturalDayRegex := regexp.MustCompile(`\d{2}|\d`)

	month := naturalMonthRegex.FindString(sentence)
	day := naturalDayRegex.FindString(sentence)

	parsedMonth, _ := time.Parse("January", month)
	parsedDay, _ := strconv.Atoi(day)

	// Returns an empty date struct if no date has been found
	if day == "" && month == "" {
		return time.Time{}
	}

	// If only the month is specified
	if day == "" {
		// Calculate the number of months to add
		calculatedMonth := parsedMonth.Month() - time.Now().Month()
		// Add a year if the month is passed
		if calculatedMonth <= 0 {
			calculatedMonth += 12
		}

		// Remove the number of days elapsed in the month to reach the first
		return time.Now().AddDate(0, int(calculatedMonth), -time.Now().Day()+1)
	}

	// Parse the date
	parsedDate := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), parsedMonth.Month(), parsedDay)
	date, err := time.Parse("2006-01-02", parsedDate)
	if err != nil {
		return time.Time{}
	}

	// If the date has been passed, add a year
	if time.Now().After(date) {
		date = date.AddDate(1, 0, 0)
	}

	return date
}

// RuleDate checks for dates written like mm/dd
func RuleDate(sentence string) time.Time {
	dateRegex := regexp.MustCompile(`(\d{2}|\d)/(\d{2}|\d)`)
	date := dateRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if date == "" {
		return time.Time{}
	}

	// Parse the found date
	parsedDate, err := time.Parse("01/02", date)
	if err != nil {
		return time.Time{}
	}

	// Add the current year to the date
	parsedDate = parsedDate.AddDate(time.Now().Year(), 0, 0)

	// Add another year if the date is passed
	if time.Now().After(parsedDate) {
		parsedDate = parsedDate.AddDate(1, 0, 0)
	}

	return parsedDate
}

// RuleTime checks for an hour written like 9pm
func RuleTime(sentence string) time.Time {
	timeRegex := regexp.MustCompile(`(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am)`)
	foundTime := timeRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if foundTime == "" {
		return time.Time{}
	}

	// Initialize the part of the day asked
	part := "am"
	if strings.Contains(foundTime, "pm") || strings.Contains(foundTime, "p.m") {
		part = "pm"
	}

	if strings.Contains(foundTime, ":") {
		// Get the hours and minutes of the given time
		hoursAndMinutesRegex := regexp.MustCompile(`(\d{2}|\d):(\d{2}|\d)`)
		timeVariables := strings.Split(hoursAndMinutesRegex.FindString(foundTime), ":")

		// Format the time with 2 digits for each
		formattedTime := fmt.Sprintf("%02s:%02s %s", timeVariables[0], timeVariables[1], part)
		response, _ := time.Parse("03:04 pm", formattedTime)

		return response
	}

	digitsRegex := regexp.MustCompile(`\d{2}|\d`)
	foundDigits := digitsRegex.FindString(foundTime)

	formattedTime := fmt.Sprintf("%02s %s", foundDigits, part)
	response, _ := time.Parse("03 pm", formattedTime)

	return response
}
