package triggers

import (
	"github.com/olivia-ai/Api/data"
	"github.com/olivia-ai/Api/language"
	"strings"
)

type Capital struct {}

// Replace the content of the response by the country and his capital
func (capital Capital) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(Response, "${CAPITAL}") {
		return Response
	}

	country := language.FindCountry(Entry)

	// If there isn't a country respond with a data.Message
	if country.Code == "" {
		return data.GetMessage("no country")
	}

	response := strings.Replace(Response, "${CAPITAL}", country.Capital, 1)
	response = strings.Replace(response, "${COUNTRY}", country.OfficialName, 1)

	return response
}
