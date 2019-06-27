package modules

import (
	"fmt"
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

var currencyTag = "currency"

func init() {
	RegisterModule(Module{
		Tag: currencyTag,
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

func CurrencyReplacer(entry, response string) (string, string) {
	country := language.FindCountry(entry)

	// If there isn't a country respond with a message from res/messages.json
	if country.Code == "" {
		responseTag := "no country"
		return responseTag, util.GetMessage(responseTag)
	}

	return currencyTag, fmt.Sprintf(response, country.CommonName, country.Currency)
}
