package language

import (
	"regexp"
	"strconv"
	"strings"
)

// MathDecimals is the map for having the regex on decimals in different languages
var MathDecimals = map[string]string{
	"en": `(\d+( |-)decimal(s)?)|(number (of )?decimal(s)? (is )?\d+)`,
	"de": `(\d+( |-)decimal(s)?)|(nummer (von )?decimal(s)? (ist )?\d+)`,
	"fr": `(\d+( |-)decimale(s)?)|(nombre (de )?decimale(s)? (est )?\d+)`,
	"es": `(\d+( |-)decimale(s)?)|(numero (de )?decimale(s)? (de )?\d+)`,
	"ca": `(\d+( |-)decimal(s)?)|(nombre (de )?decimal(s)? (de )?\d+)`,
	"it": `(\d+( |-)decimale(s)?)|(numero (di )?decimale(s)? (è )?\d+)`,
	"tr": `(\d+( |-)desimal(s)?)|(numara (dan )?desimal(s)? (mı )?\d+)`,
	"nl": `(\d+( |-)decimal(en)?)|(nummer (van )?decimal(en)? (is )?\d+)`,
	"el": `(\d+( |-)δεκαδικ(ό|ά)?)|(αριθμός (από )?δεκαδικ(ό|ά)? (είναι )?\d+)`,
}

// FindMathOperation finds a math operation in a string an returns it
func FindMathOperation(entry string) string {
	mathRegex := regexp.MustCompile(
		`((\()?(((\d+|pi)(\^\d+|!|.)?)|sqrt|cos|sin|tan|acos|asin|atan|log|ln|abs)( )?[+*\/\-x]?( )?(\))?[+*\/\-]?)+`,
	)

	operation := mathRegex.FindString(entry)
	// Replace "x" symbol by "*"
	operation = strings.Replace(operation, "x", "*", -1)
	return strings.TrimSpace(operation)
}

// FindNumberOfDecimals finds the number of decimals asked in the query
func FindNumberOfDecimals(locale, entry string) int {
	decimalsRegex := regexp.MustCompile(
		MathDecimals[locale],
	)
	numberRegex := regexp.MustCompile(`\d+`)

	decimals := numberRegex.FindString(decimalsRegex.FindString(entry))
	decimalsInt, _ := strconv.Atoi(decimals)

	return decimalsInt
}
