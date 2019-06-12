package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

func init() {
	RegisterModule(Module{
		Tag: "currency",
		Patterns: []string{
			"Which currency is used in ",
			"Give me the used currency of ",
			"Give me the currency of ",
			"What is the currency of ",
		},
		Responses: []string{
			"The currency of %s is %s",
		},
		Replacer: CurrencyReplacer,
	})
}

func CurrencyReplacer(entry, response string) string {
	country := language.FindCountry(entry)

	// If there isn't a country respond with a message from res/messages.json
	if country.Code == "" {
		return util.GetMessage("no country")
	}

	return fmt.Sprintf(response, country.CommonName, country.Currency)
}
