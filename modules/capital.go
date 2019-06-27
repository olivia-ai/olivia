package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

func init() {
	RegisterModule(Module{
		Tag: "capital",
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

func CapitalReplacer(entry, response string) string {
	country := language.FindCountry(entry)

	// If there isn't a country respond with a message from res/messages.json
	if country.Code == "" {
		return util.GetMessage("no country")
	}

	return fmt.Sprintf(response, country.CommonName, country.Capital)
}
