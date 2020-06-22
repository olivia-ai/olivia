package date

import (
	"regexp"
	"strings"
	"time"
)

// PatternTranslation are the map of regexs in different languages
var PatternTranslation = map[string]PatternTranslations{
	"en": {
		DateRegex: `(of )?(the )?((after )?tomorrow|((today|tonight)|(next )?(monday|tuesday|wednesday|thursday|friday|saturday|sunday))|(\d{2}|\d)(th|rd|st|nd)? (of )?(january|february|march|april|may|june|july|august|september|october|november|december)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(at )?(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am|p\.m|a\.m)`,
	},
	"de": {
		DateRegex: `(von )?(das )?((nach )?morgen|((heute|abends)|(nächsten )?(montag|dienstag|mittwoch|donnerstag|freitag|samstag|sonntag))|(\d{2}|\d)(th|rd|st|nd)? (of )?(januar|februar|märz|april|mai|juli|juli|august|september|oktober|november|dezember)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(um )?(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am|p\.m|a\.m)`,
	},
	"fr": {
		DateRegex: `(le )?(après )?demain|((aujourd'hui'|ce soir)|(lundi|mardi|mecredi|jeudi|vendredi|samedi|dimanche)( prochain)?|(\d{2}|\d) (janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(à )?(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am|p\.m|a\.m)`,
	},
	"es": {
		DateRegex: `(el )?((pasado )?mañana|((hoy|esta noche)|(el )?(proximo )?(lunes|martes|miercoles|jueves|viernes|sabado|domingo))|(\d{2}|\d) (de )?(enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(a )?(las )?(\d{2}|\d)(:\d{2}|\d)?( )?(de )?(la )?(pm|am|p\.m|a\.m|tarde|mañana)`,
	},
	"ca": {
		DateRegex: `(el )?((després )?(de )?demà|((avui|aquesta nit)|(el )?(proper )?(dilluns|dimarts|dimecres|dijous|divendres|dissabte|diumenge))|(\d{2}|\d) (de )?(gener|febrer|març|abril|maig|juny|juliol|agost|setembre|octubre|novembre|desembre)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(a )?(les )?(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am|p\.m|a\.m)`,
	},
	"nl": {
		DateRegex: `(van )?(de )?((na )?morgen|((vandaag|vanavond)|(volgende )?(maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag))|(\d{2}|\d)(te|de)? (vab )?(januari|februari|maart|april|mei|juni|juli|augustus|september|oktober|november|december)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(om )?(\d{2}|\d)(:\d{2}|\d)?( )?(pm|am|p\.m|a\.m)`,
	},
	"el": {
		DateRegex: `(από )?(το )?((μεθ )?αύριο|((σήμερα|απόψε)|(επόμενη )?(δευτέρα|τρίτη|τετάρτη|πέμπτη|παρασκευή|σάββατο|κυριακή))|(\d{2}|\d)(η)? (of )?(ιανουάριος|φεβρουάριος|μάρτιος|απρίλιος|μάιος|ιούνιος|ιούλιος|αύγουστος|σεπτέμβριος|οκτώβριος|νοέμβριος|δεκέμβριος)|((\d{2}|\d)/(\d{2}|\d)))`,
		TimeRegex: `(at )?(\d{2}|\d)(:\d{2}|\d)?( )?(μμ|πμ|μ\.μ|π\.μ)`,
	},
}

// PatternTranslations are the translations of the regexs for dates
type PatternTranslations struct {
	DateRegex string
	TimeRegex string
}

// SearchTime returns the found date in the given sentence and the sentence without the date, if no date has
// been found, it returns an empty date and the given sentence.
func SearchTime(locale, sentence string) (string, time.Time) {
	_time := RuleTime(sentence)
	// Set the time to 12am if no time has been found
	if _time == (time.Time{}) {
		_time = time.Date(0, 0, 0, 12, 0, 0, 0, time.UTC)
	}

	for _, rule := range rules {
		date := rule(locale, sentence)

		// If the current rule found a date
		if date != (time.Time{}) {
			date = time.Date(date.Year(), date.Month(), date.Day(), _time.Hour(), _time.Minute(), 0, 0, time.UTC)

			sentence = DeleteTimes(locale, sentence)
			return DeleteDates(locale, sentence), date
		}
	}

	return sentence, time.Now().Add(time.Hour * 24)
}

// DeleteDates removes the dates of the given sentence and returns it
func DeleteDates(locale, sentence string) string {
	// Create a regex to match the patterns of dates to remove them.
	datePatterns := regexp.MustCompile(PatternTranslation[locale].DateRegex)

	// Replace the dates by empty string
	sentence = datePatterns.ReplaceAllString(sentence, "")
	// Trim the spaces and return
	return strings.TrimSpace(sentence)
}

// DeleteTimes removes the times of the given sentence and returns it
func DeleteTimes(locale, sentence string) string {
	// Create a regex to match the patterns of times to remove them.
	timePatterns := regexp.MustCompile(PatternTranslation[locale].TimeRegex)

	// Replace the times by empty string
	sentence = timePatterns.ReplaceAllString(sentence, "")
	// Trim the spaces and return
	return strings.TrimSpace(sentence)
}
