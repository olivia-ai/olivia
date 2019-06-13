package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/util"
	"github.com/soudy/mathcat"
	"regexp"
)

const MathExpression = `((\()?(((\d+|pi)(\^\d+|!|.)?)|sqrt|cos|sin|tan|acos|asin|atan|log|ln|abs)( )?[+*\/\-]?( )?(\))?[+*\/\-]?)+`

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
	// Find the math operation in the entry
	mathRegex := regexp.MustCompile(MathExpression)
	operation := mathRegex.FindString(entry)

	res, err := mathcat.Eval(operation)
	// If the expression isn't valid reply with a message from res/messages.json
	if err != nil {
		return util.GetMessage("math not valid")
	}
	// Arbitrary choice of 6 decimals
	result := res.FloatString(6)

	// Remove trailing zeros of the result with a Regex
	trailingZerosRegex := regexp.MustCompile(`\.?0+$`)
	result = trailingZerosRegex.ReplaceAllString(result, "")

	return fmt.Sprintf(response, result)
}
