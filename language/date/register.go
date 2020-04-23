package date

import "time"

// A Rule is a function that takes the given sentence and tries to parse a specific
// rule to return a date, if not, the date is empty.
type Rule func(string, string) time.Time

var rules []Rule

// RegisterRule takes a rule in parameter and register it to the array of rules
func RegisterRule(rule Rule) {
	rules = append(rules, rule)
}
