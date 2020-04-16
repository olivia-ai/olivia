package modules

import (
	"fmt"
	"regexp"

	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
	"github.com/soudy/mathcat"
)

var mathTag = "math"

func init() {
	RegisterModule(Module{
		Tag: "math",
		Patterns: []string{
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

// MathReplacer replaces the pattern contained inside the response by the answer of the math
// expression specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func MathReplacer(entry, response, _ string) (string, string) {
	operation := language.FindMathOperation(entry)

	// If there is no operation in the entry message reply with a "don't understand" message
	if operation == "" {
		responseTag := "don't understand"
		return responseTag, util.GetMessage(responseTag)
	}

	res, err := mathcat.Eval(operation)
	// If the expression isn't valid reply with a message from res/datasets/messages.json
	if err != nil {
		responseTag := "math not valid"
		return responseTag, util.GetMessage(responseTag)
	}
	// Use number of decimals from the query
	decimals := language.FindNumberOfDecimals(entry)
	if decimals == 0 {
		decimals = 6
	}

	result := res.FloatString(decimals)

	// Remove trailing zeros of the result with a Regex
	trailingZerosRegex := regexp.MustCompile(`\.?0+$`)
	result = trailingZerosRegex.ReplaceAllString(result, "")

	return mathTag, fmt.Sprintf(response, result)
}
