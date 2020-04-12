package modules

import (
	"fmt"

	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

var capitalTag = "capital"

func init() {
	RegisterModule(Module{
		Tag: capitalTag,
		Patterns: []string{
			"What is the capital of ",
			"What's the capital of ",
			"Give me the capital of ",
		},
		Responses: []string{
			"The capital of %s is %s",
		},
		Replacer: CapitalReplacer,
	})
}

// CapitalReplacer replaces the pattern contained inside the response by the capital of the country
// specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func CapitalReplacer(entry, response, _ string) (string, string) {
	country := language.FindCountry(entry)

	// If there isn't a country respond with a message from res/datasets/messages.json
	if country.Code == "" {
		responseTag := "no country"
		return responseTag, util.GetMessage(responseTag)
	}

	return capitalTag, fmt.Sprintf(response, country.CommonName, country.Capital)
}
