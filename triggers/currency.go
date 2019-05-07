package triggers

import (
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
	"strings"
)

type Currency struct{}

func (currency Currency) ReplaceContent() string {
	if !strings.Contains(Response, "${CURRENCY}") {
		return Response
	}

	country := language.FindCountry(Entry)

	if country.Code == "" {
		return util.GetMessage("no country")
	}

	response := strings.Replace(Response, "${COUNTRY}", country.OfficialName+" "+country.Flag, 1)
	response = strings.Replace(response, "${CURRENCY}", country.Currency, 1)

	return response
}
