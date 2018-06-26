package triggers

import (
	"../data"
	"../language"
	"strings"
)

type Capital struct {
	Entry    string
	Response string
}

func (capital Capital) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(capital.Response, "${CAPITAL}") {
		return capital.Response
	}

	country := language.FindCountry(capital.Entry)

	// If there isn't a country respond with a data.Message
	if country.Code == "" {
		return data.GetMessage("no country")
	}

	response := strings.Replace(capital.Response, "${CAPITAL}", country.Capital, 1)
	response = strings.Replace(response, "${COUNTRY}", country.OfficialName, 1)

	return response
}
