package language

import (
	"regexp"
	"strconv"
	"strings"
)

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
func FindNumberOfDecimals(entry string) int {
	decimalsRegex := regexp.MustCompile(
		`(\d+( |-)decimal(s)?)|(number (of )?decimal(s)? (is )?\d+)`,
	)
	numberRegex := regexp.MustCompile(`\d+`)

	decimals := numberRegex.FindString(decimalsRegex.FindString(entry))
	decimalsInt, _ := strconv.Atoi(decimals)

	return decimalsInt
}
