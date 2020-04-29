package modules

import (
	"fmt"

	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
)

// CurrencyTag is the intent tag for its module
var CurrencyTag = "currency"

// CurrencyReplacer replaces the pattern contained inside the response by the currency of the country
// specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func CurrencyReplacer(locale, entry, response, _ string) (string, string) {
	country := language.FindCountry(locale, entry)

	// If there isn't a country respond with a message from res/datasets/messages.json
	if country.Currency == "" {
		responseTag := "no country"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	return CurrencyTag, fmt.Sprintf(response, ArticleCountries[locale](country.Name[locale]), country.Currency)
}
