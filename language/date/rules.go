package date

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olivia-ai/olivia/util"
)

const day = time.Hour * 24

// RuleTranslations are the translations of the rules in different languages
var RuleTranslations = map[string]RuleTranslation{
	"en": {
		DaysOfWeek: []string{
			"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday",
		},
		Months: []string{
			"january", "february", "march", "april", "may", "june", "july",
			"august", "september", "october", "november", "december",
		},
		RuleToday:         `today|tonight`,
		RuleTomorrow:      `(after )?tomorrow`,
		RuleAfterTomorrow: "after",
		RuleDayOfWeek:     `(next )?(monday|tuesday|wednesday|thursday|friday|saturday|sunday)`,
		RuleNextDayOfWeek: "next",
		RuleNaturalDate:   `january|february|march|april|may|june|july|august|september|october|november|december`,
	},
	"de": {
		DaysOfWeek: []string{
			"Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag", "Sonntag",
		},
		Months: []string{
			"Januar", "Februar", "Marsch", "April", "Mai", "Juni", "Juli",
			"August", "September", "Oktober", "November", "Dezember",
		},
		RuleToday:         `heute|abends`,
		RuleTomorrow:      `(nach )?tomorrow`,
		RuleAfterTomorrow: "nach",
		RuleDayOfWeek:     `(nächsten )?(Montag|Dienstag|Mittwoch|Donnerstag|Freitag|Samstag|Sonntag)`,
		RuleNextDayOfWeek: "nächste",
		RuleNaturalDate:   `Januar|Februar|März|April|Mai|Juli|Juli|August|September|Oktober|November|Dezember`,
	},
	"fr": {
		DaysOfWeek: []string{
			"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche",
		},
		Months: []string{
			"janvier", "février", "mars", "avril", "mai", "juin", "juillet",
			"août", "septembre", "octobre", "novembre", "décembre",
		},
		RuleToday:         `aujourd'hui|ce soir`,
		RuleTomorrow:      `(après )?demain`,
		RuleAfterTomorrow: "après",
		RuleDayOfWeek:     `(lundi|mardi|mecredi|jeudi|vendredi|samedi|dimanche)( prochain)?`,
		RuleNextDayOfWeek: "prochain",
		RuleNaturalDate:   `janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre`,
	},
	"es": {
		DaysOfWeek: []string{
			"lunes", "martes", "miercoles", "jueves", "viernes", "sabado", "domingo",
		},
		Months: []string{
			"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio",
			"agosto", "septiembre", "octubre", "noviembre", "diciembre",
		},
		RuleToday:         `hoy|esta noche`,
		RuleTomorrow:      `(pasado )?mañana`,
		RuleAfterTomorrow: "pasado",
		RuleDayOfWeek:     `(el )?(proximo )?(lunes|martes|miercoles|jueves|viernes|sabado|domingo))`,
		RuleNextDayOfWeek: "proximo",
		RuleNaturalDate:   `enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre`,
	},
	"ca": {
		DaysOfWeek: []string{
			"dilluns", "dimarts", "dimecres", "dijous", "divendres", "dissabte", "diumenge",
		},
		Months: []string{
			"gener", "febrer", "març", "abril", "maig", "juny", "juliol",
			"agost", "setembre", "octubre", "novembre", "desembre",
		},
		RuleToday:         `avui|aquesta nit`,
		RuleTomorrow:      `((després )?(de )?demà`,
		RuleAfterTomorrow: "després",
		RuleDayOfWeek:     `(el )?(proper )?(dilluns|dimarts|dimecres|dijous|divendres|dissabte|diumenge))`,
		RuleNextDayOfWeek: "proper",
		RuleNaturalDate:   `gener|febrer|març|abril|maig|juny|juliol|agost|setembre|octubre|novembre|desembre`,
	},
	"nl": {
		DaysOfWeek: []string{
			"maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag", "zondag",
		},
		Months: []string{
			"januari", "februari", "maart", "april", "mei", "juni", "juli",
			"augustus", "september", "oktober", "november", "december",
		},
		RuleToday:         `vandaag|vanavond`,
		RuleTomorrow:      `(na )?morgen`,
		RuleAfterTomorrow: "na",
		RuleDayOfWeek:     `(volgende )?(maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag)`,
		RuleNextDayOfWeek: "volgende",
		RuleNaturalDate:   `januari|februari|maart|april|mei|juni|juli|augustus|september|oktober|november|december`,
	},
	"el": {
		DaysOfWeek: []string{
			"δευτέρα", "τρίτη", "τετάρτη", "πέμπτη", "παρασκευή", "σάββατο", "κυριακή",
		},
		Months: []string{
			"ιανουάριος", "φεβρουάριος", "μάρτιος", "απρίλιος", "μάιος", "ιούνιος", "ιούλιος",
			"αύγουστος", "σεπτέμβριος", "οκτώβριος", "νοέμβριος", "δεκέμβριος",
		},
		RuleToday:         `σήμερα|απόψε`,
		RuleTomorrow:      `(μεθ )?άυριο`,
		RuleAfterTomorrow: "μεθ",
		RuleDayOfWeek:     `(επόμενη )?(δευτέρα|τρίτη|τετάρτη|πέμπτη|παρασκευή|σάββατο|κυριακή)`,
		RuleNextDayOfWeek: "επόμενη",
		RuleNaturalDate:   `ιανουάριος|φεβρουάριος|μάρτιος|απρίλιος|μάιος|ιούνιος|ιούλιος|αύγουστος|σεπτέμβριος|οκτώβριος|νοέμβριος|δεκέμβριος`,
	},
}

// A RuleTranslation is all the texts/regexs to match the dates
type RuleTranslation struct {
	DaysOfWeek        []string
	Months            []string
	RuleToday         string
	RuleTomorrow      string
	RuleAfterTomorrow string
	RuleDayOfWeek     string
	RuleNextDayOfWeek string
	RuleNaturalDate   string
}

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
func RuleToday(locale, sentence string) (result time.Time) {
	todayRegex := regexp.MustCompile(RuleTranslations[locale].RuleToday)
	today := todayRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if today == "" {
		return time.Time{}
	}

	return time.Now()
}

// RuleTomorrow checks for "tomorrow" and "after tomorrow" dates in the given sentence, then
// it returns the date parsed.
func RuleTomorrow(locale, sentence string) (result time.Time) {
	tomorrowRegex := regexp.MustCompile(RuleTranslations[locale].RuleTomorrow)
	date := tomorrowRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if date == "" {
		return time.Time{}
	}

	result = time.Now().Add(day)

	// If the date contains "after", we add 24 hours to tomorrow's date
	if strings.Contains(date, RuleTranslations[locale].RuleAfterTomorrow) {
		return result.Add(day)
	}

	return
}

// RuleDayOfWeek checks for the days of the week and the keyword "next" in the given sentence,
// then it returns the date parsed.
func RuleDayOfWeek(locale, sentence string) time.Time {
	dayOfWeekRegex := regexp.MustCompile(RuleTranslations[locale].RuleDayOfWeek)
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
	if strings.Contains(date, RuleTranslations[locale].RuleNextDayOfWeek) {
		calculatedDate += 7
	}

	// Then add the calculated number of day to the actual date
	return time.Now().Add(day * time.Duration(calculatedDate))
}

// RuleNaturalDate checks for the dates written in natural language in the given sentence,
// then it returns the date parsed.
func RuleNaturalDate(locale, sentence string) time.Time {
	naturalMonthRegex := regexp.MustCompile(
		RuleTranslations[locale].RuleNaturalDate,
	)
	naturalDayRegex := regexp.MustCompile(`\d{2}|\d`)

	month := naturalMonthRegex.FindString(sentence)
	day := naturalDayRegex.FindString(sentence)

	// Put the month in english to parse the time with time golang package
	if locale != "en" {
		monthIndex := util.Index(RuleTranslations[locale].Months, month)
		month = RuleTranslations["en"].Months[monthIndex]
	}

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
func RuleDate(locale, sentence string) time.Time {
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
	timeRegex := regexp.MustCompile(`(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am|p\.m|a\.m)`)
	foundTime := timeRegex.FindString(sentence)

	// Returns an empty date struct if no date has been found
	if foundTime == "" {
		return time.Time{}
	}

	// Initialize the part of the day asked
	var part string
	if strings.Contains(foundTime, "pm") || strings.Contains(foundTime, "p.m") {
		part = "pm"
	} else if strings.Contains(foundTime, "am") || strings.Contains(foundTime, "a.m") {
		part = "am"
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
