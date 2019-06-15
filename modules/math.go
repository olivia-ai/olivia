package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/util"
	"github.com/soudy/mathcat"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	RegisterModule(Module{
		Tag: "math",
		Patterns: []string{
			"What is ",
			"Give me the result of ",
			"Calculate ",
		},
		Responses: []string{
			"The result is %s",
			"That makes %s",
		},
		Replacer: MathReplacer,
	})
}

func MathReplacer(entry, response string) string {
	operation := FindMathOperation(entry)

	// If there is no operation in the entry message reply with a "don't understand" message
	if operation == "" {
		return util.GetMessage("don't understand")
	}

	res, err := mathcat.Eval(operation)
	// If the expression isn't valid reply with a message from res/messages.json
	if err != nil {
		return util.GetMessage("math not valid")
	}
	// Use number of decimals from the query
	decimals := FindNumberOfDecimals(entry)
	if decimals == 0 {
		decimals = 6
	}

	result := res.FloatString(decimals)

	// Remove trailing zeros of the result with a Regex
	trailingZerosRegex := regexp.MustCompile(`\.?0+$`)
	result = trailingZerosRegex.ReplaceAllString(result, "")

	return fmt.Sprintf(response, result)
}

// Find a math operation in a string an returns it
func FindMathOperation(entry string) string {
	mathRegex := regexp.MustCompile(
		`((\()?(((\d+|pi)(\^\d+|!|.)?)|sqrt|cos|sin|tan|acos|asin|atan|log|ln|abs)( )?[+*\/\-]?( )?(\))?[+*\/\-]?)+`,
	)

	operation := mathRegex.FindString(entry)
	return strings.TrimSpace(operation)
}

// Find the number of decimals asked in the query
func FindNumberOfDecimals(entry string) int {
	decimalsRegex := regexp.MustCompile(
		`(\d+( |-)decimal(s)?)|(number (of )?decimal(s)? (is )?\d+)`,
	)
	numberRegex := regexp.MustCompile(`\d+`)

	decimals := numberRegex.FindString(decimalsRegex.FindString(entry))
	decimalsInt, _ := strconv.Atoi(decimals)

	return decimalsInt
}
